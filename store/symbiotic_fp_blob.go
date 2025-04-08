package store

import (
	"encoding/json"
	"errors"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type SymbioticFpBlob struct {
	SignRequests SignRequest `json:"sign_requests"`
	Timestamp    uint64      `json:"timestamp"`
}

type SignRequest struct {
	StateRoot   string `json:"state_root"`
	Signature   []byte `json:"signature"`
	SignAddress string `json:"sign_address"`
}

type BatchSymbioticFpBlob struct {
	SymbioticFpBlobs []SymbioticFpBlob `json:"symbiotic_fp_blobs"`
}

func (s *Storage) SetSymbioticFpBlob(blob SymbioticFpBlob) error {
	var bSF BatchSymbioticFpBlob
	sFz, err := s.db.Get(getSymbioticFpBlobKey(blob.Timestamp), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			bSF.SymbioticFpBlobs = append(bSF.SymbioticFpBlobs, blob)
			sF, err := json.Marshal(bSF)
			if err != nil {
				return err
			}
			return s.db.Put(getSymbioticFpBlobKey(blob.Timestamp), sF, nil)
		}
	}
	if err = json.Unmarshal(sFz, &bSF); err != nil {
		return err
	}
	bSF.SymbioticFpBlobs = append(bSF.SymbioticFpBlobs, blob)
	sFz, err = json.Marshal(bSF)
	if err != nil {
		return err
	}

	return s.db.Put(getSymbioticFpBlobKey(blob.Timestamp), sFz, nil)
}

func (s *Storage) GetSymbioticFpBlobsByTimestamp(start uint64, end uint64) ([]SymbioticFpBlob, error) {
	var batchSymbioticFpBlobs []BatchSymbioticFpBlob
	var results []SymbioticFpBlob
	iter := s.db.NewIterator(&util.Range{Start: getSymbioticFpBlobKey(start), Limit: getSymbioticFpBlobKey(end)}, nil)
	defer iter.Release()

	for iter.Next() {
		var bsfb BatchSymbioticFpBlob
		if err := json.Unmarshal(iter.Value(), &bsfb); err != nil {
			return nil, err
		}
		batchSymbioticFpBlobs = append(batchSymbioticFpBlobs, bsfb)
	}

	if err := iter.Error(); err != nil {
		return nil, err
	}

	for _, batchSymbioticFpBlob := range batchSymbioticFpBlobs {
		for _, blob := range batchSymbioticFpBlob.SymbioticFpBlobs {
			results = append(results, blob)
		}
	}

	return results, nil
}
