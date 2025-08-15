package store

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/syndtr/goleveldb/leveldb"
)

type OperatorRegistered struct {
	BlockNumber int64          `json:"block_number"`
	TxHash      common.Hash    `json:"tx_hash"`
	Operator    common.Address `json:"operator"`
	NodeUrl     string         `json:"node_url"`
	Timestamp   uint64         `json:"timestamp"`
}

type NodeMembers struct {
	Members []string `json:"members"`
}

func (s *Storage) SetOperatorRegisteredEvent(event OperatorRegistered) error {
	var nodeMembers NodeMembers
	nMB, err := s.db.Get(getActiveMemberKey(), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			nodeMembers.Members = append(nodeMembers.Members, event.NodeUrl)
			bn, err := json.Marshal(nodeMembers)
			if err != nil {
				return err
			}
			return s.db.Put(getActiveMemberKey(), bn, nil)
		} else {
			return err
		}
	}

	if err = json.Unmarshal(nMB, &nodeMembers); err != nil {
		return err
	}
	nodeMembers.Members = append(nodeMembers.Members, event.NodeUrl)
	nM, err := json.Marshal(nodeMembers)
	if err != nil {
		return err
	}
	return s.db.Put(getActiveMemberKey(), nM, nil)
}

func (s *Storage) SetActiveMember(member string) error {
	var nodeMembers NodeMembers
	nMB, err := s.db.Get(getActiveMemberKey(), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			nodeMembers.Members = append(nodeMembers.Members, member)
			bn, err := json.Marshal(nodeMembers)
			if err != nil {
				return err
			}
			return s.db.Put(getActiveMemberKey(), bn, nil)
		} else {
			return err
		}
	}
	if err = json.Unmarshal(nMB, &nodeMembers); err != nil {
		return err
	}

	if contains(nodeMembers.Members, member) {
		return nil
	} else {
		nodeMembers.Members = append(nodeMembers.Members, member)
	}

	nM, err := json.Marshal(nodeMembers)
	if err != nil {
		return err
	}
	return s.db.Put(getActiveMemberKey(), nM, nil)
}

func (s *Storage) GetActiveMember() (NodeMembers, error) {
	aMB, err := s.db.Get(getActiveMemberKey(), nil)
	if err != nil {
		return handleError(NodeMembers{}, err)
	}

	var nM NodeMembers
	if err = json.Unmarshal(aMB, &nM); err != nil {
		return NodeMembers{}, err
	}
	return nM, nil
}

func (s *Storage) DeleteUnusedMembers(activeMembers []string) error {
	aMB, err := s.db.Get(getActiveMemberKey(), nil)
	if err != nil {
		return err
	}

	var nM NodeMembers
	if err = json.Unmarshal(aMB, &nM); err != nil {
		return err
	}

	nM.Members = activeMembers
	nMB, err := json.Marshal(nM)
	if err != nil {
		return err
	}
	return s.db.Put(getActiveMemberKey(), nMB, nil)
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
