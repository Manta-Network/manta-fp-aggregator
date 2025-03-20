package store

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/syndtr/goleveldb/leveldb"
)

type SymbioticFpIds struct {
	BatchId          uint64        `json:"batch_id"`
	SignRequests     []SignRequest `json:"sign_requests"`
	TotalStakeAmount *big.Int      `json:"total_stake_amount"`
}

type SignRequest struct {
	StateRoot   [32]byte `json:"state_root"`
	Signature   []byte   `json:"signature"`
	SignAddress string   `json:"sign_address"`
}

func (s *Storage) SetSymbioticFpIds(id SymbioticFpIds) error {
	sFz, err := s.db.Get(getSymbioticFpIdsKey(id.BatchId), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			sF, err := json.Marshal(id)
			if err != nil {
				return err
			}
			return s.db.Put(getSymbioticFpIdsKey(id.BatchId), sF, nil)
		}
	}
	var sF SymbioticFpIds
	if err = json.Unmarshal(sFz, &sF); err != nil {
		return err
	}

	sF.TotalStakeAmount = new(big.Int).Add(sF.TotalStakeAmount, id.TotalStakeAmount)
	sF.SignRequests = append(sF.SignRequests, id.SignRequests[0])

	sFz, err = json.Marshal(sF)
	if err != nil {
		return err
	}
	return s.db.Put(getSymbioticFpIdsKey(id.BatchId), sFz, nil)
}

func (s *Storage) GetSymbioticFpIds(batchId uint64) (SymbioticFpIds, error) {
	sFz, err := s.db.Get(getSymbioticFpIdsKey(batchId), nil)
	if err != nil {
		return handleError(SymbioticFpIds{}, err)
	}
	var sF SymbioticFpIds
	if err = json.Unmarshal(sFz, &sF); err != nil {
		return SymbioticFpIds{}, err
	}
	return sF, nil
}
