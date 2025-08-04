package synchronizer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/common/tasks"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/Manta-Network/manta-fp-aggregator/metrics"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer/node"

	client "github.com/celestiaorg/celestia-openrpc"
	"github.com/celestiaorg/celestia-openrpc/types/header"
	"github.com/celestiaorg/celestia-openrpc/types/share"
)

type CelestiaSynchronizer struct {
	client          *client.Client
	db              *store.Storage
	headers         []*header.ExtendedHeader
	LatestHeader    *header.ExtendedHeader
	HeaderTraversal *node.CelestiaHeaderTraversal
	blockStep       uint64
	tasks           tasks.Group
	log             log.Logger
	namespace       share.Namespace
	metrics         metrics.Metricer
	tickerSyncer    *time.Ticker
}

func NewCelestiaSynchronizer(ctx context.Context, cfg *config.Config, db *store.Storage, shutdown context.CancelCauseFunc, logger log.Logger, authToken string, metricer metrics.Metricer) (*CelestiaSynchronizer, error) {
	cli, err := client.NewClient(ctx, cfg.CelestiaConfig.DaRpc, authToken)
	if err != nil {
		return nil, err
	}

	nsBytes, err := hex.DecodeString(cfg.CelestiaConfig.Namespace)
	if err != nil {
		return nil, err
	}
	if len(nsBytes) != 10 {
		return nil, errors.New("wrong namespace length")
	}

	namespace, err := share.NamespaceFromBytes(append(make([]byte, 19), nsBytes...))
	if err != nil {
		return nil, err
	}

	dbLatestHeader, err := db.GetCelestiaScannedHeight()
	if err != nil {
		return nil, err
	}
	var fromHeader *header.ExtendedHeader
	if dbLatestHeader != 0 {
		logger.Info("celestia: sync detected last indexed block", "number", dbLatestHeader)
		header, err := cli.Header.GetByHeight(ctx, dbLatestHeader)
		if err != nil {
			logger.Error("failed to get celestia header by latest db header", "height", dbLatestHeader)
			return nil, fmt.Errorf("could not fetch celestia starting block header by latest db header: %w", err)
		}
		fromHeader = header
	} else if cfg.CelestiaStartingHeight > 0 {
		logger.Info("celestia: no sync indexed state starting from supplied celestia height", "height", cfg.CelestiaStartingHeight)
		header, err := cli.Header.GetByHeight(ctx, uint64(cfg.CelestiaStartingHeight))
		if err != nil {
			logger.Error("failed to get celestia header by cfg height", "height", dbLatestHeader)
			return nil, fmt.Errorf("could not fetch celestia starting block header by cfg height: %w", err)
		}
		fromHeader = header
	} else {
		logger.Info("no celestia block indexed state")
	}

	// The probability of 1 blocks being reorganized is very low for celestia
	headerTraversal := node.NewCelestiaHeaderTraversal(cli, fromHeader, big.NewInt(1))

	return &CelestiaSynchronizer{
		client:          cli,
		blockStep:       cfg.CelestiaBlockStep,
		HeaderTraversal: headerTraversal,
		LatestHeader:    fromHeader,
		db:              db,
		log:             logger,
		namespace:       namespace,
		metrics:         metricer,
		tickerSyncer:    time.NewTicker(time.Second * 2),
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in celestia synchronizer: %w", err))
		}},
	}, nil
}

func (syncer *CelestiaSynchronizer) Start() error {
	syncer.tasks.Go(func() error {
		for range syncer.tickerSyncer.C {
			done := syncer.metrics.RecordCelestiaInterval()
			if len(syncer.headers) > 0 {
				syncer.log.Info("celestia: retrying previous batch")
			} else {
				newHeaders, err := syncer.HeaderTraversal.NextHeaders(syncer.blockStep)
				if err != nil {
					syncer.log.Error("celestia error querying for headers", "err", err)
					continue
				} else if len(newHeaders) == 0 {
					syncer.log.Warn("celestia: no new headers. syncer at head?")
				} else {
					syncer.headers = newHeaders
				}
				latestHeader := syncer.HeaderTraversal.LatestHeader()
				if latestHeader != nil {
					syncer.log.Info("celestia: Latest header", "latestHeader Number", latestHeader.Height())
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

func (syncer *CelestiaSynchronizer) processBatch(headers []*header.ExtendedHeader) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	syncer.log.Info("celestia: extracting batch", "size", len(headers), "startBlock", firstHeader.Height(), "endBlock", lastHeader.Height())

	headerMap := make(map[uint64]*header.ExtendedHeader, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Height()] = header
	}
	blockHeaders := make([]store.CelestiaBlockHeader, 0, len(headers))
	for i := range headers {
		if headers[i].Hash() == nil {
			continue
		}
		cHeader := store.CelestiaBlockHeader{
			Hash:       headers[i].Hash(),
			ParentHash: headers[i].LastHeader(),
			Number:     headers[i].Height(),
			Timestamp:  headers[i].Time().Unix(),
		}
		blockHeaders = append(blockHeaders, cHeader)

		blobs, err := syncer.client.Blob.GetAll(context.Background(), cHeader.Number, []share.Namespace{syncer.namespace})
		if err != nil {
			syncer.log.Error("celestia: failed to get blob", "height", cHeader.Number)
			continue
		}

		for _, blob := range blobs {
			var signRequest store.SignRequest
			err = json.Unmarshal(blob.Data, &signRequest)
			if err != nil {
				syncer.log.Warn("celestia: failed to unmarshal symbiotic fp blob data to sign request", "err", err)
				continue
			}
			if signRequest.StateRoot != "" && signRequest.SignAddress != "" && signRequest.Signature != nil {
				err = syncer.db.SetSymbioticFpBlob(store.SymbioticFpBlob{
					SignRequests: signRequest,
					Timestamp:    uint64(cHeader.Timestamp),
				})
				if err != nil {
					return err
				}
				syncer.log.Info("celestia: success to store symbiotic fp blob data", "height", cHeader.Number, "timestamp", cHeader.Timestamp)
			}
		}
		if err = syncer.db.SetStakeDetailsByTimestamp(uint64(headers[i].Time().Unix())); err != nil {
			return err
		}
	}

	if err := syncer.db.SetCelestiaBlockHeaders(blockHeaders); err != nil {
		return err
	}
	if err := syncer.db.UpdateCelestiaHeight(lastHeader.Height()); err != nil {
		return err
	}

	syncer.metrics.RecordLatestCelestiaBlock(lastHeader.Height())
	syncer.metrics.RecordCelestiaIndexedHeaders(len(blockHeaders))

	return nil
}

func (syncer *CelestiaSynchronizer) Close() {
	syncer.client.Close()
	syncer.tickerSyncer.Stop()
}
