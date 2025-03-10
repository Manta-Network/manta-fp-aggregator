package synchronizer

import (
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer/contracts"
)

type EthEventProcess struct {
	db  *store.Storage
	log log.Logger

	contractEventChan chan store.ContractEvent

	finalityRelayerManager *contracts.FinalityRelayerManager
	l2OutputOracle         *contracts.L2OutputOracle
}

func NewEthEventProcess(db *store.Storage, logger log.Logger, contractEventChan chan store.ContractEvent) (*EthEventProcess, error) {
	finalityRelayerManager, err := contracts.NewFinalityRelayerManager(logger)
	if err != nil {
		logger.Error("new finality relayer manager fail", "err", err)
		return nil, err
	}
	l2OutputOracle, err := contracts.NewL2OutputOracle(logger)
	if err != nil {
		logger.Error("new l2 output oracle fail", "err", err)
		return nil, err
	}

	return &EthEventProcess{
		db:                     db,
		log:                    logger,
		contractEventChan:      contractEventChan,
		finalityRelayerManager: finalityRelayerManager,
		l2OutputOracle:         l2OutputOracle,
	}, nil
}

func (e *EthEventProcess) Start() error {
	for {
		select {
		case event := <-e.contractEventChan:
			if err := e.finalityRelayerManager.ProcessFinalityRelayerManagerEvent(e.db, event); err != nil {
				e.log.Error("failed to process FinalityRelayerManager event", "err", err)
				continue
			}
			if err := e.l2OutputOracle.ProcessOutputProposedEvent(e.db, event); err != nil {
				e.log.Error("failed to process OutputPropose event", "err", err)
				continue
			}
		}
	}
}
