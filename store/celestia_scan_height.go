package store

import (
	"encoding/binary"
)

func (s *Storage) UpdateCelestiaHeight(height uint64) error {
	heightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBz, height)
	return s.db.Put(getCelestiaScannedHeightKey(), heightBz, nil)
}

func (s *Storage) GetCelestiaScannedHeight() (uint64, error) {
	bz, err := s.db.Get(getCelestiaScannedHeightKey(), nil)
	if err != nil {
		return handleError(uint64(0), err)
	}
	return binary.BigEndian.Uint64(bz), nil
}

func (s *Storage) ResetCelestiaScanHeight(height uint64) error {
	heightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBz, height)
	return s.db.Put(getCelestiaScannedHeightKey(), heightBz, nil)
}
