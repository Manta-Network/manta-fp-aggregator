package manager

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/bindings/finality"
	sfp "github.com/Manta-Network/manta-fp-aggregator/bindings/sfp"
	"github.com/Manta-Network/manta-fp-aggregator/client"
	common2 "github.com/Manta-Network/manta-fp-aggregator/common"
	"github.com/Manta-Network/manta-fp-aggregator/common/httputil"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	kmssigner "github.com/Manta-Network/manta-fp-aggregator/kms"
	"github.com/Manta-Network/manta-fp-aggregator/manager/router"
	"github.com/Manta-Network/manta-fp-aggregator/manager/types"
	"github.com/Manta-Network/manta-fp-aggregator/metrics"
	"github.com/Manta-Network/manta-fp-aggregator/sign"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer"
	"github.com/Manta-Network/manta-fp-aggregator/ws/server"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	types2 "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/ethereum-optimism/optimism/op-proposer/bindings"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	errNotEnoughSignNode = errors.New("not enough available nodes to sign")
	errNotEnoughVoteNode = errors.New("not enough available nodes to vote")
)

type Manager struct {
	wg                       sync.WaitGroup
	done                     chan struct{}
	log                      log.Logger
	db                       *store.Storage
	wsServer                 server.IWebsocketManager
	NodeMembers              []string
	cfg                      *config.Config
	httpServer               *http.Server
	mu                       sync.Mutex
	windowPeriodStartTime    uint64
	outputSubmissionInterval uint64
	tickerController         bool

	ctx     context.Context
	stopped atomic.Bool

	ethChainID   uint64
	privateKey   *ecdsa.PrivateKey
	from         common.Address
	ethClient    *ethclient.Client
	frmContract  *finality.FinalityRelayerManager
	msmContract  *sfp.MantaStakingMiddleware
	l2oo         *bindings.L2OutputOracle
	batchId      uint64
	isFirstBatch bool

	babylonSynchronizer  *synchronizer.BabylonSynchronizer
	ethSynchronizer      *synchronizer.EthSynchronizer
	ethEventProcess      *synchronizer.EthEventProcess
	celestiaSynchronizer *synchronizer.CelestiaSynchronizer

	txMsgChan         chan store.TxMessage
	contractEventChan chan store.ContractEvent

	metricsRegistry *prometheus.Registry
	metrics         metrics.Metricer
	balanceMetricer io.Closer
	metricsServer   *httputil.HTTPServer

	kmsId     string
	kmsRegion string
	kmsClient *kms.Client
}

func NewFinalityManager(ctx context.Context, db *store.Storage, wsServer server.IWebsocketManager, cfg *config.Config, shutdown context.CancelCauseFunc, logger log.Logger, priv *ecdsa.PrivateKey, authToken string, kmsId string, kmsRegion string) (*Manager, error) {
	var kmsClient *kms.Client
	var from common.Address
	var pubkey *ecdsa.PublicKey
	var err error

	if cfg.EnableKms {
		kmsClient, err = kmssigner.NewKmsClientFromConfig(context.Background(), kmsRegion)
		if err != nil {
			return nil, fmt.Errorf("failed to create the kms client: %w", err)
		}
		pubkey, err = kmssigner.GetPubKeyCtx(ctx, kmsClient, kmsId)
		if err != nil {
			return nil, fmt.Errorf("failed to GetPubKey by kms: %w", err)
		}
	} else {
		pubkey = &priv.PublicKey
	}
	from = crypto.PubkeyToAddress(*pubkey)

	ethCli, err := client.DialEthClientWithTimeout(ctx, cfg.EthRpc, false)
	if err != nil {
		return nil, err
	}
	frmContract, err := finality.NewFinalityRelayerManager(common.HexToAddress(cfg.Contracts.FrmContractAddress), ethCli)
	if err != nil {
		return nil, err
	}
	msmContract, err := sfp.NewMantaStakingMiddleware(common.HexToAddress(cfg.Contracts.MsmContractAddress), ethCli)
	if err != nil {
		return nil, err
	}

	l2ooContract, err := bindings.NewL2OutputOracle(common.HexToAddress(cfg.Contracts.L2ooContractAddress), ethCli)
	if err != nil {
		return nil, err
	}

	latestBlock, err := ethCli.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block, err: %v", err)
	}

	cOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestBlock)),
		From:        from,
	}

	batchId, err := frmContract.ConfirmBatchId(cOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get batchId from fp contract, err: %v", err)
	}

	nodeMemberS := strings.Split(cfg.Manager.NodeMembers, ",")
	for _, nodeMember := range nodeMemberS {
		if err := db.SetActiveMember(nodeMember); err != nil {
			return nil, fmt.Errorf("failed to set node member, err: %v", err)
		}
	}

	registry := metrics.NewRegistry()
	metricer := metrics.NewMetrics(registry)

	txMsgChan := make(chan store.TxMessage, cfg.Manager.MaxBabylonOperatorNum)
	babylonSynchronizer, err := synchronizer.NewBabylonSynchronizer(ctx, cfg, db, shutdown, logger, txMsgChan, metricer)
	if err != nil {
		return nil, err
	}

	contractEventChan := make(chan store.ContractEvent, 100)
	ethSynchronizer, err := synchronizer.NewEthSynchronizer(cfg, db, ctx, logger, shutdown, contractEventChan, metricer)
	if err != nil {
		return nil, err
	}
	ethEventProcess, err := synchronizer.NewEthEventProcess(db, logger, contractEventChan)
	if err != nil {
		return nil, err
	}
	celestiaSynchronizer, err := synchronizer.NewCelestiaSynchronizer(ctx, cfg, db, shutdown, logger, authToken, metricer)
	if err != nil {
		return nil, err
	}

	return &Manager{
		done:                     make(chan struct{}),
		log:                      logger,
		db:                       db,
		wsServer:                 wsServer,
		NodeMembers:              nodeMemberS,
		cfg:                      cfg,
		ctx:                      ctx,
		privateKey:               priv,
		from:                     from,
		tickerController:         true,
		babylonSynchronizer:      babylonSynchronizer,
		ethSynchronizer:          ethSynchronizer,
		ethEventProcess:          ethEventProcess,
		celestiaSynchronizer:     celestiaSynchronizer,
		txMsgChan:                txMsgChan,
		contractEventChan:        contractEventChan,
		ethChainID:               cfg.EthChainID,
		ethClient:                ethCli,
		frmContract:              frmContract,
		msmContract:              msmContract,
		batchId:                  batchId.Uint64(),
		l2oo:                     l2ooContract,
		outputSubmissionInterval: uint64(cfg.Manager.OutputSubmitInterval),
		metricsRegistry:          registry,
		metrics:                  metricer,
		kmsId:                    kmsId,
		kmsRegion:                kmsRegion,
		kmsClient:                kmsClient,
	}, nil
}

func (m *Manager) Start(ctx context.Context) error {
	err := m.db.DeleteUnusedMembers(m.NodeMembers)
	if err != nil {
		m.log.Error("failed to delete unused members")
		return err
	}

	registry := router.NewRegistry(m, m.db)
	r := gin.Default()
	registry.Register(r)

	var s *http.Server
	s = &http.Server{
		Addr:    m.cfg.Manager.HttpAddr,
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			m.log.Error("api server starts failed", "err", err)
		}
	}()
	m.httpServer = s
	waitNodeTicker := time.NewTicker(5 * time.Second)
	defer waitNodeTicker.Stop()
	var done bool
	for !done {
		select {
		case <-waitNodeTicker.C:
			availableNodes := m.availableNodes(m.NodeMembers)
			if len(availableNodes) < len(m.NodeMembers) {
				m.log.Warn("wait node to connect", "availableNodesNum", len(availableNodes), "connectedNodeNum", len(m.NodeMembers))
				continue
			} else {
				done = true
				break
			}
		}
	}

	if m.batchId == 0 {
		m.isFirstBatch = true
	}

	if err := m.getWindowPeriodStartTime(); err != nil {
		m.log.Error("could not get window period start time", "err", err)
		return err
	}

	m.balanceMetricer = m.metrics.StartBalanceMetrics(m.log, m.ethClient, m.from)
	if err := m.startMetricsServer(); err != nil {
		m.log.Error("failed to start metrics Server", "err", err)
		return err
	}

	go m.babylonSynchronizer.Start()
	go m.ethSynchronizer.Start()
	go m.ethEventProcess.Start()
	go m.celestiaSynchronizer.Start()

	m.wg.Add(1)
	go m.work()
	m.log.Info("manager is starting......", "address", m.from.String())
	return nil
}

func (m *Manager) Stop(ctx context.Context) error {
	close(m.done)
	m.wg.Wait()
	if err := m.httpServer.Shutdown(ctx); err != nil {
		m.log.Error("http server forced to shutdown", "err", err)
		return err
	}
	if err := m.babylonSynchronizer.Close(); err != nil {
		m.log.Error("babylon synchronizer server forced to shutdown", "err", err)
		return err
	}
	m.ethSynchronizer.Close()
	m.celestiaSynchronizer.Close()
	if err := m.db.Close(); err != nil {
		m.log.Error("failed to close db server", "err", err)
		return err
	}
	if m.metricsServer != nil {
		if err := m.metricsServer.Close(); err != nil {
			m.log.Error("failed to close metrics server", "err", err)
		}
	}
	m.stopped.Store(true)
	m.log.Info("Server exiting")
	return nil
}

func (m *Manager) Stopped() bool {
	return m.stopped.Load()
}

func (m *Manager) getWindowPeriodStartTime() error {
	sR, err := m.db.GetLatestProcessedStateRoot()
	if err != nil {
		m.log.Error("failed to get latest processed state root", "err", err)
		return err
	}

	if sR != nil {
		m.windowPeriodStartTime = sR.Timestamp.Uint64()
	} else {
		latestHeight, err := m.ethClient.BlockNumber(m.ctx)
		if err != nil {
			m.log.Error("failed to get eth latest block height", "err", err)
			return err
		}

		cOpts := &bind.CallOpts{
			BlockNumber: big.NewInt(int64(latestHeight)),
			From:        m.from,
		}

		index, err := m.l2oo.LatestOutputIndex(cOpts)
		if err != nil {
			m.log.Error("failed to get latest output index", "height", latestHeight, "err", err)
			return err
		}

		output, err := m.l2oo.GetL2Output(cOpts, index)
		if err != nil {
			m.log.Error("failed to get l2 output", "height", latestHeight, "err", err)
			return err
		}
		m.windowPeriodStartTime = output.Timestamp.Uint64()
	}

	return nil
}

func (m *Manager) work() {
	fpTicker := time.NewTicker(m.cfg.Manager.FPTimeout)
	defer fpTicker.Stop()
	defer m.wg.Done()

	for {
		select {
		case txMsg := <-m.txMsgChan:
			go func(txMsg store.TxMessage) {
				if err := m.storeDelegateMsgData(txMsg); err != nil {
					m.log.Error("failed to store delegate msg", "err", err)
				}
			}(txMsg)
		case <-fpTicker.C:
			if !m.tickerController {
				m.log.Warn("the previous state root has not been processed yet")
				continue
			}
			opCtx, opCancel := context.WithTimeout(m.ctx, time.Duration(m.outputSubmissionInterval)*time.Second)
			defer opCancel()

			op, err := m.db.GetLatestUnprocessedStateRoot(m.windowPeriodStartTime, m.outputSubmissionInterval)
			if err != nil {
				m.log.Error("failed to get latest unprocessed state root", "err", err)
				continue
			}

			if op == nil {
				if m.ethSynchronizer.HeaderTraversal.LastTraversedHeader().Time > m.windowPeriodStartTime+m.outputSubmissionInterval {
					m.windowPeriodStartTime = m.windowPeriodStartTime + m.outputSubmissionInterval - 1
					m.log.Warn("no more state root need to processed, skip", "next_start", m.windowPeriodStartTime)
				}
				m.log.Warn("no more state root need to processed", "start", m.windowPeriodStartTime, "end", m.windowPeriodStartTime+m.outputSubmissionInterval)
				continue
			}

			if !m.checkSyncStatus(op) {
				continue
			}

			if m.isFirstBatch {
				m.windowPeriodStartTime = op.Timestamp.Uint64()
				if err = m.db.SetLatestProcessedStateRoot(*op); err != nil {
					m.log.Error("failed to set latest processed state root", "err", err)
					continue
				}
				m.isFirstBatch = false
				continue
			}

			m.tickerController = false
			m.metrics.RecordWindowPeriodStartTime(m.windowPeriodStartTime)
			m.metrics.RecordWindowPeriodEndTime(op.Timestamp.Uint64())

			done := make(chan struct{}, 1)
			errCh := make(chan error, 1)

			go func() {
				defer func() {
					close(done)
					opCancel()
				}()

				if err = m.processStateRoot(op); err != nil {
					errCh <- err
					return
				}
				m.resetState(op)
				done <- struct{}{}
			}()

			select {
			case <-done:
				m.log.Info("success to process state root", "batch_id", m.batchId)
				continue
			case err := <-errCh:
				m.log.Error("failed to process state root", "err", err)
				m.resetState(op)
				continue
			case <-opCtx.Done():
				m.log.Warn("process state root timeout, skip", "state_root", op.StateRoot)
				m.resetState(op)
				continue
			}

		case <-m.done:
			return
		}
	}
}

func (m *Manager) checkSyncStatus(op *store.OutputProposed) bool {
	babylonSynced := uint64(m.babylonSynchronizer.HeaderTraversal.LastTraversedHeader().Time.Unix()) >= op.Timestamp.Uint64()
	celestiaSynced := uint64(m.celestiaSynchronizer.HeaderTraversal.LastTraversedHeader().Time().Unix()) >= op.Timestamp.Uint64()

	if !babylonSynced {
		m.log.Warn("Babylon sync not completed",
			"required", op.Timestamp.Uint64(),
			"current", m.babylonSynchronizer.HeaderTraversal.LastTraversedHeader().Time.Unix())
		return false
	}

	if !celestiaSynced {
		m.log.Warn("Celestia sync not completed",
			"required", op.Timestamp.Uint64(),
			"current", m.celestiaSynchronizer.HeaderTraversal.LastTraversedHeader().Time().Unix())
		return false
	}
	return true
}

func (m *Manager) resetState(op *store.OutputProposed) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if err := m.db.SetLatestProcessedStateRoot(*op); err != nil {
		m.log.Error("failed to set latest processed state root")
	}
	m.windowPeriodStartTime = op.Timestamp.Uint64()
	m.tickerController = true

}

func (m *Manager) processStateRoot(op *store.OutputProposed) error {
	voteStateRoot, err := m.getMaxSignStateRoot(op.Timestamp.Uint64())
	m.log.Info("success to count fp signatures", "result", voteStateRoot)
	if err != nil {
		m.log.Error("failed to get max sign state root", "err", err)
		return err
	}

	var signature *sign.G1Point
	var g2Point *sign.G2Point
	var NonSignerPubkeys []finality.BN254G1Point
	var request types.SignMsgRequest

	if voteStateRoot.StateRoot != "" {
		request.StartTimestamp = voteStateRoot.StartTimestamp
		request.EndTimestamp = voteStateRoot.EndTimestamp
		request.L1BlockNumber = voteStateRoot.L1BlockNumber
		request.L2BlockNumber = voteStateRoot.L2BlockNumber
		request.StateRoot = voteStateRoot.StateRoot
	} else {
		m.log.Warn("no fp signature, skip this state root", "state_root", op.StateRoot)
		return errors.New("no fp signature, skip this state root")
	}

	res, err := m.SignMsgBatch(request)
	if errors.Is(err, errNotEnoughSignNode) || errors.Is(err, errNotEnoughVoteNode) {
		m.log.Error("not enough available nodes to sign or not enough available nodes to vote")
		return err
	} else if err != nil {
		m.log.Error("failed to sign msg", "err", err)
		return err
	}
	m.log.Info("success to sign msg", "signature", res.Signature)

	signature = res.Signature
	g2Point = res.G2Point
	for _, v := range res.NonSignerPubkeys {
		NonSignerPubkeys = append(NonSignerPubkeys, finality.BN254G1Point{
			X: v.X.BigInt(new(big.Int)),
			Y: v.Y.BigInt(new(big.Int)),
		})
	}
	err = m.getLatestConfirmBatchId()
	if err != nil {
		m.log.Error("failed to get latest confirm batch id", "err", err)
		return err
	}
	m.log.Info("success to get latest confirm batchId")

	err = m.db.SetBatchStakeDetails(m.batchId, voteStateRoot, m.windowPeriodStartTime, op.Timestamp.Uint64())
	if err != nil {
		m.log.Error("failed to store batch stake details", "err", err)
		return err
	}

	var opts *bind.TransactOpts
	if m.cfg.EnableKms {
		opts, err = kmssigner.NewAwsKmsTransactorWithChainIDCtx(context.Background(), m.kmsClient,
			m.kmsId, big.NewInt(int64(m.ethChainID)))
		if err != nil {
			m.log.Error("failed to new transact opts by kms", "err", err)
			return err
		}
	} else {
		opts, err = client.NewTransactOpts(m.ethChainID, m.privateKey)
		if err != nil {
			m.log.Error("failed to new transact opts", "err", err)
			return err
		}
	}
	opts.Context = context.Background()

	finalityBatch := finality.IFinalityRelayerManagerFinalityBatch{
		StateRoot:     common.HexToHash(voteStateRoot.StateRoot),
		L2BlockNumber: big.NewInt(int64(voteStateRoot.L2BlockNumber)),
		L1BlockHash:   common.HexToHash(voteStateRoot.L1BlockHash),
		L1BlockNumber: big.NewInt(int64(voteStateRoot.L1BlockNumber)),
		MsgHash:       crypto.Keccak256Hash(common.Hex2Bytes(voteStateRoot.StateRoot)),
	}

	totalBtcStake, err := m.db.GetBatchTotalBabylonStakeAmount(m.batchId)
	if err != nil {
		m.log.Error("failed to get total btc stake", "err", err)
		return err
	}

	finalityNonSignerAndSignature := finality.IBLSApkRegistryFinalityNonSignerAndSignature{
		NonSignerPubkeys: NonSignerPubkeys,
		ApkG2: finality.BN254G2Point{
			X: [2]*big.Int{g2Point.X.A1.BigInt(new(big.Int)), g2Point.X.A0.BigInt(new(big.Int))},
			Y: [2]*big.Int{g2Point.Y.A1.BigInt(new(big.Int)), g2Point.Y.A0.BigInt(new(big.Int))},
		},
		Sigma: finality.BN254G1Point{
			X: signature.X.BigInt(new(big.Int)),
			Y: signature.Y.BigInt(new(big.Int)),
		},
		TotalBtcStake:   big.NewInt(int64(totalBtcStake)),
		TotalMantaStake: voteStateRoot.TotalSymbioticFpStaked,
	}

	signatureIsValid, err := sign.VerifySig(signature.G1Affine, g2Point.G2Affine, crypto.Keccak256Hash(common.Hex2Bytes(voteStateRoot.StateRoot)))
	if err != nil {
		m.log.Error("failed to check signature is valid", "err", err)
		return err
	}
	m.log.Info("signature", "is", signatureIsValid)

	tx, err := m.frmContract.VerifyFinalitySignature(opts, finalityBatch, finalityNonSignerAndSignature, big.NewInt(100000))
	if err != nil {
		m.log.Error("failed to send VerifyFinalitySignature transaction", "err", err)
		return err
	}

	receipt, err := client.GetTransactionReceipt(context.Background(), m.ethClient, tx, time.Second*10, m.log)
	if err != nil {
		m.log.Error("failed to get verify finality transaction receipt", "err", err)
		m.metrics.RecordGetReceiptError(tx.Hash().String())
		return err
	}

	m.log.Info("success to send verify finality signature transaction", "tx_hash", receipt.TxHash.String())

	if err = m.db.DeleteStakeDetailsByTimestamp(m.babylonSynchronizer.StartTimestamp, m.windowPeriodStartTime); err != nil {
		m.log.Error("failed to delete old stake details data", "err", err)
		return err
	}

	if err = m.db.ChangeBatchStakeDetailsStatus(m.batchId, store.Confirmed); err != nil {
		m.log.Error("failed to change batch stake details status", "err", err)
		return err
	}

	m.metrics.RecordBatchId(m.batchId)

	return nil
}

func (m *Manager) SignMsgBatch(request types.SignMsgRequest) (*types.SignResult, error) {
	m.log.Info("received sign request", "sign_type", request)

	activeMember, err := m.db.GetActiveMember()
	if err != nil {
		m.log.Error("failed to get active member from db", "err", err)
		return nil, err
	}
	availableNodes := m.availableNodes(activeMember.Members)
	if len(availableNodes) < len(activeMember.Members) {
		m.log.Warn("not enough sign node", "availableNodes", availableNodes)
		return nil, errNotEnoughSignNode
	}

	ctx := types.NewContext().
		WithAvailableNodes(availableNodes).
		WithRequestId(randomRequestId())

	var resp types.SignResult
	var signErr error
	resp, signErr = m.sign(ctx, request, types.SignMsgBatch)
	if signErr != nil {
		return nil, signErr
	}
	if resp.Signature == nil {
		return nil, errNotEnoughVoteNode
	}

	return &resp, nil
}

func (m *Manager) availableNodes(nodeMembers []string) []string {
	aliveNodes := m.wsServer.AliveNodes()
	log.Info("check available nodes", "expected", fmt.Sprintf("%v", nodeMembers), "alive nodes", fmt.Sprintf("%v", aliveNodes))

	availableNodes := make([]string, 0)
	for _, n := range aliveNodes {
		if ExistsIgnoreCase(nodeMembers, n) {
			availableNodes = append(availableNodes, n)
		}
	}
	return availableNodes
}

func (m *Manager) storeDelegateMsgData(txMsg store.TxMessage) error {
	switch txMsg.Type {
	case common2.MsgCreateBTCDelegation:
		var mCBD types2.MsgCreateBTCDelegation
		mCBD.Unmarshal(txMsg.Data)
		if err := m.db.SetCreateBTCDelegationMsg(store.CreateBTCDelegation{
			CBD:    mCBD,
			TxHash: txMsg.TransactionHash,
		}); err != nil {
			return err
		}
		btcTx, err := types2.NewBtcTransaction(mCBD.StakingTx)
		if err != nil {
			m.log.Error("failed to new btc transaction", "err", err)
			return err
		}
		if err = m.db.SetBabylonDelegationKey(txMsg.TransactionHash, []byte(btcTx.Transaction.TxHash().String())); err != nil {
			m.log.Error("failed to store babylon delegation key", "err", err)
			return err
		}
		m.log.Info("success to store babylon delegation msg", "tx", hexutil.Encode(txMsg.TransactionHash))
		return nil
	case common2.MsgBTCUndelegate:
		var mBU types2.MsgBTCUndelegate
		mBU.Unmarshal(txMsg.Data)
		if err := m.db.SetBtcUndelegateMsg(store.BtcUndelegate{
			BU:     mBU,
			TxHash: txMsg.TransactionHash,
		}); err != nil {
			return err
		}
		m.log.Info("success to store babylon undelegation msg", "tx", hexutil.Encode(txMsg.TransactionHash))
		return nil
	case common2.MsgCreateFinalityProvider:
		return nil
	case common2.MsgCommitPubRandList:
		return nil
	case common2.MsgSelectiveSlashingEvidence:
		return nil
	case common2.MsgSubmitFinalitySignatureType:
		return nil
	default:
		return errors.New("unknown babylon tx msg type")
	}
}

func randomRequestId() string {
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	return time.Now().Format("20060102150405") + code
}

func ExistsIgnoreCase(slice []string, target string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, target) {
			return true
		}
	}
	return false
}

func (m *Manager) getMaxSignStateRoot(end uint64) (*types.VoteStateRoot, error) {
	stateRootCountCache := make(map[string]int64)
	babylonFpSignCache := make(map[string]string)
	symbioticFpSignCache := make(map[string]string)
	var symbioticFpSignList []string
	var totalSymbioticFpStaked = big.NewInt(0)

	op, err := m.db.GetLatestProcessedStateRoot()
	if err != nil {
		return nil, err
	}
	m.log.Info("start counting fp signatures", "start", op.Timestamp.Uint64(), "end", end, "stateroot", op.StateRoot)

	babylonFinalitySignatures, err := m.db.GetBabylonFinalitySignatureByTimestamp(op.Timestamp.Uint64(), end)
	if err != nil {
		return nil, err
	}
	// celestia block may have an earlier timestamp than eth block
	symbioticFinalitySignatures, err := m.db.GetSymbioticFpBlobsByTimestamp(op.Timestamp.Uint64()-uint64(time.Minute.Seconds()), end)
	if err != nil {
		return nil, err
	}

	for _, bfs := range babylonFinalitySignatures {
		if bfs.SubmitFinalitySignature.L1BlockNumber == op.L1BlockNumber && bfs.SubmitFinalitySignature.Height == op.L2BlockNumber.Uint64() {
			stateRootCountCache[bfs.SubmitFinalitySignature.StateRoot]++
			if babylonFpSignCache[bfs.SubmitFinalitySignature.FpPubkeyHex] == "" {
				babylonFpSignCache[bfs.SubmitFinalitySignature.FpPubkeyHex] = bfs.SubmitFinalitySignature.StateRoot
			}
		}
	}

	for _, sfs := range symbioticFinalitySignatures {
		if sfs.SignRequests.L1BlockNumber == op.L1BlockNumber && sfs.SignRequests.L2BlockNumber == op.L2BlockNumber.Uint64() {
			cOpts := &bind.CallOpts{
				BlockNumber: big.NewInt(int64(op.L1BlockNumber)),
				From:        m.from,
			}
			operator, err := m.msmContract.Operators(cOpts, common.HexToAddress(sfs.SignRequests.SignAddress))
			if err != nil {
				m.log.Error(fmt.Errorf("failed to get operator info at block: %v, err: %v", cOpts.BlockNumber, err).Error())
				continue
			}
			if operator.OperatorName == "" {
				m.log.Warn(fmt.Sprintf("node %s is not operator", sfs.SignRequests.SignAddress))
				continue
			} else {
				if operator.Paused {
					m.log.Warn("operator is paused", "address", sfs.SignRequests.SignAddress)
					continue
				}
			}
			if symbioticFpSignCache[sfs.SignRequests.SignAddress] == "" {
				vault, err := sfp.NewSymbioticVault(operator.Vault, m.ethClient)
				if err != nil {
					m.log.Error("failed to get operator vault", "address", sfs.SignRequests.SignAddress, "err", err)
					continue
				}
				amount, err := vault.ActiveStakeAt(cOpts, op.Timestamp, nil)
				if err != nil {
					m.log.Error("failed to get operator stake amount", "address", sfs.SignRequests.SignAddress, "err", err)
					continue
				}
				minMantaStakeAmount, _ := new(big.Int).SetString(m.cfg.MinMantaStakeAmount, 10)
				if amount.Cmp(minMantaStakeAmount) >= 0 {
					symbioticFpSignList = append(symbioticFpSignList, sfs.SignRequests.SignAddress)
					totalSymbioticFpStaked = new(big.Int).Add(totalSymbioticFpStaked, amount)
					stateRootCountCache[sfs.SignRequests.StateRoot]++
					symbioticFpSignCache[sfs.SignRequests.SignAddress] = sfs.SignRequests.StateRoot
				}
			}
		}
	}

	var maxSignStateRoot string
	var maxStateRootCount int64

	for k, v := range stateRootCountCache {
		if v > maxStateRootCount || maxSignStateRoot == "" {
			maxSignStateRoot = k
			maxStateRootCount = v
		}
	}

	var voteStateRoot = types.VoteStateRoot{
		StartTimestamp:         op.Timestamp.Uint64(),
		EndTimestamp:           end,
		BabylonHeight:          uint64(m.babylonSynchronizer.HeaderTraversal.LastTraversedHeader().Height),
		L1BlockNumber:          op.L1BlockNumber,
		L1BlockHash:            op.L1BlockHash.String(),
		L2BlockNumber:          op.L2BlockNumber.Uint64(),
		StateRoot:              maxSignStateRoot,
		BabylonFpSignCache:     babylonFpSignCache,
		SymbioticFpSignList:    symbioticFpSignList,
		TotalSymbioticFpStaked: totalSymbioticFpStaked,
	}

	return &voteStateRoot, nil
}

func (m *Manager) getLatestConfirmBatchId() error {
	latestBlock, err := m.ethClient.BlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get latest block, err: %v", err)
	}

	cOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestBlock)),
		From:        m.from,
	}

	id, err := m.frmContract.ConfirmBatchId(cOpts)
	if err != nil {
		return err
	}

	m.batchId = id.Uint64()

	return nil
}

func (m *Manager) startMetricsServer() error {
	m.log.Info("starting metrics server", "addr", m.cfg.Metrics.ListenAddr, "port", m.cfg.Metrics.ListenPort)
	metricsSrv, err := metrics.StartServer(m.metricsRegistry, m.cfg.Metrics.ListenAddr, m.cfg.Metrics.ListenPort)
	if err != nil {
		return fmt.Errorf("failed to start metrics server: %w", err)
	}
	m.log.Info("started metrics server", "addr", metricsSrv.Addr())
	m.metricsServer = metricsSrv
	return nil
}
