package manager

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/bindings/bls"
	"github.com/Manta-Network/manta-fp-aggregator/bindings/finality"
	"github.com/Manta-Network/manta-fp-aggregator/bindings/sfp"
	"github.com/Manta-Network/manta-fp-aggregator/client"
	common2 "github.com/Manta-Network/manta-fp-aggregator/common"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/Manta-Network/manta-fp-aggregator/manager/celestia"
	"github.com/Manta-Network/manta-fp-aggregator/manager/router"
	"github.com/Manta-Network/manta-fp-aggregator/manager/types"
	"github.com/Manta-Network/manta-fp-aggregator/sign"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer"
	"github.com/Manta-Network/manta-fp-aggregator/ws/server"

	types4 "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	types2 "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/ethereum-optimism/optimism/op-proposer/bindings"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

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
	httpAddr                 string
	httpServer               *http.Server
	grpcAddr                 string
	gs                       *grpc.Server
	sStakeUrl                string
	mu                       sync.Mutex
	windowPeriodStartTime    uint64
	outputSubmissionInterval uint64

	ctx     context.Context
	stopped atomic.Bool

	ethChainID      uint64
	privateKey      *ecdsa.PrivateKey
	from            common.Address
	ethClient       *ethclient.Client
	celestiaDa      *celestia.DAClient
	frmContract     *finality.FinalityRelayerManager
	frmContractAddr common.Address
	rawFrmContract  *bind.BoundContract
	barContract     *bls.BLSApkRegistry
	barContractAddr common.Address
	rawBarContract  *bind.BoundContract
	msmContract     *sfp.MantaStakingMiddleware
	l2oo            *bindings.L2OutputOracle
	batchId         uint64
	isFirstBatch    bool

	signTimeout time.Duration
	fPTimeout   time.Duration

	babylonSynchronizer *synchronizer.BabylonSynchronizer
	ethSynchronizer     *synchronizer.EthSynchronizer
	ethEventProcess     *synchronizer.EthEventProcess

	txMsgChan         chan store.TxMessage
	contractEventChan chan store.ContractEvent
}

func NewFinalityManager(ctx context.Context, db *store.Storage, wsServer server.IWebsocketManager, cfg *config.Config, shutdown context.CancelCauseFunc, logger log.Logger, priv *ecdsa.PrivateKey) (*Manager, error) {
	ethCli, err := client.DialEthClientWithTimeout(ctx, cfg.EthRpc, false)
	if err != nil {
		return nil, err
	}
	frmContract, err := finality.NewFinalityRelayerManager(common.HexToAddress(cfg.Contracts.FrmContractAddress), ethCli)
	if err != nil {
		return nil, err
	}
	fParsed, err := abi.JSON(strings.NewReader(
		finality.FinalityRelayerManagerABI,
	))
	if err != nil {
		return nil, err
	}
	rawfrmContract := bind.NewBoundContract(
		common.HexToAddress(cfg.Contracts.FrmContractAddress), fParsed, ethCli, ethCli,
		ethCli,
	)

	barContract, err := bls.NewBLSApkRegistry(common.HexToAddress(cfg.Contracts.BarContactAddress), ethCli)
	if err != nil {
		return nil, err
	}
	bParsed, err := bls.BLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	rawBarContract := bind.NewBoundContract(
		common.HexToAddress(cfg.Contracts.BarContactAddress), *bParsed, ethCli, ethCli,
		ethCli,
	)

	//msmContract, err := sfp.NewMantaStakingMiddleware(common.HexToAddress(cfg.Contracts.MsmContractAddress), ethCli)
	//if err != nil {
	//	return nil, err
	//}

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
		From:        crypto.PubkeyToAddress(priv.PublicKey),
	}

	batchId, err := frmContract.ConfirmBatchId(cOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get batchId from fp contract, err: %v", err)
	}

	nodeMemberS := strings.Split(cfg.Manager.NodeMembers, ",")

	txMsgChan := make(chan store.TxMessage, cfg.Manager.MaxBabylonOperatorNum)
	babylonSynchronizer, err := synchronizer.NewBabylonSynchronizer(ctx, cfg, db, shutdown, logger, txMsgChan)
	if err != nil {
		return nil, err
	}

	contractEventChan := make(chan store.ContractEvent, 100)
	ethSynchronizer, err := synchronizer.NewEthSynchronizer(cfg, db, ctx, logger, shutdown, contractEventChan)
	if err != nil {
		return nil, err
	}
	ethEventProcess, err := synchronizer.NewEthEventProcess(db, logger, contractEventChan)
	if err != nil {
		return nil, err
	}
	celestiaDa, err := celestia.NewDAClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Manager{
		done:                make(chan struct{}),
		log:                 logger,
		db:                  db,
		wsServer:            wsServer,
		httpAddr:            cfg.Manager.HttpAddr,
		grpcAddr:            cfg.Manager.SymbioticFpRpc,
		sStakeUrl:           cfg.Manager.SymbioticStakeUrl,
		celestiaDa:          celestiaDa,
		NodeMembers:         nodeMemberS,
		ctx:                 ctx,
		privateKey:          priv,
		from:                crypto.PubkeyToAddress(priv.PublicKey),
		signTimeout:         cfg.Manager.SignTimeout,
		fPTimeout:           cfg.Manager.FPTimeout,
		babylonSynchronizer: babylonSynchronizer,
		ethSynchronizer:     ethSynchronizer,
		ethEventProcess:     ethEventProcess,
		txMsgChan:           txMsgChan,
		contractEventChan:   contractEventChan,
		ethChainID:          cfg.EthChainID,
		ethClient:           ethCli,
		frmContract:         frmContract,
		frmContractAddr:     common.HexToAddress(cfg.Contracts.FrmContractAddress),
		rawFrmContract:      rawfrmContract,
		barContract:         barContract,
		barContractAddr:     common.HexToAddress(cfg.Contracts.BarContactAddress),
		rawBarContract:      rawBarContract,
		//msmContract:         msmContract,
		batchId: batchId.Uint64(),
		l2oo:    l2ooContract,
	}, nil
}

func (m *Manager) Start(ctx context.Context) error {
	waitNodeTicker := time.NewTicker(5 * time.Second)
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

	registry := router.NewRegistry(m, m.db)
	r := gin.Default()
	registry.Register(r)

	var s *http.Server
	s = &http.Server{
		Addr:    m.httpAddr,
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			m.log.Error("api server starts failed", "err", err)
		}
	}()
	m.httpServer = s

	//go func(m *Manager) {
	//	m.log.Info("start rpc server ", "addr", m.grpcAddr)
	//
	//	listener, err := net.Listen("tcp", m.grpcAddr)
	//	if err != nil {
	//		m.log.Error("Could not start tcp listener")
	//	}
	//
	//	opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)
	//
	//	gs := grpc.NewServer(
	//		opt,
	//		grpc.ChainUnaryInterceptor(),
	//	)
	//	reflection.Register(gs)
	//	pb.RegisterCelestiaServiceServer(gs, m)
	//
	//	m.log.Info("grpc info ", "address", listener.Addr().String())
	//	if err := gs.Serve(listener); err != nil {
	//		m.log.Error("Could not start GRPC server")
	//	}
	//	m.gs = gs
	//}(m)

	if m.batchId == 0 {
		m.isFirstBatch = true
	}

	if err := m.getWindowPeriodStartTime(); err != nil {
		m.log.Error("Could not get window period start time", "err", err)
		// for test
		m.outputSubmissionInterval = 360
		m.windowPeriodStartTime = uint64(time.Now().Unix())
		//return err
	}

	go m.babylonSynchronizer.Start()
	go m.ethSynchronizer.Start()
	go m.ethEventProcess.Start()

	m.wg.Add(1)
	go m.work()
	m.log.Info("manager is starting......")
	return nil
}

func (m *Manager) Stop(ctx context.Context) error {
	close(m.done)
	if err := m.httpServer.Shutdown(ctx); err != nil {
		m.log.Error("http server forced to shutdown", "err", err)
		return err
	}
	if err := m.babylonSynchronizer.Close(); err != nil {
		m.log.Error("babylon synchronizer server forced to shutdown", "err", err)
		return err
	}
	if err := m.ethSynchronizer.Close(); err != nil {
		m.log.Error("eth synchronizer server forced to shutdown", "err", err)
		return err
	}
	m.stopped.Store(true)
	m.log.Info("Server exiting")
	return nil
}

func (m *Manager) Stopped() bool {
	return m.stopped.Load()
}

func (m *Manager) getWindowPeriodStartTime() error {
	latestScannedHeight, err := m.db.GetEthScannedHeight()
	if err != nil {
		m.log.Error("failed to get latest scanned height", "err", err)
		return err
	}

	cOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestScannedHeight)),
		From:        m.from,
	}

	index, err := m.l2oo.LatestOutputIndex(cOpts)
	if err != nil {
		m.log.Error("failed to get latest output index", "height", latestScannedHeight, "err", err)
		return err
	}

	output, err := m.l2oo.GetL2Output(cOpts, index)
	if err != nil {
		m.log.Error("failed to get l2 output", "height", latestScannedHeight, "err", err)
		return err
	}

	submissionInterval, err := m.l2oo.SubmissionInterval(cOpts)
	if err != nil {
		m.log.Error("failed to get output submission interval", "err", err)
		return err
	}

	m.outputSubmissionInterval = submissionInterval.Uint64()
	m.windowPeriodStartTime = output.Timestamp.Uint64()

	return nil
}

func (m *Manager) work() {
	fpTicker := time.NewTicker(m.fPTimeout)
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
			fmt.Println(m.windowPeriodStartTime)
			fmt.Println(m.babylonSynchronizer.HeaderTraversal.LatestHeader().Time.Unix())
			fmt.Println(m.babylonSynchronizer.HeaderTraversal.LatestHeader().Height)
			op, err := m.db.GetLatestUnprocessedStateRoot(m.windowPeriodStartTime, m.outputSubmissionInterval)
			if err != nil {
				m.log.Error("failed to get latest unprocessed state root", "err", err)
			}

			if op == nil {
				m.log.Warn("no more state root need to processed", "start", m.windowPeriodStartTime, "end", m.windowPeriodStartTime+m.outputSubmissionInterval+100)
				continue
			}

			if m.isFirstBatch {
				m.windowPeriodStartTime = op.Timestamp.Uint64()
				m.isFirstBatch = false
				continue
			}

			if uint64(m.babylonSynchronizer.HeaderTraversal.LatestHeader().Time.Unix()) < op.Timestamp.Uint64() {
				m.log.Info(fmt.Sprintf("waiting for the babylon block to be synchronized to the timestamp at %v, scanned timestamp: %v", op.Timestamp.Uint64(), m.babylonSynchronizer.HeaderTraversal.LatestHeader().Time.Unix()))
				continue
			}

			m.log.Info("start counting fp signatures", "start", m.windowPeriodStartTime, "end", op.Timestamp.Uint64())

			finalitySignature, fpSignCache, err := m.getMaxSignStateRoot(m.windowPeriodStartTime, op.Timestamp.Uint64())
			if err != nil {
				m.log.Error("failed to get max sign state root", "err", err)
				continue
			}

			if finalitySignature == nil {
				m.log.Warn("no fp signature, skip this state root", "state_root", op.StateRoot)
				m.windowPeriodStartTime = op.Timestamp.Uint64()
				continue
			}

			fmt.Println("==========")
			fmt.Println(finalitySignature)
			fmt.Println(fpSignCache)
			fmt.Println("==========")

			var request types.SignMsgRequest
			var signature *sign.G1Point
			var g2Point *sign.G2Point
			var NonSignerPubkeys []finality.BN254G1Point

			request.BlockNumber = big.NewInt(int64(finalitySignature.BlockNumber))
			request.TxType = common2.MsgSubmitFinalitySignatureType
			request.TxHash = finalitySignature.TransactionHash

			res, err := m.SignMsgBatch(request)
			if errors.Is(err, errNotEnoughSignNode) || errors.Is(err, errNotEnoughVoteNode) {
				m.log.Error("not enough available nodes to sign or not enough available nodes to vote")
				continue
			} else if err != nil {
				m.log.Error("failed to sign msg", "err", err)
				continue
			}
			m.log.Info("success to sign msg", "txHash", hexutil.Encode(request.TxHash), "signature", res.Signature, "block_number", request.BlockNumber.Int64())

			signature = res.Signature
			g2Point = res.G2Point
			for _, v := range res.NonSignerPubkeys {
				NonSignerPubkeys = append(NonSignerPubkeys, finality.BN254G1Point{
					X: v.X.BigInt(new(big.Int)),
					Y: v.Y.BigInt(new(big.Int)),
				})
			}
			if err = m.db.SetSignature(store.Signature{
				BlockNumber:     request.BlockNumber.Int64(),
				TransactionHash: request.TxHash,
				Data:            signature.Serialize(),
				Timestamp:       time.Now().Unix(),
			}); err != nil {
				m.log.Error("failed to store signature", "err", err)
				return
			}

			err = m.db.SetBatchStakeDetails(m.batchId, fpSignCache, *finalitySignature)
			if err != nil {
				m.log.Error("failed to store batch stake details", "err", err)
				continue
			}

			opts, err := client.NewTransactOpts(m.ctx, m.ethChainID, m.privateKey)
			if err != nil {
				m.log.Error("failed to new transact opts", "err", err)
				continue
			}
			fmt.Println(finalitySignature.SubmitFinalitySignature.StateRoot)
			finalityBatch := finality.IFinalityRelayerManagerFinalityBatch{
				StateRoot:     common.HexToHash(finalitySignature.SubmitFinalitySignature.StateRoot),
				L2BlockNumber: big.NewInt(int64(finalitySignature.SubmitFinalitySignature.L2BlockNumber)),
				L1BlockHash:   common.HexToHash(finalitySignature.SubmitFinalitySignature.L1BlockHash),
				L1BlockNumber: big.NewInt(int64(finalitySignature.SubmitFinalitySignature.L1BlockNumber)),
				MsgHash:       crypto.Keccak256Hash(common.Hex2Bytes(finalitySignature.SubmitFinalitySignature.StateRoot)),
			}

			totalBtcStake, err := m.db.GetBatchTotalBabylonStakeAmount(m.batchId)
			if err != nil {
				m.log.Error("failed to get total btc stake", "err", err)
				continue
			}
			//symbioticStake, err := m.db.GetSymbioticFpIds(m.batchId)
			//if err != nil {
			//	m.log.Error("failed to get symbiotic stake", "err", err)
			//	continue
			//}

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
				TotalMantaStake: big.NewInt(1),
			}

			tx, err := m.frmContract.VerifyFinalitySignature(opts, finalityBatch, finalityNonSignerAndSignature, big.NewInt(1))
			if err != nil {
				m.log.Error("failed to craft VerifyFinalitySignature transaction", "err", err)
				continue
			}
			rTx, err := m.rawFrmContract.RawTransact(opts, tx.Data())
			if err != nil {
				m.log.Error("failed to raw VerifyFinalitySignature transaction", "err", err)
				continue
			}
			err = m.ethClient.SendTransaction(m.ctx, tx)
			if err != nil {
				m.log.Error("failed to send VerifyFinalitySignature transaction", "err", err)
				continue
			}

			receipt, err := client.GetTransactionReceipt(m.ctx, m.ethClient, rTx.Hash())
			if err != nil {
				m.log.Error("failed to get verify finality transaction receipt", "err", err)
				continue
			}

			m.log.Info("success to send verify finality signature transaction", "tx_hash", receipt.TxHash.String())

			m.batchId++
			m.windowPeriodStartTime = op.Timestamp.Uint64()
		case <-m.done:
			return
		}
	}
}

func (m *Manager) SignMsgBatch(request types.SignMsgRequest) (*types.SignResult, error) {
	m.log.Info("received sign request", "tx_type", request.TxType, "block_number", request.BlockNumber.Uint64(), "tx_hash", hexutil.Encode(request.TxHash))

	activeMember, err := m.db.GetActiveMember()
	if err != nil {
		m.log.Error("failed to get active member from db", "err", err)
		return nil, err
	}
	availableNodes := m.availableNodes(activeMember.Members)
	if len(availableNodes) == 0 {
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
		var txInfo types4.TransactionInfo
		mCBD.Unmarshal(txMsg.Data)
		if err := m.db.SetCreateBTCDelegationMsg(store.CreateBTCDelegation{
			CBD:    mCBD,
			TxHash: txMsg.TransactionHash,
		}); err != nil {
			return err
		}
		txInfo.Unmarshal(mCBD.StakingTx)
		btcTx, err := types2.NewBtcTransaction(txInfo.Transaction)
		if err != nil {
			m.log.Error("failed to new btc transaction", "err", err)
			return err
		}
		if err = m.db.SetBabylonDelegationKey(txMsg.TransactionHash, []byte(btcTx.Transaction.TxHash().String())); err != nil {
			m.log.Error("failed to store babylon delegation key", "err", err)
			return err
		}
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

//func (m *Manager) StateRootSignIDs(ctx context.Context, request *pb.StateRootSignIDsRequest) (*pb.StateRootSignIDsResponse, error) {
//	if request == nil {
//		log.Warn("invalid request: request body is empty")
//		return nil, errors.New("invalid request: request body is empty")
//	}
//
//	m.log.Info("celestia: blob request", "id", request.Ids)
//
//	idsByte, err := hex.DecodeString(request.Ids)
//	if err != nil {
//		m.log.Error("failed to decode symbiotic fp ids to byte", "err", err)
//		return nil, err
//	}
//	ctx2, cancel := context.WithTimeout(ctx, m.celestiaDa.GetTimeout)
//	blobs, err := m.celestiaDa.Client.Get(ctx2, [][]byte{idsByte}, m.celestiaDa.Namespace)
//	cancel()
//	if err != nil || len(blobs) != 1 {
//		m.log.Error(fmt.Errorf("celestia: failed to resolve frame: %w, len= %v", err, len(blobs)).Error())
//		return nil, fmt.Errorf("celestia: failed to resolve frame: %w, len= %v", err, len(blobs))
//	}
//
//	var signRequest store.SignRequest
//	err = json.Unmarshal(blobs[0], &signRequest)
//	if err != nil {
//		m.log.Error("failed to unmarshal symbiotic fp ids to sign request", "err", err)
//		return nil, err
//	}
//
//	cOpts := &bind.CallOpts{
//		BlockNumber: big.NewInt(int64(request.L1BlockNumber)),
//		From:        m.from,
//	}
//	operator, err := m.msmContract.Operators(cOpts, common.HexToAddress(signRequest.SignAddress))
//	if err != nil {
//		return nil, fmt.Errorf("failed to get operator info at block: %v, err: %v", cOpts.BlockNumber, err)
//	}
//
//	if operator.OperatorName == "" {
//		m.log.Warn(fmt.Sprintf("node %s is not operator", signRequest.SignAddress))
//		return nil, nil
//	} else {
//		if operator.Paused {
//			m.log.Warn("operator is paused", "address", signRequest.SignAddress)
//			return nil, nil
//		}
//	}
//
//	stakeAmount, err := m.getSymbioticOperatorStakeAmount(signRequest.SignAddress)
//	if err != nil {
//		m.log.Error("failed to get operator stake amount", "address", signRequest.SignAddress, "err", err)
//		return nil, nil
//	}
//
//	var sFI = store.SymbioticFpIds{
//		BatchId:          m.batchId,
//		TotalStakeAmount: stakeAmount,
//	}
//
//	sFI.SignRequests = append(sFI.SignRequests, signRequest)
//
//	err = m.db.SetSymbioticFpIds(sFI)
//	if err != nil {
//		m.log.Error("failed to store symbiotic fp ids", "err", err)
//		return nil, err
//	}
//
//	return &pb.StateRootSignIDsResponse{
//		Success: true,
//		Message: "call state root sign ids success",
//	}, nil
//}

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

func (m *Manager) getMaxSignStateRoot(start, end uint64) (*store.WrapperSFs, map[string]string, error) {
	stateRootCountCache := make(map[string]int64)
	stateRootMsgCache := make(map[string]store.WrapperSFs)
	fpSignCache := make(map[string]string)

	finalitySignatures, err := m.db.GetBabylonFinalitySignatureByTimestamp(start, end)
	if err != nil {
		return nil, nil, err
	}

	op, err := m.db.GetOutputProposedByTimestamp(start)
	if err != nil {
		return nil, nil, err
	}

	for _, fs := range finalitySignatures {
		fmt.Println("=========333333=========")
		fmt.Println(fs)
		fmt.Println("=========333333=========")
		if fs.SubmitFinalitySignature.L2BlockNumber == op.L2BlockNumber.Uint64() {
			if stateRootCountCache[fs.SubmitFinalitySignature.StateRoot] == 0 {
				stateRootMsgCache[fs.SubmitFinalitySignature.StateRoot] = fs
			}
			stateRootCountCache[fs.SubmitFinalitySignature.StateRoot]++
			fpSignCache[fs.SubmitFinalitySignature.FpPubkeyHex] = fs.SubmitFinalitySignature.StateRoot
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
	submitFinalitySignature := stateRootMsgCache[maxSignStateRoot]

	return &submitFinalitySignature, fpSignCache, nil
}

func (m *Manager) getSymbioticOperatorStakeAmount(operator string) (*big.Int, error) {
	query := fmt.Sprintf(`{"query":"query {\n  vaultUpdates(first: 1, where: {operator: \"%s\"}, orderBy: timestamp, orderDirection: desc) {\n    vaultTotalActiveStaked\n  }\n}"}`, operator)
	jsonQuery := []byte(query)

	req, err := http.NewRequest("POST", m.sStakeUrl, bytes.NewBuffer(jsonQuery))
	if err != nil {
		m.log.Error("Error creating HTTP request:", "err", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, multipart/mixed")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		m.log.Error("Error sending HTTP request:", "err", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		m.log.Error("Error reading response body:", "err", err)
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		m.log.Error("Error parsing JSON response:", "err", err)
		return nil, err
	}

	var totalStaked = big.NewInt(0)
	if data, exists := result["data"]; exists {
		if vaultUpdates, exists := data.(map[string]interface{})["vaultUpdates"]; exists {
			if len(vaultUpdates.([]interface{})) > 0 {
				vaultTotalActiveStaked := vaultUpdates.([]interface{})[0].(map[string]interface{})["vaultTotalActiveStaked"]
				totalStaked, _ = new(big.Int).SetString(vaultTotalActiveStaked.(string), 10)
				m.log.Info(fmt.Sprintf("operator %s vaultTotalActiveStaked: %s", operator, vaultTotalActiveStaked))
			} else {
				m.log.Warn(fmt.Sprintf("operator %s no vault updates found", operator))
			}
		} else {
			m.log.Warn(fmt.Sprintf("operator %s no vaultUpdates field found in response data", operator))
		}
	} else {
		m.log.Warn(fmt.Sprintf("operator %s no data field found in JSON response", operator))
	}

	return totalStaked, nil
}
