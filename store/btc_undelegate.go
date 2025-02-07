package store

import (
	"encoding/binary"
	"encoding/json"
	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
)

type BtcUndelegate struct {
	BU     types.MsgBTCUndelegate
	TxHash []byte `json:"tx_hash"`
}

func (s *Storage) SetBtcUndelegateMsg(batchId uint64, blockHeight uint64, msg BtcUndelegate) error {
	bz, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	key, err := s.GetBabylonDelegationKey([]byte(msg.BU.StakingTxHash))
	if err != nil {
		return err
	}
	exist, cBD := s.GetCreateBTCDelegationMsg(key)
	if exist {
		err = s.setBTCUnDelegateAmount([]byte(cBD.CBD.StakerAddr), uint64(cBD.CBD.StakingValue))
		if err != nil {
			return err
		}
		err = s.SetStakeDetails(batchId, blockHeight, cBD, UndelegateType)
		if err != nil {
			return err
		}
	} else {
		return nil
	}

	return s.db.Put(getBtcUndelegateKey(msg.TxHash), bz, nil)
}

func (s *Storage) GetBtcUndelegateMsg(txHash []byte) (bool, BtcUndelegate) {
	cBDB, err := s.db.Get(getBtcUndelegateKey(txHash), nil)
	if err != nil {
		return handleError2(BtcUndelegate{}, err)
	}
	var cBD BtcUndelegate
	if err = json.Unmarshal(cBDB, &cBD); err != nil {
		return false, BtcUndelegate{}
	}
	return true, cBD
}

func (s *Storage) setBTCUnDelegateAmount(address []byte, amount uint64) error {
	amountBz := make([]byte, 8)
	amountB, err := s.db.Get(getBTCDelegateAmountKey(address), nil)
	if err != nil {
		return err
	}
	a1 := binary.BigEndian.Uint64(amountB)
	binary.BigEndian.PutUint64(amountBz, a1-amount)
	return s.db.Put(getBTCDelegateAmountKey(address), amountBz, nil)
}
