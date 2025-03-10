package contracts

import (
	"context"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/store"

	"github.com/ethereum-optimism/optimism/op-proposer/bindings"
	"go.uber.org/zap"
)

type L2OutputOracle struct {
	L2ooFilterer *bindings.L2OutputOracleFilterer
	L2ooABI      *abi.ABI
	L2ooCtx      context.Context
	log          log.Logger
}

func NewL2OutputOracle(log log.Logger) (*L2OutputOracle, error) {
	l2OutputOracleAbi, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		log.Error("get l2 output oracle abi fail", zap.String("err", err.Error()))
		return nil, err
	}

	l2OutputOracleUnpack, err := bindings.NewL2OutputOracleFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new l2output oracle fail", zap.String("err", err.Error()))
		return nil, err
	}

	return &L2OutputOracle{
		log:          log,
		L2ooFilterer: l2OutputOracleUnpack,
		L2ooABI:      l2OutputOracleAbi,
		L2ooCtx:      context.Background(),
	}, nil
}

func (l2oo *L2OutputOracle) ProcessOutputProposedEvent(db *store.Storage, event store.ContractEvent) error {
	var l2OutputOracle *store.L2OutputOracle
	if event.EventSignature.String() == l2oo.L2ooABI.Events["OutputProposed"].ID.String() {
		outputProposedEvent, err := l2oo.L2ooFilterer.ParseOutputProposed(*event.RLPLog)
		if err != nil {
			l2oo.log.Error("parse outputProposedEvent fail", "err", err)
			return err
		}
		log.Info("parse outputProposedEvent success",
			"state_root", hex.EncodeToString(outputProposedEvent.OutputRoot[:]))

		l2OutputOracle = &store.L2OutputOracle{
			StateRoot:     hex.EncodeToString(outputProposedEvent.OutputRoot[:]),
			L2BlockNumber: outputProposedEvent.L2BlockNumber,
			L2OutputIndex: outputProposedEvent.L2OutputIndex,
			L1BlockHash:   outputProposedEvent.Raw.BlockHash,
			L1BlockNumber: outputProposedEvent.Raw.BlockNumber,
			Timestamp:     outputProposedEvent.L1Timestamp,
		}
	}

	if l2OutputOracle != nil {
		if err := db.SetL2OutputOracle(*l2OutputOracle); err != nil {
			return err
		}
		l2oo.log.Info("store outputProposedEvent success")
	}

	return nil
}
