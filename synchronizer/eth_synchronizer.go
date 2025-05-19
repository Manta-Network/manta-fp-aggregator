package synchronizer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/common/tasks"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/Manta-Network/manta-fp-aggregator/metrics"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer/node"
	node2 "github.com/Manta-Network/manta-fp-aggregator/synchronizer/node"
)

type EthSynchronizer struct {
	ethClient         node2.EthClient
	db                *store.Storage
	headers           []types.Header
	latestHeader      *types.Header
	HeaderTraversal   *node.EthHeaderTraversal
	blockStep         uint64
	contracts         []common.Address
	startHeight       *big.Int
	confirmationDepth *big.Int
	resourceCtx       context.Context
	resourceCancel    context.CancelFunc
	contractEventChan chan store.ContractEvent
	log               log.Logger
	metrics           metrics.Metricer
	tasks             tasks.Group
}

func NewEthSynchronizer(cfg *config.Config, db *store.Storage, ctx context.Context, logger log.Logger, shutdown context.CancelCauseFunc, contractEventChan chan store.ContractEvent, metricer metrics.Metricer) (*EthSynchronizer, error) {
	client, err := node.DialEthClient(ctx, cfg.EthRpc)
	if err != nil {
		return nil, err
	}

	dbLatestHeader, err := db.GetEthScannedHeight()
	if err != nil {
		return nil, err
	}
	var fromHeader *types.Header
	if dbLatestHeader != 0 {
		logger.Info("eth: sync detected last indexed block", "number", dbLatestHeader)
		header, err := client.BlockHeaderByNumber(big.NewInt(int64(dbLatestHeader)))
		if err != nil {
			logger.Error("failed to get eth block header", "height", dbLatestHeader)
		}
		fromHeader = header
	} else if cfg.EthStartingHeight > 0 {
		logger.Info("eth: no sync indexed state starting from supplied ethereum height", "height", cfg.EthStartingHeight)
		header, err := client.BlockHeaderByNumber(big.NewInt(cfg.EthStartingHeight))
		if err != nil {
			return nil, fmt.Errorf("could not fetch eth starting block header: %w", err)
		}
		fromHeader = header
	} else {
		logger.Info("no ethereum block indexed state")
	}

	headerTraversal := node.NewEthHeaderTraversal(client, fromHeader, big.NewInt(0))

	var contracts []common.Address
	contracts = append(contracts, common.HexToAddress(cfg.Contracts.FrmContractAddress))
	contracts = append(contracts, common.HexToAddress(cfg.Contracts.L2ooContractAddress))

	resCtx, resCancel := context.WithCancel(context.Background())
	return &EthSynchronizer{
		HeaderTraversal:   headerTraversal,
		ethClient:         client,
		latestHeader:      fromHeader,
		db:                db,
		blockStep:         cfg.EthBlockStep,
		contracts:         contracts,
		resourceCtx:       resCtx,
		resourceCancel:    resCancel,
		log:               logger,
		contractEventChan: contractEventChan,
		metrics:           metricer,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in eth synchronizer: %w", err))
		}},
	}, nil
}

func (syncer *EthSynchronizer) Start() error {
	tickerSyncer := time.NewTicker(time.Second * 2)
	syncer.tasks.Go(func() error {
		for range tickerSyncer.C {
			done := syncer.metrics.RecordEthInterval()
			if len(syncer.headers) > 0 {
				syncer.log.Info("eth: retrying previous batch")
			} else {
				newHeaders, err := syncer.HeaderTraversal.NextHeaders(syncer.blockStep)
				if err != nil {
					syncer.log.Error("eth: error querying for headers", "err", err)
					continue
				} else if len(newHeaders) == 0 {
					syncer.log.Warn("eth: no new headers. syncer at head?")
				} else {
					syncer.headers = newHeaders
				}
				latestHeader := syncer.HeaderTraversal.LatestHeader()
				if latestHeader != nil {
					syncer.log.Info("eth: Latest header", "latestHeader Number", latestHeader.Number)
				}
			}
			err := syncer.processBatch(syncer.headers)
			if err == nil {
				syncer.headers = nil
			}
			done(err)
		}
		return nil
	})
	return nil
}

func (syncer *EthSynchronizer) processBatch(headers []types.Header) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	syncer.log.Info("eth: extracting batch", "size", len(headers), "startBlock", firstHeader.Number.String(), "endBlock", lastHeader.Number.String())

	headerMap := make(map[common.Hash]*types.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Hash()] = &header
	}
	syncer.log.Info("eth: chainCfg Contracts", "contract address", syncer.contracts)
	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number, Addresses: syncer.contracts}
	logs, err := syncer.ethClient.FilterLogs(filterQuery)
	if err != nil {
		syncer.log.Info("failed to extract logs", "err", err)
		return err
	}

	if logs.ToBlockHeader.Number.Cmp(lastHeader.Number) != 0 {
		return fmt.Errorf("eth: mismatch in FilterLog#ToBlock number")
	} else if logs.ToBlockHeader.Hash() != lastHeader.Hash() {
		return fmt.Errorf("eth: mismatch in FitlerLog#ToBlock block hash")
	}

	if len(logs.Logs) > 0 {
		syncer.log.Info("eth: detected logs", "size", len(logs.Logs))
	}

	blockHeaders := make([]store.EthBlockHeader, 0, len(headers))
	for i := range headers {
		if headers[i].Number == nil {
			continue
		}
		eHeader := store.EthBlockHeader{
			Hash:       headers[i].TxHash,
			ParentHash: headers[i].ParentHash,
			Number:     headers[i].Number.Int64(),
			Timestamp:  int64(headers[i].Time),
		}
		blockHeaders = append(blockHeaders, eHeader)
	}

	chainContractEvent := make([]store.ContractEvent, len(logs.Logs))
	for i := range logs.Logs {
		logEvent := logs.Logs[i]
		if _, ok := headerMap[logEvent.BlockHash]; !ok {
			continue
		}
		timestamp := headerMap[logEvent.BlockHash].Time
		chainContractEvent[i] = store.ContractEventFromLog(&logs.Logs[i], timestamp)
		syncer.contractEventChan <- chainContractEvent[i]
		syncer.metrics.RecordEthBatchLog(chainContractEvent[i].ContractAddress.String())
	}

	if err := syncer.db.SetEthBlockHeaders(blockHeaders); err != nil {
		return err
	}
	if err := syncer.db.SetContractEvents(chainContractEvent); err != nil {
		return err
	}
	if err := syncer.db.UpdateEthHeight(lastHeader.Number.Uint64()); err != nil {
		return err
	}

	syncer.metrics.RecordLatestEthBlock(lastHeader.Number.Uint64())
	syncer.metrics.RecordEthIndexedHeaders(len(blockHeaders))

	return nil
}
func (syncer *EthSynchronizer) Close() error {
	return nil
}
