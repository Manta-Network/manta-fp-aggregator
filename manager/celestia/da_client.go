package celestia

import (
	"encoding/hex"
	"errors"
	"github.com/Manta-Network/manta-fp-aggregator/config"
	"github.com/rollkit/go-da"
	"time"

	"github.com/celestiaorg/go-square/blob"
	"github.com/celestiaorg/go-square/inclusion"
	"github.com/celestiaorg/go-square/namespace"
	"github.com/rollkit/go-da/proxy"
	"github.com/tendermint/tendermint/crypto/merkle"
)

type DAClient struct {
	Client     da.DA
	Namespace  da.Namespace
	GetTimeout time.Duration
}

func NewDAClient(cfg *config.Config) (*DAClient, error) {
	nsBytes, err := hex.DecodeString(cfg.Manager.CelestiaConfig.Namespace)
	if err != nil {
		return nil, err
	}
	if len(nsBytes) != 10 {
		return nil, errors.New("wrong namespace length")
	}
	var client da.DA
	if cfg.Manager.CelestiaConfig.DaRpc != "" {
		client, err = proxy.NewClient(cfg.Manager.CelestiaConfig.DaRpc, cfg.Manager.CelestiaConfig.AuthToken)
		if err != nil {
			return nil, err
		}
	}

	return &DAClient{
		Client:     client,
		Namespace:  append(make([]byte, 19), nsBytes...),
		GetTimeout: cfg.Manager.CelestiaConfig.Timeout,
	}, nil
}

func CreateCommitment(data da.Blob, ns da.Namespace) ([]byte, error) {
	ins, err := namespace.From(ns)
	if err != nil {
		return nil, err
	}
	return inclusion.CreateCommitment(blob.New(ins, data, 0), merkle.HashFromByteSlices, 64)
}
