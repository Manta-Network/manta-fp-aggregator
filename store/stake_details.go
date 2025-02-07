package store

import (
	"encoding/json"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
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
	FpBtcPk      string   `json:"fpName"`
	FpVoteWeight uint64   `json:"fpVoteWeight"`
	Staker       []Staker `json:"staker"`
}

type StakeDetails struct {
	BatchID       uint64       `json:"batchId"`
	TotalBTCVote  uint64       `json:"tatolBtcVote"`
	BabylonBlock  uint64       `json:"babylonBlock"`
	BitcoinQuorum []QuorumNode `json:"BitcoinQuorum"`
}

func (s *Storage) SetStakeDetails(batchId uint64, blockHeight uint64, msg CreateBTCDelegation, stakeType int8) error {
	if stakeType == DelegateType {
		stakeDB, err := s.db.Get(getStakeDetailsKey(batchId), nil)
		if err != nil {
			if errors.Is(err, leveldb.ErrNotFound) {
				var sD = StakeDetails{
					BatchID:      batchId,
					TotalBTCVote: uint64(msg.CBD.StakingValue),
					BabylonBlock: blockHeight,
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
				return s.db.Put(getStakeDetailsKey(batchId), bz, nil)
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
		return s.db.Put(getStakeDetailsKey(batchId), bsD, nil)
	} else if stakeType == UndelegateType {
		stakeDB, err := s.db.Get(getStakeDetailsKey(batchId), nil)
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
		return s.db.Put(getStakeDetailsKey(batchId), bsD, nil)
	} else {
		return errors.New("unknown stake type")
	}
}

func (s *Storage) GetStakeDetails(batchId uint64) (StakeDetails, error) {
	sDB, err := s.db.Get(getStakeDetailsKey(batchId), nil)
	if err != nil {
		return handleError(StakeDetails{}, err)
	}
	var sD StakeDetails
	if err = json.Unmarshal(sDB, &sD); err != nil {
		return StakeDetails{}, err
	}
	return sD, nil
}
