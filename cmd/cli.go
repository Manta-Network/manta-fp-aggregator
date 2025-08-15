package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/common/cliapp"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/Manta-Network/manta-fp-aggregator/manager"
	"github.com/Manta-Network/manta-fp-aggregator/node"
	"github.com/Manta-Network/manta-fp-aggregator/node/conversion"
	"github.com/Manta-Network/manta-fp-aggregator/sign"
	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/ws/server"

	"github.com/urfave/cli/v2"
)

var (
	PrivateKeyFlagName        = "private-key"
	KeyPairFlagName           = "key-pair"
	DefaultPubKeyFilename     = "fn_bls.pub"
	DefaultPrivKeyFilename    = "fn_bls.piv"
	CelestiaAuthTokenFlagName = "auth-token"
	KmsIdFlagName             = "kms.id"
	KmsRegionFlagName         = "kms.region"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Value:   "./manta-fp-aggregator.yaml",
		Aliases: []string{"c"},
		Usage:   "path to config file",
		EnvVars: []string{"FP_AGGREGATOR_CONFIG"},
	}
	PrivateKeyFlag = &cli.StringFlag{
		Name:    PrivateKeyFlagName,
		Usage:   "Private Key corresponding to l2fp aggregator",
		EnvVars: []string{"FP_AGGREGATOR_PRIVATE_KEY"},
	}
	KeyPairFlag = &cli.StringFlag{
		Name:    KeyPairFlagName,
		Usage:   "key pair corresponding to l2fp aggregator",
		EnvVars: []string{"FP_AGGREGATOR_KEY_PAIR"},
	}
	CelestiaAuthTokenFlag = &cli.StringFlag{
		Name:    CelestiaAuthTokenFlagName,
		Usage:   "the auth token of celestia node",
		EnvVars: []string{"FP_AGGREGATOR_AUTH_TOKEN"},
	}
	KmsIdFlag = &cli.StringFlag{
		Name:    KmsIdFlagName,
		Usage:   "KMS ID the client will reference",
		EnvVars: []string{"FP_AGGREGATOR_KMS_ID"},
	}
	KmsRegion = &cli.StringFlag{
		Name:    KmsRegionFlagName,
		Usage:   "AWS region the client will connect to",
		EnvVars: []string{"FP_AGGREGATOR_KMS_REGION"},
	}
)

func newCli(GitCommit string, GitDate string) *cli.App {
	nodeFlags := []cli.Flag{ConfigFlag, PrivateKeyFlag, KeyPairFlag, CelestiaAuthTokenFlag, KmsIdFlag, KmsRegion}
	managerFlags := []cli.Flag{ConfigFlag, PrivateKeyFlag, CelestiaAuthTokenFlag, KmsIdFlag, KmsRegion}
	peerIDFlags := []cli.Flag{PrivateKeyFlag}
	return &cli.App{
		Version:              VersionWithCommit(GitCommit, GitDate),
		Description:          "A decentralized Relayer that synchronizes contract events from Babylon",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "node",
				Flags:       nodeFlags,
				Description: "Runs the relayer node service",
				Action:      cliapp.LifecycleCmd(runNode),
			},
			{
				Name:        "manager",
				Flags:       managerFlags,
				Description: "Runs the relayer manager service",
				Action:      cliapp.LifecycleCmd(runManager),
			},
			{
				Name:        "parse-peer-id",
				Flags:       peerIDFlags,
				Description: "Parse peer id of the key",
				Action:      runParsePeerID,
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}

func runNode(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelDebug, true))
	log.SetDefault(logger)

	cfg, err := config.NewConfig(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}

	var kmsId string
	var kmsRegion string
	var privKey *ecdsa.PrivateKey
	var shouldRegister bool
	if cfg.EnableKms {
		if ctx.IsSet(KmsIdFlagName) {
			kmsId = ctx.String(KmsIdFlagName)
		} else {
			return nil, errors.New("need to config kms id")
		}
		if ctx.IsSet(KmsRegionFlagName) {
			kmsRegion = ctx.String(KmsRegionFlagName)
		} else {
			return nil, errors.New("need to config kms region")
		}
	} else {
		if ctx.IsSet(PrivateKeyFlagName) {
			privKey, err = crypto.HexToECDSA(ctx.String(PrivateKeyFlagName))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("need to config private key")
		}
	}

	var authToken string
	if ctx.IsSet(CelestiaAuthTokenFlagName) {
		authToken = ctx.String(CelestiaAuthTokenFlagName)
	} else {
		return nil, errors.New("need to config celestia auth token")
	}

	var keyPairs *sign.KeyPair
	if ctx.IsSet(KeyPairFlagName) {
		keyPairs, err = sign.MakeKeyPairFromString(ctx.String(KeyPairFlagName))
		if err != nil {
			return nil, err
		}
	} else {
		privKeyPath := cfg.Node.KeyPath + "/" + DefaultPrivKeyFilename
		pubKeyPath := cfg.Node.KeyPath + "/" + DefaultPubKeyFilename
		if _, err := os.Stat(privKeyPath); err == nil {
			privKeyData, err := os.ReadFile(privKeyPath)
			if err != nil {
				return nil, err
			}
			keyPairs, err = sign.MakeKeyPairFromString(string(privKeyData))
			if err != nil {
				return nil, err
			}
		} else if errors.Is(err, os.ErrNotExist) {
			keyPairs, err = sign.GenRandomBlsKeys()
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(pubKeyPath, []byte(keyPairs.PubKey.String()), 0o600)
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(privKeyPath, []byte(keyPairs.PrivKey.String()), 0o600)
			if err != nil {
				return nil, err
			}
			shouldRegister = true
		} else {
			return nil, err
		}
	}
	logger.Info("create keyPairs", "pubkey:", keyPairs.PubKey.String())
	logger.Info("should register", "is", shouldRegister)

	db, err := store.NewStorage(cfg.Node.LevelDbFolder)
	if err != nil {
		return nil, err
	}

	return node.NewFinalityNode(ctx.Context, db, privKey, keyPairs, shouldRegister, cfg, logger, shutdown, authToken, kmsId, kmsRegion)
}

func runManager(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelDebug, true))
	log.SetDefault(logger)

	cfg, err := config.NewConfig(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}

	var kmsId string
	var kmsRegion string
	var privKey *ecdsa.PrivateKey
	if cfg.EnableKms {
		if ctx.IsSet(KmsIdFlagName) {
			kmsId = ctx.String(KmsIdFlagName)
		} else {
			return nil, errors.New("need to config kms id")
		}
		if ctx.IsSet(KmsRegionFlagName) {
			kmsRegion = ctx.String(KmsRegionFlagName)
		} else {
			return nil, errors.New("need to config kms region")
		}
	} else {
		if ctx.IsSet(PrivateKeyFlagName) {
			privKey, err = crypto.HexToECDSA(ctx.String(PrivateKeyFlagName))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("need to config private key")
		}
	}

	var authToken string
	if ctx.IsSet(CelestiaAuthTokenFlagName) {
		authToken = ctx.String(CelestiaAuthTokenFlagName)
	} else {
		return nil, errors.New("need to config celestia auth token")
	}

	db, err := store.NewStorage(cfg.Manager.LevelDbFolder)
	if err != nil {
		return nil, err
	}

	wsServer, err := server.NewWSServer(cfg.Manager.WsAddr)
	if err != nil {
		return nil, err
	}

	return manager.NewFinalityManager(ctx.Context, db, wsServer, cfg, shutdown, logger, privKey, authToken, kmsId, kmsRegion)
}

func runParsePeerID(ctx *cli.Context) error {
	privateKey := ctx.String(PrivateKeyFlag.Name)

	var publicBz []byte
	if len(privateKey) != 0 {
		privKey, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return err
		}
		pubkeybytes := crypto.CompressPubkey(&privKey.PublicKey)
		publicBz = pubkeybytes
	} else {
		return errors.New("pri-key needs to be specified")
	}
	pubkeyHex := hex.EncodeToString(publicBz)
	fmt.Println(fmt.Sprintf("pub key is (%s)", pubkeyHex))

	peerId, err := conversion.GetPeerIDFromSecp256PubKey(publicBz)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("peerId is (%s)", peerId))
	return nil
}
