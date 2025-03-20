package store

import (
	"encoding/json"
)

type CelestiaBlockHeader struct {
	Hash       []byte `json:"hash"`
	ParentHash []byte `json:"parent_hash"`
	Number     uint64 `json:"number"`
	Timestamp  int64  `json:"timestamp"`
}

func (s *Storage) SetCelestiaBlockHeader(header CelestiaBlockHeader) error {
	cz, err := json.Marshal(header)
	if err != nil {
		return err
	}
	return s.db.Put(getCelestiaBlockHeaderKey(header.Number), cz, nil)
}

func (s *Storage) SetCelestiaBlockHeaders(headers []CelestiaBlockHeader) error {
	for _, header := range headers {
		err := s.SetCelestiaBlockHeader(header)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) GetCelestiaBlockHeader(number uint64) (bool, CelestiaBlockHeader) {
	chz, err := s.db.Get(getCelestiaBlockHeaderKey(number), nil)
	if err != nil {
		return handleError2(CelestiaBlockHeader{}, err)
	}
	var ch CelestiaBlockHeader
	if err = json.Unmarshal(chz, &ch); err != nil {
		return false, CelestiaBlockHeader{}
	}
	return true, ch
}
