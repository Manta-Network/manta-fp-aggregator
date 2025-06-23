package client

import (
	"context"
	"crypto/ecdsa"
	"crypto/tls"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	defaultDialTimeout = 5 * time.Second
)

func DialEthClientWithTimeout(ctx context.Context, url string, disableHTTP2 bool) (
	*ethclient.Client, error) {
	ctxt, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()
	if strings.HasPrefix(url, "http") {
		httpClient := new(http.Client)
		if disableHTTP2 {
			log.Debug("Disabled HTTP/2 support in  eth client")
			httpClient.Transport = &http.Transport{
				TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
			}
		}
		rpcClient, err := rpc.DialHTTPWithClient(url, httpClient)
		if err != nil {
			return nil, err
		}
		return ethclient.NewClient(rpcClient), nil
	}
	return ethclient.DialContext(ctxt, url)
}

func NewTransactOpts(chainId uint64, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	var opts *bind.TransactOpts
	var err error

	if privateKey == nil {
		return nil, errors.New("no private key provided")
	}

	opts, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetUint64(chainId))
	if err != nil {
		return nil, fmt.Errorf("new keyed transactor fail, err: %v", err)
	}

	return opts, err
}

func GetTransactionReceipt(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error

	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		receipt, err = client.TransactionReceipt(ctx, txHash)
		if err != nil && !errors.Is(err, ethereum.NotFound) {
			return nil, err
		}

		if errors.Is(err, ethereum.NotFound) {
			continue
		}
		return receipt, nil
	}
}
