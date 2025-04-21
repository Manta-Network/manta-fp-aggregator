package node

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/bindings/finality"
	"github.com/Manta-Network/manta-fp-aggregator/client"
	common3 "github.com/Manta-Network/manta-fp-aggregator/common"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/Manta-Network/manta-fp-aggregator/manager/types"
	common2 "github.com/Manta-Network/manta-fp-aggregator/node/common"
	"github.com/Manta-Network/manta-fp-aggregator/sign"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer"
	wsclient "github.com/Manta-Network/manta-fp-aggregator/ws/client"

	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

type Node struct {
	wg         sync.WaitGroup
	done       chan struct{}
	log        log.Logger
	db         *store.Storage
	privateKey *ecdsa.PrivateKey
	from       common.Address

	ctx      context.Context
	cancel   context.CancelFunc
	stopChan chan struct{}
	stopped  atomic.Bool

	wsClient *wsclient.WSClients
	keyPairs *sign.KeyPair

	signTimeout      time.Duration
	waitScanInterval time.Duration
	signRequestChan  chan tdtypes.RPCRequest
	synchronizer     *synchronizer.BabylonSynchronizer
	txMsgChan        chan store.TxMessage
}

func NewFinalityNode(ctx context.Context, db *store.Storage, privKey *ecdsa.PrivateKey, keyPairs *sign.KeyPair, shouldRegister bool, cfg *config.Config, logger log.Logger, shutdown context.CancelCauseFunc) (*Node, error) {
	from := crypto.PubkeyToAddress(privKey.PublicKey)

	pubkey := crypto.CompressPubkey(&privKey.PublicKey)
	pubkeyHex := hex.EncodeToString(pubkey)
	logger.Info(fmt.Sprintf("pub key is (%s) \n", pubkeyHex))
	if shouldRegister || from.String() == "0xCa7c9907EC2e6641Aed7ed9e380Ce61879D47891" {
		logger.Info("register to operator ...")
		tx, err := registerOperator(ctx, cfg, privKey, pubkeyHex, keyPairs)
		if err != nil {
			logger.Error("failed to register operator", "err", err)
			return nil, err
		}
		logger.Info("success to register operator", "tx_hash", tx.Hash())
	}

	wsClient, err := wsclient.NewWSClient(cfg.Node.WsAddr, "/ws", privKey, pubkeyHex)
	if err != nil {
		return nil, err
	}
	txMsgChan := make(chan store.TxMessage, 100)
	synchronizer, err := synchronizer.NewBabylonSynchronizer(ctx, cfg, db, shutdown, logger, txMsgChan)
	if err != nil {
		return nil, err
	}

	return &Node{
		wg:               sync.WaitGroup{},
		done:             make(chan struct{}),
		stopChan:         make(chan struct{}),
		log:              logger,
		db:               db,
		privateKey:       privKey,
		from:             from,
		ctx:              ctx,
		wsClient:         wsClient,
		keyPairs:         keyPairs,
		signRequestChan:  make(chan tdtypes.RPCRequest, 100),
		signTimeout:      cfg.Node.SignTimeout,
		waitScanInterval: cfg.Node.WaitScanInterval,
		synchronizer:     synchronizer,
		txMsgChan:        txMsgChan,
	}, nil
}

func (n *Node) Start(ctx context.Context) error {
	n.wg.Add(3)
	go n.ProcessMessage()
	go n.sign()
	go n.synchronizer.Start()
	go n.work()
	return nil
}

func (n *Node) Stop(ctx context.Context) error {
	n.cancel()
	close(n.done)
	n.wg.Wait()
	n.synchronizer.Close()
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
			if err := n.synchronizer.ProcessNewFinalityProvider(txMsg); err != nil {
				n.log.Error("failed to process NewFinalityProvider msg", "err", err)
				continue
			}
			if err := n.synchronizer.ProcessCreateBTCDelegation(txMsg); err != nil {
				n.log.Error("failed to process CreateBTCDelegation msg", "err", err)
				continue
			}
			if err := n.synchronizer.ProcessCommitPubRandList(txMsg); err != nil {
				n.log.Error("failed to process CommitPubRandList msg", "err", err)
				continue
			}
			if err := n.synchronizer.ProcessBTCUndelegate(txMsg); err != nil {
				n.log.Error("failed to process BTCUndelegate msg", "err", err)
				continue
			}
			if err := n.synchronizer.ProcessSelectiveSlashingEvidence(txMsg); err != nil {
				n.log.Error("failed to process SelectiveSlashingEvidence msg", "err", err)
				continue
			}
			if err := n.synchronizer.ProcessSubmitFinalitySignature(txMsg); err != nil {
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
				if requestBody.SignType == common3.BabylonSignType {
					if len(requestBody.TxHash) == 0 || requestBody.BlockNumber.Uint64() <= 0 {
						n.log.Error("tx hash and babylon block number must not be nil or negative")
						RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", "tx hash and l2 block number must not be nil or negative")
						if err := n.wsClient.SendMsg(RpcResponse); err != nil {
							n.log.Error("failed to send msg to manager", "err", err)
						}
						continue
					}
				} else if requestBody.SignType == common3.SymbioticSignType {
					if requestBody.StateRoot == "" {
						n.log.Error("state root must not be nil or negative")
						RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", "state root must not be nil")
						if err := n.wsClient.SendMsg(RpcResponse); err != nil {
							n.log.Error("failed to send msg to manager", "err", err)
						}
						continue
					}
				}

				nodeSignRequest.RequestBody = requestBody
				if requestBody.SignType == common3.BabylonSignType {
					go n.handleBabylonSign(req.ID.(tdtypes.JSONRPCStringID), nodeSignRequest)
				} else if requestBody.SignType == common3.SymbioticSignType {
					go n.handleSymbioticSign(req.ID.(tdtypes.JSONRPCStringID), nodeSignRequest)
				}
			}
		}
	}()
}

func (n *Node) handleBabylonSign(resId tdtypes.JSONRPCStringID, req types.NodeSignRequest) error {
	var err error
	var bSign *sign.Signature

	requestBody := req.RequestBody.(types.SignMsgRequest)
	height, err := n.db.GetBabylonScannedHeight()
	if err != nil {
		n.log.Error("node failed to get scanned height", "err", err)
		return err
	}
	if requestBody.BlockNumber.Uint64() <= height {
		bSign, err = n.SignMessage(requestBody)
		if err != nil {
			n.log.Error("node failed to sign messages", "err", err)
			return err
		}
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
		} else {
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
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), n.signTimeout)
		ticker := time.NewTicker(n.waitScanInterval)
		defer cancel()
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				height, err := n.db.GetBabylonScannedHeight()
				if err != nil {
					n.log.Error("node failed to get babylon scanned height", "err", err)
					return err
				}
				if requestBody.BlockNumber.Uint64() > height {
					n.log.Warn(fmt.Sprintf("node received the task from the manager, the height is %v, but the synchronized height is %v", requestBody.BlockNumber.Uint64(), height))
					continue
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
							n.log.Info("send sign response to finality manager successfully ")
							return nil
						}
					} else {
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
	return nil
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
	if requestBody.SignType == common3.BabylonSignType {
		if requestBody.TxType == common3.MsgSubmitFinalitySignatureType {
			exist, sFS := n.db.GetSubmitFinalitySignatureMsg(requestBody.TxHash)
			if exist {
				var sigParams store.WrapperSFs
				if err := json.Unmarshal(sFS.SFSByte, &sigParams); err != nil {
					n.log.Error("failed to unmarshal submitFinalitySignature JSON:", "err", err)
					return nil, err
				}
				decodedStateRootBytes, err := base64.StdEncoding.DecodeString(sigParams.SubmitFinalitySignature.StateRoot)
				if err != nil {
					n.log.Error("failed to decode stateRoot:", "err", err)
					return nil, err
				}
				byteData := crypto.Keccak256Hash(decodedStateRootBytes)
				bSign = n.keyPairs.SignMessage(byteData)
				n.log.Info("success to sign SubmitFinalitySignatureMsg", "signature", bSign.String())
			} else {
				return nil, nil
			}
		}
	} else if requestBody.SignType == common3.SymbioticSignType {
		bSign = n.keyPairs.SignMessage(crypto.Keccak256Hash(common.Hex2Bytes(requestBody.StateRoot)))
		n.log.Info("success to sign SubmitFinalitySignatureMsg", "signature", bSign.String())
	}

	return bSign, nil
}

func registerOperator(ctx context.Context, cfg *config.Config, priKey *ecdsa.PrivateKey, node string, keyPairs *sign.KeyPair) (*types2.Transaction, error) {
	ethCli, err := client.DialEthClientWithTimeout(ctx, cfg.EthRpc, false)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client, err: %v", err)
	}
	frmContract, err := finality.NewFinalityRelayerManager(common.HexToAddress(cfg.Contracts.FrmContractAddress), ethCli)
	if err != nil {
		return nil, fmt.Errorf("failed to new FinalityRelayerManager contract, err: %v", err)
	}
	//bar, err := bls.NewBLSApkRegistry(common.HexToAddress(cfg.Contracts.BarContactAddress), ethCli)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to new BLSApkRegistry contract, err: %v", err)
	//}
	fParsed, err := finality.FinalityRelayerManagerMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get FinalityRelayerManager contract abi, err: %v", err)
	}
	rawFrmContract := bind.NewBoundContract(
		common.HexToAddress(cfg.Contracts.FrmContractAddress), *fParsed, ethCli, ethCli,
		ethCli,
	)
	//bParsed, err := bls.BLSApkRegistryMetaData.GetAbi()
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get BLSApkRegistry contract abi, err: %v", err)
	//}
	//rawBarContract := bind.NewBoundContract(
	//	common.HexToAddress(cfg.Contracts.BarContactAddress), *bParsed, ethCli, ethCli,
	//	ethCli,
	//)

	topts, err := client.NewTransactOpts(ctx, cfg.EthChainID, priKey)
	if err != nil {
		return nil, fmt.Errorf("failed to new transaction option, err: %v", err)
	}

	//nodeAddr := crypto.PubkeyToAddress(priKey.PublicKey)
	//latestBlock, err := ethCli.BlockNumber(ctx)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get latest block, err: %v", err)
	//}

	//cOpts := &bind.CallOpts{
	//	BlockNumber: big.NewInt(int64(latestBlock)),
	//	From:        nodeAddr,
	//}

	//msg, err := bar.GetPubkeyRegMessageHash(cOpts, nodeAddr)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get PubkeyRegistrationMessageHash, err: %v", err)
	//}

	//sigMsg := new(bn254.G1Affine).ScalarMultiplication(sign.NewG1Point(msg.X, msg.Y).G1Affine, keyPairs.PrivKey.BigInt(new(big.Int)))

	//params := bls.IBLSApkRegistryPubkeyRegistrationParams{
	//	PubkeyRegistrationSignature: bls.BN254G1Point{
	//		X: sigMsg.X.BigInt(new(big.Int)),
	//		Y: sigMsg.Y.BigInt(new(big.Int)),
	//	},
	//	PubkeyG1: bls.BN254G1Point{
	//		X: keyPairs.GetPubKeyG1().X.BigInt(new(big.Int)),
	//		Y: keyPairs.GetPubKeyG1().Y.BigInt(new(big.Int)),
	//	},
	//	PubkeyG2: bls.BN254G2Point{
	//		X: [2]*big.Int{keyPairs.GetPubKeyG2().X.A1.BigInt(new(big.Int)), keyPairs.GetPubKeyG2().X.A0.BigInt(new(big.Int))},
	//		Y: [2]*big.Int{keyPairs.GetPubKeyG2().Y.A1.BigInt(new(big.Int)), keyPairs.GetPubKeyG2().Y.A0.BigInt(new(big.Int))},
	//	},
	//}

	//regBlsTx, err := bar.RegisterBLSPublicKey(topts, nodeAddr, params, msg)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to craft RegisterBLSPublicKey transaction, err: %v", err)
	//}
	//fRegBlsTx, err := rawBarContract.RawTransact(topts, regBlsTx.Data())
	//if err != nil {
	//	return nil, fmt.Errorf("failed to raw RegisterBLSPublicKey transaction, err: %v", err)
	//}
	//err = ethCli.SendTransaction(ctx, fRegBlsTx)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to send RegisterBLSPublicKey transaction, err: %v", err)
	//}
	//
	//_, err = client.GetTransactionReceipt(ctx, ethCli, fRegBlsTx.Hash())
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get RegisterBLSPublicKey transaction receipt, err: %v, tx_hash: %v", err, fRegBlsTx.Hash().String())
	//}

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
