package manager

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"net"
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
	"github.com/Manta-Network/manta-fp-aggregator/client"
	common2 "github.com/Manta-Network/manta-fp-aggregator/common"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/Manta-Network/manta-fp-aggregator/manager/celestia"
	"github.com/Manta-Network/manta-fp-aggregator/manager/protobuf/pb"
	"github.com/Manta-Network/manta-fp-aggregator/manager/router"
	"github.com/Manta-Network/manta-fp-aggregator/manager/types"
	"github.com/Manta-Network/manta-fp-aggregator/sign"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer/node"
	"github.com/Manta-Network/manta-fp-aggregator/ws/server"

	types4 "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	types2 "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	types3 "github.com/babylonlabs-io/babylon/x/finality/types"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

var (
	errNotEnoughSignNode = errors.New("not enough available nodes to sign")
	errNotEnoughVoteNode = errors.New("not enough available nodes to vote")
)

type Manager struct {
	wg          sync.WaitGroup
	done        chan struct{}
	log         log.Logger
	db          *store.Storage
	wsServer    server.IWebsocketManager
	NodeMembers []string
	httpAddr    string
	httpServer  *http.Server
	grpcAddr    string
	gs          *grpc.Server
	mu          sync.Mutex

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
	batchId         uint64
	l2BlockNumber   uint64

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

	ethNodeClient, err := node.DialEthClient(ctx, cfg.EthRpc)
	if err != nil {
		return nil, err
	}
	contractEventChan := make(chan store.ContractEvent, 100)
	ethSynchronizer, err := synchronizer.NewEthSynchronizer(cfg, db, ethNodeClient, logger, shutdown, contractEventChan)
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
		batchId:             batchId.Uint64(),
		l2BlockNumber:       0,
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

	go func(m *Manager) {
		m.log.Info("start rpc server ", "addr", m.grpcAddr)

		listener, err := net.Listen("tcp", m.grpcAddr)
		if err != nil {
			m.log.Error("Could not start tcp listener")
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(),
		)
		reflection.Register(gs)
		pb.RegisterCelestiaServiceServer(gs, m)

		m.log.Info("grpc info ", "address", listener.Addr().String())
		if err := gs.Serve(listener); err != nil {
			m.log.Error("Could not start GRPC server")
		}
		m.gs = gs
	}(m)

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

func (m *Manager) work() {
	fpTicker := time.NewTicker(m.fPTimeout)
	defer m.wg.Done()

	stateRootCountCache := make(map[string]int64)
	stateRootMsgCache := make(map[string]store.TxMessage)
	fpSignCache := make(map[string]string)

	for {
		select {
		case txMsg := <-m.txMsgChan:
			go func(txMsg store.TxMessage) {
				if txMsg.Type == common2.MsgSubmitFinalitySignatureType {
					var sigParams store.WrapperSFs
					if err := json.Unmarshal(txMsg.Data, &sigParams); err != nil {
						m.log.Error("failed to unmarshal submitFinalitySignature JSON:", "err", err)
						return
					}

					decodedStateRootBytes, err := base64.StdEncoding.DecodeString(sigParams.SubmitFinalitySignature.StateRoot)
					if err != nil {
						m.log.Error("failed to decode stateRoot:", "err", err)
						return
					}
					stateRoot := common.Bytes2Hex(decodedStateRootBytes)

					m.mu.Lock()
					defer m.mu.Unlock()
					if stateRootCountCache[stateRoot] == 0 {
						stateRootMsgCache[stateRoot] = txMsg
					}

					fpSignCache[sigParams.SubmitFinalitySignature.FpPubkeyHex] = stateRoot
					stateRootCountCache[stateRoot]++
				} else {
					if err := m.storeDelegateMsgData(txMsg); err != nil {
						m.log.Error("failed to store delegate msg data", "err", err)
					}
				}
			}(txMsg)
		case <-fpTicker.C:
			if len(stateRootCountCache) == 0 {
				continue
			}

			stateRoot := getMaxSignStateRoot(stateRootCountCache)
			msg := stateRootMsgCache[stateRoot]

			var sigParams store.WrapperSFs
			if err := json.Unmarshal(msg.Data, &sigParams); err != nil {
				m.log.Error("failed to unmarshal submitFinalitySignature JSON:", "err", err)
				return
			}
			if sigParams.SubmitFinalitySignature.L2BlockNumber <= m.l2BlockNumber {
				continue
			}

			var request types.SignMsgRequest
			var signature *sign.G1Point
			var g2Point *sign.G2Point
			var NonSignerPubkeys []finality.BN254G1Point

			request.BlockNumber = big.NewInt(int64(msg.BlockHeight))
			request.TxType = msg.Type
			request.TxHash = msg.TransactionHash

			if sig, err := m.db.GetSignature(request.BlockNumber.Int64()); len(sig.Data) > 0 {
				if err != nil {
					m.log.Error("failed to get signature by tx hash", "tx_hash", hexutil.Encode(request.TxHash), "err", err)
					continue
				}
				signature, err = new(sign.G1Point).Deserialize(sig.Data)
				if err != nil {
					m.log.Error("failed to deserialize signature", "err", err)
					continue
				}
				m.log.Info("get stored signature ", "tx_hash", hexutil.Encode(request.TxHash), "sig", sig)
			} else {
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
			}

			opts, err := client.NewTransactOpts(m.ctx, m.ethChainID, m.privateKey)
			if err != nil {
				m.log.Error("failed to new transact opts", "err", err)
				continue
			}

			finalityBatch := finality.IFinalityRelayerManagerFinalityBatch{
				StateRoot:     common.HexToHash(stateRoot),
				L2BlockNumber: big.NewInt(int64(sigParams.SubmitFinalitySignature.L2BlockNumber)),
				L1BlockHash:   common.HexToHash(sigParams.SubmitFinalitySignature.L1BlockHash),
				//todo for test
				L1BlockNumber: big.NewInt(1),
				//L1BlockNumber: big.NewInt(int64(sigParams.SubmitFinalitySignature.L1BlockNumber)),
				MsgHash: crypto.Keccak256Hash(msg.Data),
			}

			totalBtcStake, err := m.db.GetBTCDelegateAmount()
			if err != nil {
				m.log.Error("failed to get total btc stake", "err", err)
				continue
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

			err = m.db.SetBatchStakeDetails(m.batchId-1, fpSignCache, stateRoot, msg.BlockHeight)
			if err != nil {
				m.log.Error("failed to store batch stake details", "err", err)
				continue
			}

			m.l2BlockNumber = sigParams.SubmitFinalitySignature.L2BlockNumber
			// init cache
			stateRootCountCache = make(map[string]int64)
			stateRootMsgCache = make(map[string]store.TxMessage)
			fpSignCache = make(map[string]string)
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
	//todo 2/3 signer to vote
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
	case common2.MsgCreateFinalityProvider:
		var mCFP types2.MsgCreateFinalityProvider
		mCFP.Unmarshal(txMsg.Data)
		return nil
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
	case common2.MsgCommitPubRandList:
		var mCPR types3.MsgCommitPubRandList
		mCPR.Unmarshal(txMsg.Data)
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
	case common2.MsgSelectiveSlashingEvidence:
		var mSSE types2.MsgSelectiveSlashingEvidence
		mSSE.Unmarshal(txMsg.Data)
		return nil
	default:
		return errors.New("unknown babylon tx msg type")
	}
}

func (m *Manager) StateRootSignIDs(ctx context.Context, request *pb.StateRootSignIDsRequest) (*pb.StateRootSignIDsResponse, error) {
	if request == nil {
		log.Warn("invalid request: request body is empty")
		return nil, errors.New("invalid request: request body is empty")
	}

	m.log.Info("celestia: blob request", "id", request.Ids)

	idsByte, err := hex.DecodeString(request.Ids)
	if err != nil {
		m.log.Error("failed to decode symbiotic fp ids to byte", "err", err)
		return nil, err
	}
	ctx2, cancel := context.WithTimeout(ctx, m.celestiaDa.GetTimeout)
	blobs, err := m.celestiaDa.Client.Get(ctx2, [][]byte{idsByte}, m.celestiaDa.Namespace)
	cancel()
	if err != nil || len(blobs) != 1 {
		m.log.Error(fmt.Errorf("celestia: failed to resolve frame: %w, len=", err, len(blobs)).Error())
		return nil, fmt.Errorf("celestia: failed to resolve frame: %w, len=", err, len(blobs))
	}

	var signRequest store.SignRequest
	err = json.Unmarshal(blobs[0], &signRequest)
	if err != nil {
		m.log.Error("failed to unmarshal symbiotic fp ids to sign request", "err", err)
		return nil, err
	}

	var sFI = store.SymbioticFpIds{
		BatchId: m.batchId,
	}

	sFI.SignRequests = append(sFI.SignRequests, signRequest)

	err = m.db.SetSymbioticFpIds(sFI)
	if err != nil {
		m.log.Error("failed to store symbiotic fp ids", "err", err)
		return nil, err
	}

	return &pb.StateRootSignIDsResponse{
		Success: true,
		Message: "call state root sign ids success",
	}, nil
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

func getMaxSignStateRoot(m map[string]int64) string {
	var maxSignStateRoot string
	var maxStateRootCount int64

	for k, v := range m {
		if v > maxStateRootCount || maxSignStateRoot == "" {
			maxSignStateRoot = k
			maxStateRootCount = v
		}
	}

	return maxSignStateRoot
}
