package node

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/bindings/bls"
	"github.com/Manta-Network/manta-fp-aggregator/bindings/finality"
	"github.com/Manta-Network/manta-fp-aggregator/bindings/sfp"
	"github.com/Manta-Network/manta-fp-aggregator/client"
	"github.com/Manta-Network/manta-fp-aggregator/common/httputil"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/Manta-Network/manta-fp-aggregator/manager/types"
	"github.com/Manta-Network/manta-fp-aggregator/metrics"
	common2 "github.com/Manta-Network/manta-fp-aggregator/node/common"
	"github.com/Manta-Network/manta-fp-aggregator/sign"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer"
	wsclient "github.com/Manta-Network/manta-fp-aggregator/ws/client"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/prometheus/client_golang/prometheus"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

type Node struct {
	wg         sync.WaitGroup
	done       chan struct{}
	log        log.Logger
	db         *store.Storage
	privateKey *ecdsa.PrivateKey
	from       common.Address

	cfg      *config.Config
	ctx      context.Context
	cancel   context.CancelFunc
	stopChan chan struct{}
	stopped  atomic.Bool

	wsClient *wsclient.WSClients
	keyPairs *sign.KeyPair

	signTimeout          time.Duration
	waitScanInterval     time.Duration
	signRequestChan      chan tdtypes.RPCRequest
	babylonSynchronizer  *synchronizer.BabylonSynchronizer
	celestiaSynchronizer *synchronizer.CelestiaSynchronizer
	txMsgChan            chan store.TxMessage

	msmContract *sfp.MantaStakingMiddleware

	metricsRegistry *prometheus.Registry
	metrics         metrics.Metricer
	metricsServer   *httputil.HTTPServer
}

func NewFinalityNode(ctx context.Context, db *store.Storage, privKey *ecdsa.PrivateKey, keyPairs *sign.KeyPair, shouldRegister bool, cfg *config.Config, logger log.Logger, shutdown context.CancelCauseFunc, authToken string) (*Node, error) {
	from := crypto.PubkeyToAddress(privKey.PublicKey)

	pubkey := crypto.CompressPubkey(&privKey.PublicKey)
	pubkeyHex := hex.EncodeToString(pubkey)
	logger.Info(fmt.Sprintf("pub key is (%s) \n", pubkeyHex))

	ethCli, err := client.DialEthClientWithTimeout(ctx, cfg.EthRpc, false)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client, err: %v", err)
	}
	if shouldRegister {
		logger.Info("register to operator ...")
		tx, err := registerOperator(ctx, ethCli, cfg, privKey, pubkeyHex, keyPairs)
		if err != nil {
			logger.Error("failed to register operator", "err", err)
			return nil, err
		}
		logger.Info("success to register operator", "tx_hash", tx.Hash())
	}

	msmContract, err := sfp.NewMantaStakingMiddleware(common.HexToAddress(cfg.Contracts.MsmContractAddress), ethCli)
	if err != nil {
		return nil, err
	}

	wsClient, err := wsclient.NewWSClient(cfg.Node.WsAddr, "/ws", privKey, pubkeyHex)
	if err != nil {
		return nil, err
	}

	registry := metrics.NewRegistry()
	metricer := metrics.NewMetrics(registry)

	txMsgChan := make(chan store.TxMessage, 100)
	babylonSynchronizer, err := synchronizer.NewBabylonSynchronizer(ctx, cfg, db, shutdown, logger, txMsgChan, metricer)
	if err != nil {
		return nil, err
	}
	celestiaSynchronizer, err := synchronizer.NewCelestiaSynchronizer(ctx, cfg, db, shutdown, logger, authToken, metricer)
	if err != nil {
		return nil, err
	}

	return &Node{
		wg:                   sync.WaitGroup{},
		done:                 make(chan struct{}),
		stopChan:             make(chan struct{}),
		log:                  logger,
		db:                   db,
		privateKey:           privKey,
		from:                 from,
		cfg:                  cfg,
		ctx:                  ctx,
		msmContract:          msmContract,
		wsClient:             wsClient,
		keyPairs:             keyPairs,
		signRequestChan:      make(chan tdtypes.RPCRequest, 100),
		babylonSynchronizer:  babylonSynchronizer,
		celestiaSynchronizer: celestiaSynchronizer,
		txMsgChan:            txMsgChan,
		metricsRegistry:      registry,
	}, nil
}

func (n *Node) Start(ctx context.Context) error {
	n.wg.Add(3)
	go n.ProcessMessage()
	go n.sign()
	go n.work()

	go n.babylonSynchronizer.Start()
	go n.celestiaSynchronizer.Start()

	if err := n.startMetricsServer(); err != nil {
		n.log.Error("failed to start metrics Server", "err", err)
		return err
	}
	return nil
}

func (n *Node) Stop(ctx context.Context) error {
	n.cancel()
	close(n.done)
	n.wg.Wait()
	n.babylonSynchronizer.Close()
	n.celestiaSynchronizer.Close()
	if n.metricsServer != nil {
		if err := n.metricsServer.Close(); err != nil {
			n.log.Error("failed to close metrics server", "err", err)
		}
	}
	n.stopped.Store(true)
	return nil
}

func (n *Node) Stopped() bool {
	return n.stopped.Load()
}

func (n *Node) work() {
	defer n.wg.Done()
	for {
		select {
		case txMsg := <-n.txMsgChan:
			if err := n.babylonSynchronizer.ProcessNewFinalityProvider(txMsg); err != nil {
				n.log.Error("failed to process NewFinalityProvider msg", "err", err)
				continue
			}
			if err := n.babylonSynchronizer.ProcessCreateBTCDelegation(txMsg); err != nil {
				n.log.Error("failed to process CreateBTCDelegation msg", "err", err)
				continue
			}
			if err := n.babylonSynchronizer.ProcessCommitPubRandList(txMsg); err != nil {
				n.log.Error("failed to process CommitPubRandList msg", "err", err)
				continue
			}
			if err := n.babylonSynchronizer.ProcessBTCUndelegate(txMsg); err != nil {
				n.log.Error("failed to process BTCUndelegate msg", "err", err)
				continue
			}
			if err := n.babylonSynchronizer.ProcessSelectiveSlashingEvidence(txMsg); err != nil {
				n.log.Error("failed to process SelectiveSlashingEvidence msg", "err", err)
				continue
			}
			if err := n.babylonSynchronizer.ProcessSubmitFinalitySignature(txMsg); err != nil {
				n.log.Error("failed to process SubmitFinalitySignature msg", "err", err)
				continue
			}
		}
	}
}

func (n *Node) sign() {
	defer n.wg.Done()

	n.log.Info("start to sign message")

	go func() {
		defer func() {
			n.log.Info("exit sign process")
		}()
		for {
			select {
			case <-n.stopChan:
				return
			case req := <-n.signRequestChan:
				var resId = req.ID.(tdtypes.JSONRPCStringID).String()
				n.log.Info(fmt.Sprintf("dealing resId (%s) ", resId))

				var nodeSignRequest types.NodeSignRequest
				rawMsg := json.RawMessage{}
				nodeSignRequest.RequestBody = &rawMsg

				if err := json.Unmarshal(req.Params, &nodeSignRequest); err != nil {
					n.log.Error("failed to unmarshal ask request")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					if err := n.wsClient.SendMsg(RpcResponse); err != nil {
						n.log.Error("failed to send msg to manager", "err", err)
					}
					continue
				}
				var requestBody types.SignMsgRequest
				if err := json.Unmarshal(rawMsg, &requestBody); err != nil {
					n.log.Error("failed to unmarshal asker's params request body")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					if err := n.wsClient.SendMsg(RpcResponse); err != nil {
						n.log.Error("failed to send msg to manager", "err", err)
					}
					continue
				}

				if requestBody.StartTimestamp <= 0 || requestBody.EndTimestamp <= 0 {
					n.log.Error("start timestamp and end timestamp must not be nil or negative")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", "start timestamp and end timestamp must not be nil or negative")
					if err := n.wsClient.SendMsg(RpcResponse); err != nil {
						n.log.Error("failed to send msg to manager", "err", err)
					}
					continue
				}

				nodeSignRequest.RequestBody = requestBody

				go n.handleSign(req.ID.(tdtypes.JSONRPCStringID), nodeSignRequest)
			}
		}
	}()
}

func (n *Node) handleSign(resId tdtypes.JSONRPCStringID, req types.NodeSignRequest) error {
	var err error
	var bSign *sign.Signature
	requestBody := req.RequestBody.(types.SignMsgRequest)
	ctx, cancel := context.WithTimeout(context.Background(), n.cfg.Node.SignTimeout)
	ticker := time.NewTicker(n.cfg.Node.WaitScanInterval)

	defer cancel()
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if !n.checkSyncStatus(requestBody.EndTimestamp) {
				continue
			}

			maxSignStateRoot, err := n.getMaxSignStateRoot(requestBody)
			if maxSignStateRoot != requestBody.StateRoot {
				signResponse := types.SignMsgResponse{
					G2Point:         nil,
					Signature:       nil,
					NonSignerPubkey: n.keyPairs.GetPubKeyG1().Serialize(),
					Vote:            uint8(common2.DisagreeVote),
				}
				RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
				n.log.Info("node disagree the msg, start to send response to finality manager")

				err = n.wsClient.SendMsg(RpcResponse)
				if err != nil {
					n.log.Error("failed to sendMsg to finality manager", "err", err)
					return err
				} else {
					n.log.Info("send sign response to finality manager successfully ")
					return nil
				}
			} else {
				bSign, err = n.SignMessage(requestBody)
				if bSign != nil {
					signResponse := types.SignMsgResponse{
						G2Point:   n.keyPairs.GetPubKeyG2().Serialize(),
						Signature: bSign.Serialize(),
						Vote:      uint8(common2.AgreeVote),
					}
					RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
					n.log.Info("node agree the msg, start to send response to finality manager")

					err = n.wsClient.SendMsg(RpcResponse)
					if err != nil {
						n.log.Error("failed to sendMsg to finality manager", "err", err)
						return err
					} else {
						n.log.Info("send sign response to finality manager successfully")
						return nil
					}
				}
			}
		case <-ctx.Done():
			n.log.Warn("sign messages timeout !")
			signResponse := types.SignMsgResponse{
				Signature:       nil,
				G2Point:         nil,
				NonSignerPubkey: n.keyPairs.GetPubKeyG1().Serialize(),
				Vote:            uint8(common2.DidNotVote),
			}
			RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
			n.log.Info("node did not vote msg, start to send response to finality manager")

			err = n.wsClient.SendMsg(RpcResponse)
			if err != nil {
				n.log.Error("failed to sendMsg to finality manager", "err", err)
				return err
			} else {
				n.log.Info("send sign response to finality manager successfully ")
				return nil
			}
		}
	}
}

func (n *Node) handleSymbioticSign(resId tdtypes.JSONRPCStringID, req types.NodeSignRequest) error {
	var err error
	var bSign *sign.Signature
	requestBody := req.RequestBody.(types.SignMsgRequest)
	bSign, err = n.SignMessage(requestBody)
	if bSign != nil {
		signResponse := types.SignMsgResponse{
			G2Point:   n.keyPairs.GetPubKeyG2().Serialize(),
			Signature: bSign.Serialize(),
			Vote:      uint8(common2.AgreeVote),
		}
		RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
		n.log.Info("node agree the msg, start to send response to finality manager")

		err = n.wsClient.SendMsg(RpcResponse)
		if err != nil {
			n.log.Error("failed to sendMsg to finality manager", "err", err)
			return err
		} else {
			n.log.Info("send sign response to finality manager successfully ")
			return nil
		}
	}
	return nil
}

func (n *Node) SignMessage(requestBody types.SignMsgRequest) (*sign.Signature, error) {
	var bSign *sign.Signature
	bSign = n.keyPairs.SignMessage(crypto.Keccak256Hash(common.Hex2Bytes(requestBody.StateRoot)))
	n.log.Info("success to sign SubmitFinalitySignatureMsg", "signature", bSign.String())

	return bSign, nil
}

func (n *Node) checkSyncStatus(end uint64) bool {
	babylonSynced := uint64(n.babylonSynchronizer.HeaderTraversal.LastTraversedHeader().Time.Unix()) >= end
	celestiaSynced := uint64(n.celestiaSynchronizer.HeaderTraversal.LastTraversedHeader().Time().Unix()) >= end

	if !babylonSynced {
		n.log.Warn("Babylon sync not completed",
			"required", end,
			"current", n.babylonSynchronizer.HeaderTraversal.LastTraversedHeader().Time.Unix())
		return false
	}

	if !celestiaSynced {
		n.log.Warn("Celestia sync not completed",
			"required", end,
			"current", n.celestiaSynchronizer.HeaderTraversal.LastTraversedHeader().Time().Unix())
		return false
	}
	return true
}

func (n *Node) getMaxSignStateRoot(request types.SignMsgRequest) (string, error) {
	stateRootCountCache := make(map[string]int64)
	symbioticFpSignCache := make(map[string]string)

	n.log.Info("start counting fp signatures", "start", request.StartTimestamp, "end", request.EndTimestamp)

	babylonFinalitySignatures, err := n.db.GetBabylonFinalitySignatureByTimestamp(request.StartTimestamp, request.EndTimestamp)
	if err != nil {
		return "", err
	}
	// celestia block may have an earlier timestamp than eth block
	symbioticFinalitySignatures, err := n.db.GetSymbioticFpBlobsByTimestamp(request.StartTimestamp-uint64(time.Minute.Seconds()), request.EndTimestamp)
	if err != nil {
		return "", err
	}

	for _, bfs := range babylonFinalitySignatures {
		if bfs.SubmitFinalitySignature.L2BlockNumber == request.L2BlockNumber {
			stateRootCountCache[bfs.SubmitFinalitySignature.StateRoot]++
		}
	}

	for _, sfs := range symbioticFinalitySignatures {
		if sfs.SignRequests.L1BlockNumber == request.L1BlockNumber && sfs.SignRequests.L2BlockNumber == request.L2BlockNumber {
			cOpts := &bind.CallOpts{
				BlockNumber: big.NewInt(int64(request.L1BlockNumber)),
				From:        n.from,
			}
			operator, err := n.msmContract.Operators(cOpts, common.HexToAddress(sfs.SignRequests.SignAddress))
			if err != nil {
				n.log.Error(fmt.Errorf("failed to get operator info at block: %v, err: %v", cOpts.BlockNumber, err).Error())
				continue
			}

			if operator.OperatorName == "" {
				n.log.Warn(fmt.Sprintf("node %s is not operator", sfs.SignRequests.SignAddress))
				continue
			} else {
				if operator.Paused {
					n.log.Warn("operator is paused", "address", sfs.SignRequests.SignAddress)
					continue
				}
			}

			if symbioticFpSignCache[sfs.SignRequests.SignAddress] == "" {
				amount, err := n.getSymbioticOperatorStakeAmount(strings.ToLower(sfs.SignRequests.SignAddress))
				if err != nil {
					n.log.Error("failed to get operator stake amount", "address", sfs.SignRequests.SignAddress, "err", err)
					continue
				}
				if amount.Cmp(big.NewInt(0)) > 0 {
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

	return maxSignStateRoot, nil
}

func (n *Node) getSymbioticOperatorStakeAmount(operator string) (*big.Int, error) {
	query := fmt.Sprintf(`{"query":"query {\n  vaultUpdates(first: 1, where: {operator: \"%s\"}, orderBy: timestamp, orderDirection: desc) {\n    vaultTotalActiveStaked\n  }\n}"}`, operator)
	jsonQuery := []byte(query)

	req, err := http.NewRequest("POST", n.cfg.SymbioticStakeUrl, bytes.NewBuffer(jsonQuery))
	if err != nil {
		n.log.Error("Error creating HTTP request:", "err", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, multipart/mixed")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		n.log.Error("Error sending HTTP request:", "err", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		n.log.Error("Error reading response body:", "err", err)
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		n.log.Error("Error parsing JSON response:", "err", err)
		return nil, err
	}

	var totalStaked = big.NewInt(0)
	if data, exists := result["data"]; exists {
		if vaultUpdates, exists := data.(map[string]interface{})["vaultUpdates"]; exists {
			if len(vaultUpdates.([]interface{})) > 0 {
				vaultTotalActiveStaked := vaultUpdates.([]interface{})[0].(map[string]interface{})["vaultTotalActiveStaked"]
				totalStaked, _ = new(big.Int).SetString(vaultTotalActiveStaked.(string), 10)
				n.log.Info(fmt.Sprintf("operator %s vaultTotalActiveStaked: %s", operator, vaultTotalActiveStaked))
			} else {
				n.log.Warn(fmt.Sprintf("operator %s no vault updates found", operator))
			}
		} else {
			n.log.Warn(fmt.Sprintf("operator %s no vaultUpdates field found in response data", operator))
		}
	} else {
		n.log.Warn(fmt.Sprintf("operator %s no data field found in JSON response", operator))
	}

	return totalStaked, nil
}

func registerOperator(ctx context.Context, ethCli *ethclient.Client, cfg *config.Config, priKey *ecdsa.PrivateKey, node string, keyPairs *sign.KeyPair) (*types2.Transaction, error) {
	frmContract, err := finality.NewFinalityRelayerManager(common.HexToAddress(cfg.Contracts.FrmContractAddress), ethCli)
	if err != nil {
		return nil, fmt.Errorf("failed to new FinalityRelayerManager contract, err: %v", err)
	}
	bar, err := bls.NewBLSApkRegistry(common.HexToAddress(cfg.Contracts.BarContactAddress), ethCli)
	if err != nil {
		return nil, fmt.Errorf("failed to new BLSApkRegistry contract, err: %v", err)
	}
	fParsed, err := finality.FinalityRelayerManagerMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get FinalityRelayerManager contract abi, err: %v", err)
	}
	rawFrmContract := bind.NewBoundContract(
		common.HexToAddress(cfg.Contracts.FrmContractAddress), *fParsed, ethCli, ethCli,
		ethCli,
	)
	bParsed, err := bls.BLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get BLSApkRegistry contract abi, err: %v", err)
	}
	rawBarContract := bind.NewBoundContract(
		common.HexToAddress(cfg.Contracts.BarContactAddress), *bParsed, ethCli, ethCli,
		ethCli,
	)

	topts, err := client.NewTransactOpts(ctx, cfg.EthChainID, priKey)
	if err != nil {
		return nil, fmt.Errorf("failed to new transaction option, err: %v", err)
	}

	nodeAddr := crypto.PubkeyToAddress(priKey.PublicKey)
	latestBlock, err := ethCli.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block, err: %v", err)
	}

	cOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestBlock)),
		From:        nodeAddr,
	}

	msg, err := bar.GetPubkeyRegMessageHash(cOpts, nodeAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get PubkeyRegistrationMessageHash, err: %v", err)
	}

	sigMsg := new(bn254.G1Affine).ScalarMultiplication(sign.NewG1Point(msg.X, msg.Y).G1Affine, keyPairs.PrivKey.BigInt(new(big.Int)))

	params := bls.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: bls.BN254G1Point{
			X: sigMsg.X.BigInt(new(big.Int)),
			Y: sigMsg.Y.BigInt(new(big.Int)),
		},
		PubkeyG1: bls.BN254G1Point{
			X: keyPairs.GetPubKeyG1().X.BigInt(new(big.Int)),
			Y: keyPairs.GetPubKeyG1().Y.BigInt(new(big.Int)),
		},
		PubkeyG2: bls.BN254G2Point{
			X: [2]*big.Int{keyPairs.GetPubKeyG2().X.A1.BigInt(new(big.Int)), keyPairs.GetPubKeyG2().X.A0.BigInt(new(big.Int))},
			Y: [2]*big.Int{keyPairs.GetPubKeyG2().Y.A1.BigInt(new(big.Int)), keyPairs.GetPubKeyG2().Y.A0.BigInt(new(big.Int))},
		},
	}

	regBlsTx, err := bar.RegisterBLSPublicKey(topts, nodeAddr, params, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to craft RegisterBLSPublicKey transaction, err: %v", err)
	}
	fRegBlsTx, err := rawBarContract.RawTransact(topts, regBlsTx.Data())
	if err != nil {
		return nil, fmt.Errorf("failed to raw RegisterBLSPublicKey transaction, err: %v", err)
	}
	err = ethCli.SendTransaction(ctx, fRegBlsTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send RegisterBLSPublicKey transaction, err: %v", err)
	}

	_, err = client.GetTransactionReceipt(ctx, ethCli, fRegBlsTx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get RegisterBLSPublicKey transaction receipt, err: %v, tx_hash: %v", err, fRegBlsTx.Hash().String())
	}

	regOTx, err := frmContract.RegisterOperator(topts, node)
	if err != nil {
		return nil, fmt.Errorf("failed to craft RegisterOperator transaction, err: %v", err)
	}
	fRegOTx, err := rawFrmContract.RawTransact(topts, regOTx.Data())
	if err != nil {
		return nil, fmt.Errorf("failed to raw RegisterOperator transaction, err: %v", err)
	}
	err = ethCli.SendTransaction(ctx, fRegOTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send RegisterOperator transaction, err: %v", err)
	}
	_, err = client.GetTransactionReceipt(ctx, ethCli, fRegOTx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get RegisterOperator transaction receipt, err: %v, tx_hash: %v", err, fRegOTx.Hash().String())
	}

	return fRegOTx, nil
}

func (n *Node) startMetricsServer() error {
	n.log.Info("starting metrics server", "addr", n.cfg.Metrics.ListenAddr, "port", n.cfg.Metrics.ListenPort)
	metricsSrv, err := metrics.StartServer(n.metricsRegistry, n.cfg.Metrics.ListenAddr, n.cfg.Metrics.ListenPort)
	if err != nil {
		return fmt.Errorf("failed to start metrics server: %w", err)
	}
	n.log.Info("started metrics server", "addr", metricsSrv.Addr())
	n.metricsServer = metricsSrv
	return nil
}
