package store

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
)

type CreateBTCDelegation struct {
	CBD    types.MsgCreateBTCDelegation
	TxHash []byte `json:"tx_hash"`
}

func (s *Storage) SetCreateBTCDelegationMsg(batchId uint64, blockHeight uint64, msg CreateBTCDelegation) error {
	bz, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = s.setBTCDelegateAmount([]byte(msg.CBD.StakerAddr), uint64(msg.CBD.StakingValue))
	if err != nil {
		return err
	}

	err = s.SetStakeDetails(batchId, blockHeight, msg, DelegateType)
	if err != nil {
		return err
	}

	return s.db.Put(getCreateBTCDelegationKey(msg.TxHash), bz, nil)
}

func (s *Storage) GetCreateBTCDelegationMsg(txHash []byte) (bool, CreateBTCDelegation) {
	cBDB, err := s.db.Get(getCreateBTCDelegationKey(txHash), nil)
	if err != nil {
		return handleError2(CreateBTCDelegation{}, err)
	}
	var cBD CreateBTCDelegation
	if err = json.Unmarshal(cBDB, &cBD); err != nil {
		return false, CreateBTCDelegation{}
	}
	return true, cBD
}

func (s *Storage) GetBTCDelegateAmount(address []byte) (uint64, error) {
	amountB, err := s.db.Get(getBTCDelegateAmountKey(address), nil)
	if err != nil {
		return handleError(uint64(0), err)
	}
	return binary.BigEndian.Uint64(amountB), nil
}

func (s *Storage) setBTCDelegateAmount(address []byte, amount uint64) error {
	amountBz := make([]byte, 8)
	amountB, err := s.db.Get(getBTCDelegateAmountKey(address), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			binary.BigEndian.PutUint64(amountBz, amount)
			return s.db.Put(getBTCDelegateAmountKey(address), amountBz, nil)
		} else {
			return err
		}
	}
	a1 := binary.BigEndian.Uint64(amountB)
	binary.BigEndian.PutUint64(amountBz, a1+amount)
	return s.db.Put(getBTCDelegateAmountKey(address), amountBz, nil)
}

func (s *Storage) SetBabylonDelegationKey(babylonTx []byte, btcTx []byte) error {
	return s.db.Put(getBabylonDelegationKey(btcTx), babylonTx, nil)
}

func (s *Storage) GetBabylonDelegationKey(btcTx []byte) ([]byte, error) {
	babylonTx, err := s.db.Get(getBabylonDelegationKey(btcTx), nil)
	if err != nil {
		return handleError([]byte(""), err)
	}
	return babylonTx, nil
}
