package node

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
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
	common3 "github.com/Manta-Network/manta-fp-aggregator/common"
	"github.com/Manta-Network/manta-fp-aggregator/common/httputil"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	kmssigner "github.com/Manta-Network/manta-fp-aggregator/kms"
	"github.com/Manta-Network/manta-fp-aggregator/manager/types"
	"github.com/Manta-Network/manta-fp-aggregator/metrics"
	common2 "github.com/Manta-Network/manta-fp-aggregator/node/common"
	"github.com/Manta-Network/manta-fp-aggregator/node/router"
	"github.com/Manta-Network/manta-fp-aggregator/sign"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer"
	wsclient "github.com/Manta-Network/manta-fp-aggregator/ws/client"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/gin-gonic/gin"
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

	cfg     *config.Config
	ctx     context.Context
	stopped atomic.Bool

	wsClient   *wsclient.WSClients
	httpServer *http.Server
	keyPairs   *sign.KeyPair

	signRequestChan      chan tdtypes.RPCRequest
	babylonSynchronizer  *synchronizer.BabylonSynchronizer
	celestiaSynchronizer *synchronizer.CelestiaSynchronizer
	txMsgChan            chan store.TxMessage

	msmContract *sfp.MantaStakingMiddleware

	metricsRegistry *prometheus.Registry
	metricsServer   *httputil.HTTPServer
}

func NewFinalityNode(ctx context.Context, db *store.Storage, privKey *ecdsa.PrivateKey, keyPairs *sign.KeyPair, shouldRegister bool, cfg *config.Config, logger log.Logger, shutdown context.CancelCauseFunc, authToken string, kmsId string, kmsRegion string) (*Node, error) {
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
	} else {
		pubkey = &privKey.PublicKey
	}

	from = crypto.PubkeyToAddress(*pubkey)
	pubkeyHex := hex.EncodeToString(crypto.CompressPubkey(pubkey))
	logger.Info(fmt.Sprintf("pub key is (%s) \n", pubkeyHex))

	ethCli, err := client.DialEthClientWithTimeout(ctx, cfg.EthRpc, false)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client, err: %v", err)
	}
	if shouldRegister {
		logger.Info("register to operator ...")
		tx, err := registerOperator(ctx, ethCli, cfg, privKey, pubkeyHex, from, keyPairs, kmsId, kmsClient, logger)
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

	wsClient, err := wsclient.NewWSClient(cfg.Node.WsAddr, "/ws", pubkeyHex)
	if err != nil {
		return nil, err
	}

	registry := metrics.NewRegistry()
	metricer := metrics.NewMetrics(registry)

	txMsgChan := make(chan store.TxMessage, 100)
	//babylonSynchronizer, err := synchronizer.NewBabylonSynchronizer(ctx, cfg, db, shutdown, logger, txMsgChan, metricer)
	//if err != nil {
	//	return nil, err
	//}
	celestiaSynchronizer, err := synchronizer.NewCelestiaSynchronizer(ctx, cfg, db, shutdown, logger, authToken, metricer)
	if err != nil {
		return nil, err
	}

	return &Node{
		wg:              sync.WaitGroup{},
		done:            make(chan struct{}),
		log:             logger,
		db:              db,
		privateKey:      privKey,
		from:            from,
		cfg:             cfg,
		ctx:             ctx,
		msmContract:     msmContract,
		wsClient:        wsClient,
		keyPairs:        keyPairs,
		signRequestChan: make(chan tdtypes.RPCRequest, 100),
		//babylonSynchronizer:  babylonSynchronizer,
		celestiaSynchronizer: celestiaSynchronizer,
		txMsgChan:            txMsgChan,
		metricsRegistry:      registry,
	}, nil
}

func (n *Node) Start(ctx context.Context) error {
	registry := router.NewRegistry()
	r := gin.Default()
	registry.Register(r)

	var s *http.Server
	s = &http.Server{
		Addr:    n.cfg.Node.HttpAddr,
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			n.log.Error("api server starts failed", "err", err)
		}
	}()
	n.httpServer = s

	n.wg.Add(3)
	go n.ProcessMessage()
	go n.sign()
	//go n.work()

	//go n.babylonSynchronizer.Start()
	go n.celestiaSynchronizer.Start()

	if err := n.startMetricsServer(); err != nil {
		n.log.Error("failed to start metrics Server", "err", err)
		return err
	}
	return nil
}

func (n *Node) Stop(ctx context.Context) error {
	close(n.done)
	n.wg.Wait()
	//n.babylonSynchronizer.Close()
	n.celestiaSynchronizer.Close()
	if err := n.db.Close(); err != nil {
		n.log.Error("failed to close db server", "err", err)
		return err
	}
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
		case <-n.done:
			return
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
			case <-n.done:
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
			if err != nil {
				n.log.Error("failed to get max sign state root", "err", err)
			}
			if maxSignStateRoot != requestBody.StateRoot || err != nil {
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
				bSign = n.SignMessage(requestBody)
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

func (n *Node) SignMessage(requestBody types.SignMsgRequest) *sign.Signature {
	var bSign *sign.Signature
	bSign = n.keyPairs.SignMessage(crypto.Keccak256Hash(common.Hex2Bytes(requestBody.StateRoot)))
	n.log.Info("success to sign SubmitFinalitySignatureMsg", "signature", bSign.String())

	return bSign
}

func (n *Node) checkSyncStatus(end uint64) bool {
	//babylonSynced := uint64(n.babylonSynchronizer.HeaderTraversal.LastTraversedHeader().Time.Unix()) >= end
	celestiaSynced := uint64(n.celestiaSynchronizer.HeaderTraversal.LastTraversedHeader().Time().Unix()) >= end

	//if !babylonSynced {
	//	n.log.Warn("Babylon sync not completed",
	//		"required", end,
	//		"current", n.babylonSynchronizer.HeaderTraversal.LastTraversedHeader().Time.Unix())
	//	return false
	//}

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
				minMantaStakeAmount, _ := new(big.Int).SetString(n.cfg.MinMantaStakeAmount, 10)
				if amount.Cmp(minMantaStakeAmount) >= 0 {
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
	query := fmt.Sprintf(`
		query {
			vaultUpdates(
				first: 1, 
				where: {operator: "%s"}, 
				orderBy: timestamp, 
				orderDirection: desc
			) {
				vaultTotalActiveStaked
			}
		}`, operator)

	requestBody := common3.SymbioticStakeRequest{Query: query}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		n.log.Error("Error marshaling JSON request", "err", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", n.cfg.SymbioticStakeUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		n.log.Error("Error creating HTTP request", "err", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	hClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := hClient.Do(req)
	if err != nil {
		n.log.Error("Error sending HTTP request", "err", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		n.log.Error("GraphQL server returned error status",
			"status", resp.Status,
			"response", string(body))
		return nil, fmt.Errorf("graphql server error: %s", resp.Status)
	}

	var response common3.SymbioticStakeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		n.log.Error("Error decoding JSON response", "err", err)
		return nil, err
	}

	if len(response.Errors) > 0 {
		errMsg := response.Errors[0].Message
		n.log.Error("GraphQL query returned errors", "errors", response.Errors)
		return nil, fmt.Errorf("graphql error: %s", errMsg)
	}

	if len(response.Data.VaultUpdates) == 0 {
		n.log.Warn("No vault updates found for operator", "operator", operator)
		return big.NewInt(0), nil
	}

	stakeStr := response.Data.VaultUpdates[0].VaultTotalActiveStaked
	totalStaked, ok := new(big.Int).SetString(stakeStr, 10)
	if !ok {
		n.log.Error("Invalid stake amount format",
			"value", stakeStr,
			"operator", operator)
		return nil, fmt.Errorf("invalid stake amount: %s", stakeStr)
	}

	n.log.Info("Retrieved operator stake amount",
		"operator", operator,
		"amount", totalStaked.String())

	return totalStaked, nil
}

func registerOperator(ctx context.Context, ethCli *ethclient.Client, cfg *config.Config, priKey *ecdsa.PrivateKey, node string, from common.Address, keyPairs *sign.KeyPair, kmsId string, kmsClient *kms.Client, logger log.Logger) (*types2.Transaction, error) {
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
	var topts *bind.TransactOpts
	if cfg.EnableKms {
		topts, err = kmssigner.NewAwsKmsTransactorWithChainIDCtx(ctx, kmsClient,
			kmsId, big.NewInt(int64(cfg.EthChainID)))
	} else {
		topts, err = client.NewTransactOpts(cfg.EthChainID, priKey)
		if err != nil {
			return nil, fmt.Errorf("failed to new transaction option, err: %v", err)
		}
	}
	topts.Context = ctx
	topts.NoSend = true

	latestBlock, err := ethCli.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block, err: %v", err)
	}

	cOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestBlock)),
		From:        from,
	}

	msg, err := bar.GetPubkeyRegMessageHash(cOpts, from)
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

	regBlsTx, err := bar.RegisterBLSPublicKey(topts, from, params, msg)
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

	_, err = client.GetTransactionReceipt(ctx, ethCli, fRegBlsTx, time.Second*10, logger)
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
	_, err = client.GetTransactionReceipt(ctx, ethCli, fRegOTx, time.Second*10, logger)
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
