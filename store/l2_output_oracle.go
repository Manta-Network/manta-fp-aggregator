package store

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type L2OutputOracle struct {
	StateRoot     string      `json:"state_root"`
	L2BlockNumber *big.Int    `json:"l2_block_number"`
	L2OutputIndex *big.Int    `json:"l2_output_index"`
	L1BlockHash   common.Hash `json:"l1_block_hash"`
	L1BlockNumber uint64      `json:"l1_block_number"`
	Timestamp     *big.Int    `json:"timestamp"`
}

func (s *Storage) SetL2OutputOracle(l2oo L2OutputOracle) error {
	bz, err := json.Marshal(l2oo)
	if err != nil {
		return err
	}
	return s.db.Put(getL2OutputOracleKey(l2oo.L1BlockNumber), bz, nil)
}

func (s *Storage) GetL2OutputOracle(l1BlockNumber uint64) (bool, L2OutputOracle) {
	l2oob, err := s.db.Get(getL2OutputOracleKey(l1BlockNumber), nil)
	if err != nil {
		return handleError2(L2OutputOracle{}, err)
	}
	var l2oo L2OutputOracle
	if err = json.Unmarshal(l2oob, &l2oo); err != nil {
		return false, L2OutputOracle{}
	}
	return true, l2oo
}
