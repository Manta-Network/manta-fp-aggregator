package manager

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	types4 "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
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
	"github.com/Manta-Network/manta-fp-aggregator/client"
	common2 "github.com/Manta-Network/manta-fp-aggregator/common"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/Manta-Network/manta-fp-aggregator/manager/router"
	"github.com/Manta-Network/manta-fp-aggregator/manager/rpc"
	"github.com/Manta-Network/manta-fp-aggregator/manager/types"
	"github.com/Manta-Network/manta-fp-aggregator/sign"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer/node"
	"github.com/Manta-Network/manta-fp-aggregator/ws/server"

	types2 "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	types3 "github.com/babylonlabs-io/babylon/x/finality/types"
	"github.com/gin-gonic/gin"
)

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

	ctx     context.Context
	stopped atomic.Bool

	ethChainID      uint64
	privateKey      *ecdsa.PrivateKey
	from            common.Address
	ethClient       *ethclient.Client
	frmContract     *finality.FinalityRelayerManager
	frmContractAddr common.Address
	rawFrmContract  *bind.BoundContract
	barContract     *bls.BLSApkRegistry
	barContractAddr common.Address
	rawBarContract  *bind.BoundContract
	batchId         uint64

	signTimeout time.Duration

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

	service := NewFinalityService(db, logger)
	if cfg.Manager.SdkRpc != "" {
		go rpc.NewAndStartFinalityRpcServer(ctx, cfg.Manager.SdkRpc, service)
	}
	txMsgChan := make(chan store.TxMessage, 100)
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

	return &Manager{
		done:                make(chan struct{}),
		log:                 logger,
		db:                  db,
		wsServer:            wsServer,
		NodeMembers:         nodeMemberS,
		ctx:                 ctx,
		privateKey:          priv,
		from:                crypto.PubkeyToAddress(priv.PublicKey),
		signTimeout:         cfg.Manager.SignTimeout,
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
			log.Error("api server starts failed", err)
		}
	}()
	m.httpServer = s
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
	defer m.wg.Done()

	for {
		select {
		case txMsg := <-m.txMsgChan:
			go func(txMsg store.TxMessage) {
				if txMsg.Type == common2.MsgSubmitFinalitySignatureType {
					var request types.SignMsgRequest
					var signature *sign.G1Point
					var g2Point *sign.G2Point
					var NonSignerPubkeys []finality.BN254G1Point

					request.BlockNumber = big.NewInt(int64(txMsg.BlockHeight))
					request.TxType = txMsg.Type
					request.TxHash = txMsg.TransactionHash

					data, err := m.processTxMsgData(txMsg)
					if err != nil {
						m.log.Error("failed to process tx msg data", "err", err)
						return
					}

					if sig, err := m.db.GetSignature(request.BlockNumber.Int64()); len(sig.Data) > 0 {
						if err != nil {
							m.log.Error("failed to get signature by tx hash", "tx_hash", hexutil.Encode(request.TxHash), "err", err)
							return
						}
						signature, err = new(sign.G1Point).Deserialize(sig.Data)
						if err != nil {
							m.log.Error("failed to deserialize signature", "err", err)
							return
						}
						m.log.Info("get stored signature ", "tx_hash", hexutil.Encode(request.TxHash), "sig", sig)
					} else {
						res, err := m.SignMsgBatch(request)
						if errors.Is(err, errNotEnoughSignNode) || errors.Is(err, errNotEnoughVoteNode) {
							m.log.Error("not enough available nodes to sign or not enough available nodes to vote")
							return
						} else if err != nil {
							m.log.Error("failed to sign msg", "err", err)
							return
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
						return
					}

					finalityBatch := finality.IFinalityRelayerManagerFinalityBatch{
						StateRoot:       common.HexToHash("1"),
						L2BlockNumber:   big.NewInt(1),
						L1BlockHash:     common.HexToHash("1"),
						L1BlockNumber:   big.NewInt(int64(1)),
						MsgHash:         crypto.Keccak256Hash(data),
						DisputeGameType: 0,
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
						TotalBtcStake:   big.NewInt(1),
						TotalMantaStake: big.NewInt(1),
					}

					tx, err := m.frmContract.VerifyFinalitySignature(opts, finalityBatch, finalityNonSignerAndSignature, big.NewInt(1))
					if err != nil {
						m.log.Error("failed to craft VerifyFinalitySignature transaction", "err", err)
						return
					}
					rTx, err := m.rawFrmContract.RawTransact(opts, tx.Data())
					if err != nil {
						m.log.Error("failed to raw VerifyFinalitySignature transaction", "err", err)
						return
					}
					err = m.ethClient.SendTransaction(m.ctx, tx)
					if err != nil {
						m.log.Error("failed to send VerifyFinalitySignature transaction", "err", err)
						return
					}

					receipt, err := client.GetTransactionReceipt(m.ctx, m.ethClient, rTx.Hash())
					if err != nil {
						m.log.Error("failed to get verify finality transaction receipt", "err", err)
						return
					}
					m.batchId++

					m.log.Info("success to send verify finality signature transaction", "tx_hash", receipt.TxHash.String())
				}
			}(txMsg)
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

func (m *Manager) processTxMsgData(txMsg store.TxMessage) ([]byte, error) {
	switch txMsg.Type {
	case common2.MsgCreateFinalityProvider:
		var mCFP types2.MsgCreateFinalityProvider
		mCFP.Unmarshal(txMsg.Data)
		return mCFP.Marshal()
	case common2.MsgCreateBTCDelegation:
		var mCBD types2.MsgCreateBTCDelegation
		var txInfo types4.TransactionInfo
		mCBD.Unmarshal(txMsg.Data)
		if err := m.db.SetCreateBTCDelegationMsg(m.batchId, txMsg.BlockHeight, store.CreateBTCDelegation{
			CBD:    mCBD,
			TxHash: txMsg.TransactionHash,
		}); err != nil {
			return nil, err
		}
		txInfo.Unmarshal(mCBD.StakingTx)
		btcTx, err := types2.NewBtcTransaction(txInfo.Transaction)
		if err != nil {
			m.log.Error("failed to new btc transaction", "err", err)
			return nil, err
		}
		if err = m.db.SetBabylonDelegationKey(txMsg.TransactionHash, []byte(btcTx.Transaction.TxHash().String())); err != nil {
			m.log.Error("failed to store babylon delegation key", "err", err)
			return nil, err
		}
		return mCBD.Marshal()
	case common2.MsgCommitPubRandList:
		var mCPR types3.MsgCommitPubRandList
		mCPR.Unmarshal(txMsg.Data)
		return mCPR.Marshal()
	case common2.MsgBTCUndelegate:
		var mBU types2.MsgBTCUndelegate
		mBU.Unmarshal(txMsg.Data)
		if err := m.db.SetBtcUndelegateMsg(m.batchId, txMsg.BlockHeight, store.BtcUndelegate{
			BU:     mBU,
			TxHash: txMsg.TransactionHash,
		}); err != nil {
			return nil, err
		}
		return mBU.Marshal()
	case common2.MsgSelectiveSlashingEvidence:
		var mSSE types2.MsgSelectiveSlashingEvidence
		mSSE.Unmarshal(txMsg.Data)
		return mSSE.Marshal()
	default:
		return nil, errors.New("unknown babylon tx msg type")
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
