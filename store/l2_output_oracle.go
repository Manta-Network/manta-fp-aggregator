package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/syndtr/goleveldb/leveldb/util"
)

type OutputProposed struct {
	StateRoot     string      `json:"state_root"`
	L2BlockNumber *big.Int    `json:"l2_block_number"`
	L2OutputIndex *big.Int    `json:"l2_output_index"`
	L1BlockHash   common.Hash `json:"l1_block_hash"`
	L1BlockNumber uint64      `json:"l1_block_number"`
	Timestamp     *big.Int    `json:"timestamp"`
}

func (s *Storage) SetOutputProposed(op OutputProposed) error {
	bz, err := json.Marshal(op)
	if err != nil {
		return err
	}
	return s.db.Put(getOutputProposedKey(op.Timestamp.Uint64()), bz, nil)
}

func (s *Storage) GetOutputProposedByTimestamp(timestamp uint64) (OutputProposed, error) {
	opb, err := s.db.Get(getOutputProposedKey(timestamp), nil)
	if err != nil {
		return handleError(OutputProposed{}, err)
	}

	var op OutputProposed
	if err = json.Unmarshal(opb, &op); err != nil {
		return OutputProposed{}, err
	}
	return op, nil
}

func (s *Storage) GetLatestUnprocessedStateRoot(timestamp, limit uint64) (*OutputProposed, error) {
	iter := s.db.NewIterator(&util.Range{Start: getOutputProposedKey(timestamp + 1), Limit: getOutputProposedKey(timestamp + limit)}, nil)
	defer iter.Release()

	if iter.Next() {
		var output OutputProposed
		if err := json.Unmarshal(iter.Value(), &output); err != nil {
			return nil, err
		}
		return &output, nil
	}

	if err := iter.Error(); err != nil {
		return nil, fmt.Errorf("iterator error: %v", err)
	}

	return nil, nil

}

func (s *Storage) SetLatestProcessedStateRoot(op OutputProposed) error {
	bz, err := json.Marshal(op)
	if err != nil {
		return err
	}
	return s.db.Put(getFinalityOutputProposedKey(), bz, nil)
}

func (s *Storage) GetLatestProcessedStateRoot() (*OutputProposed, error) {
	opb, err := s.db.Get(getFinalityOutputProposedKey(), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	var op OutputProposed
	if err = json.Unmarshal(opb, &op); err != nil {
		return nil, err
	}
	return &op, nil
}
