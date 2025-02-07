package synchronizer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/common/tasks"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer/node"

	"github.com/cometbft/cometbft/rpc/client/http"
	types2 "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"
	cTx "github.com/cosmos/cosmos-sdk/types/tx"
)

var validMsgTypes = map[string]bool{
	"/babylon.btcstaking.v1.MsgCreateFinalityProvider":    true,
	"/babylon.btcstaking.v1.MsgCreateBTCDelegation":       true,
	"/babylon.finality.v1.MsgCommitPubRandList":           true,
	"/babylon.btcstaking.v1.MsgBTCUndelegate":             true,
	"/babylon.btcstaking.v1.MsgSelectiveSlashingEvidence": true,
}

var ContractMsgType = "/cosmwasm.wasm.v1.MsgExecuteContract"

type BabylonSynchronizer struct {
	client               *http.HTTP
	db                   *store.Storage
	headers              []types2.Header
	latestHeader         *types2.Header
	headerTraversal      *node.BabylonHeaderTraversal
	opFinalityGadgetAddr string
	blockStep            uint64
	startHeight          *big.Int
	confirmationDepth    *big.Int
	resourceCtx          context.Context
	resourceCancel       context.CancelFunc
	tasks                tasks.Group
	log                  log.Logger
	txMsgChan            chan store.TxMessage
}

func NewBabylonSynchronizer(ctx context.Context, cfg *config.Config, db *store.Storage, shutdown context.CancelCauseFunc, logger log.Logger, txMsgChan chan store.TxMessage) (*BabylonSynchronizer, error) {

	cli, err := client.NewClientFromNode(cfg.BabylonRpc)
	if err != nil {
		fmt.Printf("Error creating client: %v", err)
	}

	dbLatestHeader, err := db.GetBabylonScannedHeight()
	if err != nil {
		return nil, err
	}
	var fromHeader *types2.Header
	if dbLatestHeader != 0 {
		logger.Info("sync detected last indexed block", "number", dbLatestHeader)
		height := int64(dbLatestHeader)
		block, err := cli.Block(ctx, &height)
		if err != nil {
			logger.Error("failed to get babylon block", "height", dbLatestHeader)
		}
		fromHeader = &block.Block.Header
	} else if cfg.BabylonStartingHeight > 0 {
		logger.Info("no sync indexed state starting from supplied babylon height", "height", cfg.BabylonStartingHeight)
		block, err := cli.Block(ctx, &cfg.BabylonStartingHeight)
		if err != nil {
			return nil, fmt.Errorf("could not fetch babylon starting block header: %w", err)
		}
		fromHeader = &block.Block.Header
	} else {
		logger.Info("no ethereum block indexed state")
	}

	headerTraversal := node.NewBabylonHeaderTraversal(cli, fromHeader, big.NewInt(0))

	resCtx, resCancel := context.WithCancel(context.Background())
	return &BabylonSynchronizer{
		client:               cli,
		blockStep:            cfg.BabylonBlockStep,
		headerTraversal:      headerTraversal,
		latestHeader:         fromHeader,
		db:                   db,
		resourceCtx:          resCtx,
		resourceCancel:       resCancel,
		opFinalityGadgetAddr: cfg.Contracts.OpFinalityGadgat,
		log:                  logger,
		txMsgChan:            txMsgChan,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in Synchronizer: %w", err))
		}},
	}, nil
}

func (syncer *BabylonSynchronizer) Start() error {
	tickerSyncer := time.NewTicker(time.Second * 2)
	syncer.tasks.Go(func() error {
		for range tickerSyncer.C {
			if len(syncer.headers) > 0 {
				syncer.log.Info("retrying previous batch")
			} else {
				newHeaders, err := syncer.headerTraversal.NextHeaders(syncer.blockStep)
				if err != nil {
					syncer.log.Error("error querying for headers", "err", err)
					continue
				} else if len(newHeaders) == 0 {
					syncer.log.Warn("no new headers. syncer at head?")
				} else {
					syncer.headers = newHeaders
				}
				latestHeader := syncer.headerTraversal.LatestHeader()
				if latestHeader != nil {
					syncer.log.Info("Latest header", "latestHeader Number", latestHeader.Height)
				}
			}
			err := syncer.processBatch(syncer.headers)
			if err == nil {
				syncer.headers = nil
			}
		}
		return nil
	})
	return nil
}

func (syncer *BabylonSynchronizer) processBatch(headers []types2.Header) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	syncer.log.Info("extracting batch", "size", len(headers), "startBlock", firstHeader.Height, "endBlock", lastHeader.Height)

	headerMap := make(map[int64]*types2.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Height] = &header
	}
	blockHeaders := make([]store.BabylonBlockHeader, 0, len(headers))
	var txMessages []store.TxMessage
	for i := range headers {
		if headers[i].Hash() == nil {
			continue
		}
		bHeader := store.BabylonBlockHeader{
			Hash:       headers[i].Hash(),
			ParentHash: headers[i].LastResultsHash.Bytes(),
			Number:     headers[i].Height,
			Timestamp:  headers[i].Time.Unix(),
		}
		blockHeaders = append(blockHeaders, bHeader)

		block, err := syncer.client.Block(syncer.resourceCtx, &headers[i].Height)
		if err != nil {
			syncer.log.Error("failed to get block", "err", err, "height", headers[i].Height)
			return err
		}
		for _, transaction := range block.Block.Txs {
			var tx cTx.Tx
			tx.Unmarshal(transaction)

			for _, msg := range tx.Body.Messages {
				if validMsgTypes[msg.TypeUrl] {
					if err != nil {
						syncer.log.Error("failed to marshal event", "err", err)
						continue
					}
					txMessage := store.TxMessage{
						BlockHeight:     uint64(block.Block.Height),
						TransactionHash: transaction.Hash(),
						Type:            msg.TypeUrl,
						Data:            msg.Value,
						Timestamp:       time.Now().Unix(),
					}
					txMessages = append(txMessages, txMessage)
					syncer.txMsgChan <- txMessage
				} else if msg.TypeUrl == ContractMsgType {
					var sMsg store.SubmitFinalitySignatureMsgValue
					sMsg.Unmarshal(msg.Value)
					if sMsg.Contract == syncer.opFinalityGadgetAddr && sMsg.StateRoot != nil {
						txMessage := store.TxMessage{
							BlockHeight:     uint64(block.Block.Height),
							TransactionHash: transaction.Hash(),
							Type:            "/MsgSubmitFinalitySignature",
							Data:            msg.Value,
							Timestamp:       time.Now().Unix(),
						}
						txMessages = append(txMessages, txMessage)
						syncer.txMsgChan <- txMessage
					}
				}
			}
		}
	}

	if err := syncer.db.SetBabylonBlockHeaders(blockHeaders); err != nil {
		return err
	}
	if err := syncer.db.SetTxMessages(txMessages); err != nil {
		return err
	}
	if err := syncer.db.UpdateBabylonHeight(uint64(lastHeader.Height)); err != nil {
		return err
	}

	return nil
}

func (syncer *BabylonSynchronizer) Close() error {
	return nil
}
