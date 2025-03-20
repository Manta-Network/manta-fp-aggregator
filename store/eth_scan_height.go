package store

import (
	"encoding/binary"
)

func (s *Storage) UpdateEthHeight(height uint64) error {
	heightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBz, height)
	return s.db.Put(getEthScannedHeightKey(), heightBz, nil)
}

func (s *Storage) GetEthScannedHeight() (uint64, error) {
	bz, err := s.db.Get(getEthScannedHeightKey(), nil)
	if err != nil {
		return handleError(uint64(0), err)
	}
	return binary.BigEndian.Uint64(bz), nil
}

func (s *Storage) ResetEthScanHeight(height uint64) error {
	heightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBz, height)
	return s.db.Put(getEthScannedHeightKey(), heightBz, nil)
}
