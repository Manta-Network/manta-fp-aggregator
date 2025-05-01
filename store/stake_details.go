package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Manta-Network/manta-fp-aggregator/manager/types"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

const (
	DelegateType   = 0
	UndelegateType = iota
)

type Staker struct {
	StakingAmount uint64 `json:"stakingAmount"`
	Address       string `json:"address"`
}

type QuorumNode struct {
	FpBtcPk      string   `json:"fpBtcPk"`
	FpVoteWeight uint64   `json:"fpVoteWeight"`
	IsSign       bool     `json:"isSign"`
	Staker       []Staker `json:"staker"`
}

type StakeDetails struct {
	BatchID           uint64       `json:"batchId"`
	TotalBTCVote      uint64       `json:"totalBtcVote"`
	BabylonBlock      uint64       `json:"babylonBlock"`
	StateRoot         string       `json:"stateRoot"`
	EthBlock          uint64       `json:"ethBlock"`
	BitcoinQuorum     []QuorumNode `json:"bitcoinQuorum"`
	SymbioticSignNode []string     `json:"symbioticSignNode"`
}

func (s *Storage) SetStakeDetailsTest(sD StakeDetails) error {
	bsD, err := json.Marshal(sD)
	if err != nil {
		return err
	}
	return s.db.Put(getStakeDetailsKey(), bsD, nil)
}

func (s *Storage) SetStakeDetails(msg CreateBTCDelegation, stakeType int8) error {
	if stakeType == DelegateType {
		stakeDB, err := s.db.Get(getStakeDetailsKey(), nil)
		if err != nil {
			if errors.Is(err, leveldb.ErrNotFound) {
				var sD = StakeDetails{
					BitcoinQuorum: []QuorumNode{
						{
							FpBtcPk:      msg.CBD.FpBtcPkList[0].MarshalHex(),
							FpVoteWeight: uint64(msg.CBD.StakingValue),
							Staker: []Staker{
								{
									StakingAmount: uint64(msg.CBD.StakingValue),
									Address:       msg.CBD.StakerAddr,
								},
							},
						},
					},
				}
				bz, err := json.Marshal(sD)
				if err != nil {
					return err
				}
				return s.db.Put(getStakeDetailsKey(), bz, nil)
			} else {
				return err
			}
		}

		var sD StakeDetails
		if err = json.Unmarshal(stakeDB, &sD); err != nil {
			return err
		}
		var existFpBtcPk bool
		var existStaker bool
		for _, quorum := range sD.BitcoinQuorum {
			if quorum.FpBtcPk == msg.CBD.FpBtcPkList[0].MarshalHex() {
				quorum.FpVoteWeight += uint64(msg.CBD.StakingValue)
				existFpBtcPk = true
				for _, staker := range quorum.Staker {
					if staker.Address == msg.CBD.StakerAddr {
						existStaker = true
						staker.StakingAmount += uint64(msg.CBD.StakingValue)
					}
				}
			}
		}

		if !existFpBtcPk {
			var quorum = QuorumNode{
				FpBtcPk:      msg.CBD.FpBtcPkList[0].MarshalHex(),
				FpVoteWeight: uint64(msg.CBD.StakingValue),
				Staker: []Staker{
					{
						StakingAmount: uint64(msg.CBD.StakingValue),
						Address:       msg.CBD.StakerAddr,
					},
				},
			}
			sD.BitcoinQuorum = append(sD.BitcoinQuorum, quorum)
		} else {
			if !existStaker {
				for _, quorum := range sD.BitcoinQuorum {
					if quorum.FpBtcPk == msg.CBD.FpBtcPkList[0].MarshalHex() {
						var staker = Staker{
							StakingAmount: uint64(msg.CBD.StakingValue),
							Address:       msg.CBD.StakerAddr,
						}
						quorum.Staker = append(quorum.Staker, staker)
					}
				}
			}
		}

		bsD, err := json.Marshal(sD)
		if err != nil {
			return err
		}
		return s.db.Put(getStakeDetailsKey(), bsD, nil)
	} else if stakeType == UndelegateType {
		stakeDB, err := s.db.Get(getStakeDetailsKey(), nil)
		if err != nil {
			return err
		}
		var sD StakeDetails
		if err = json.Unmarshal(stakeDB, &sD); err != nil {
			return err
		}
		for _, quorum := range sD.BitcoinQuorum {
			if quorum.FpBtcPk == msg.CBD.FpBtcPkList[0].MarshalHex() {
				for _, staker := range quorum.Staker {
					staker.StakingAmount -= uint64(msg.CBD.StakingValue)
				}
			}
		}
		bsD, err := json.Marshal(sD)
		if err != nil {
			return err
		}
		return s.db.Put(getStakeDetailsKey(), bsD, nil)
	} else {
		return errors.New("unknown stake type")
	}
}

func (s *Storage) GetStakeDetails() (*StakeDetails, error) {
	sDB, err := s.db.Get(getStakeDetailsKey(), nil)
	if err != nil {
		return handleError(&StakeDetails{}, err)
	}
	var sD StakeDetails
	if err = json.Unmarshal(sDB, &sD); err != nil {
		return &StakeDetails{}, err
	}
	return &sD, nil
}

func (s *Storage) SetStakeDetailsByTimestamp(timestamp uint64) error {
	sDB, err := s.db.Get(getStakeDetailsKey(), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil
		} else {
			return err
		}
	}
	return s.db.Put(getStakeDetailsByTimestampKey(timestamp), sDB, nil)
}

func (s *Storage) GetStakeDetailsByTimestamp(start uint64, end uint64) (*StakeDetails, error) {
	iter := s.db.NewIterator(&util.Range{Start: getStakeDetailsByTimestampKey(start), Limit: getStakeDetailsByTimestampKey(end)}, nil)
	defer iter.Release()

	for iter.Next() {
		var sd StakeDetails
		if err := json.Unmarshal(iter.Value(), &sd); err != nil {
			return nil, err
		}
		return &sd, nil
	}

	if err := iter.Error(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Storage) DeleteStakeDetailsByTimestamp(start uint64, end uint64) error {
	iter := s.db.NewIterator(&util.Range{Start: getStakeDetailsByTimestampKey(start), Limit: getStakeDetailsByTimestampKey(end)}, nil)
	defer iter.Release()

	for iter.Next() {
		if err := s.db.Delete(iter.Key(), nil); err != nil {
			return err
		}
	}

	if err := iter.Error(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) SetBatchStakeDetails(batchID uint64, fpSignCache map[string]string, vs *types.VoteStateRoot, symbioticFpSignCache []string, start uint64, end uint64) error {
	var sD *StakeDetails
	sD, err := s.GetStakeDetailsByTimestamp(start, end)
	if err != nil {
		return err
	}

	if sD == nil {
		sD, err = s.GetStakeDetails()
		if err != nil {
			return err
		}
	}

	sD.BatchID = batchID
	sD.BabylonBlock = vs.BabylonHeight
	sD.StateRoot = vs.StateRoot
	sD.EthBlock = vs.L1BlockNumber
	sD.SymbioticSignNode = symbioticFpSignCache

	for fpPubkeyHex, sR := range fpSignCache {
		for i, quorum := range sD.BitcoinQuorum {
			if strings.ToLower(fpPubkeyHex) == strings.ToLower(quorum.FpBtcPk) && sR == vs.StateRoot {
				sD.BitcoinQuorum[i].IsSign = true
			}
		}
	}

	bsD, err := json.Marshal(sD)
	if err != nil {
		return err
	}
	fmt.Println("==============")
	fmt.Println(batchID)
	fmt.Println(sD)
	fmt.Println("=================")

	return s.db.Put(getBatchStakeDetailsKey(batchID), bsD, nil)
}

func (s *Storage) GetBatchStakeDetails(batchID uint64) (StakeDetails, error) {
	bSDB, err := s.db.Get(getBatchStakeDetailsKey(batchID), nil)
	if err != nil {
		return handleError(StakeDetails{}, err)
	}

	var sD StakeDetails
	if err = json.Unmarshal(bSDB, &sD); err != nil {
		return StakeDetails{}, err
	}
	sD.BatchID = batchID

	return sD, nil
}

func (s *Storage) GetBatchTotalBabylonStakeAmount(batchID uint64) (uint64, error) {
	bSDB, err := s.db.Get(getBatchStakeDetailsKey(batchID), nil)
	if err != nil {
		return 0, err
	}

	var sD StakeDetails
	if err = json.Unmarshal(bSDB, &sD); err != nil {
		return 0, err
	}

	var totalAmount uint64
	for _, quorum := range sD.BitcoinQuorum {
		if quorum.IsSign {
			for _, staker := range quorum.Staker {
				totalAmount += staker.StakingAmount
			}
		}
	}

	return totalAmount, nil
}
