package node

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/common/bigint"

	client "github.com/celestiaorg/celestia-openrpc"
	"github.com/celestiaorg/celestia-openrpc/types/header"
)

type CelestiaHeaderTraversal struct {
	celestiaClient *client.Client

	latestHeader        *header.ExtendedHeader
	lastTraversedHeader *header.ExtendedHeader

	blockConfirmationDepth *big.Int
}

func NewCelestiaHeaderTraversal(celestiaClient *client.Client, fromHeader *header.ExtendedHeader, confDepth *big.Int) *CelestiaHeaderTraversal {
	return &CelestiaHeaderTraversal{
		celestiaClient:         celestiaClient,
		lastTraversedHeader:    fromHeader,
		blockConfirmationDepth: confDepth,
	}
}

func (f *CelestiaHeaderTraversal) LatestHeader() *header.ExtendedHeader {
	return f.latestHeader
}

func (f *CelestiaHeaderTraversal) LastTraversedHeader() *header.ExtendedHeader {
	return f.lastTraversedHeader
}

func (f *CelestiaHeaderTraversal) NextHeaders(maxSize uint64) ([]*header.ExtendedHeader, error) {
	ctx := context.Background()
	latestHeader, err := f.celestiaClient.Header.LocalHead(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to query latest block: %w", err)
	} else if latestHeader == nil {
		return nil, fmt.Errorf("latest header unreported")
	} else {
		f.latestHeader = latestHeader
	}
	log.Info("celestia header traversal db latest header: ", "height", latestHeader.Height())

	endHeight := new(big.Int).Sub(big.NewInt(int64(latestHeader.Height())), f.blockConfirmationDepth)
	if endHeight.Sign() < 0 {
		// No blocks with the provided confirmation depth available
		return nil, nil
	}

	if f.lastTraversedHeader != nil {
		cmp := big.NewInt(int64(f.lastTraversedHeader.Height())).Cmp(endHeight)
		if cmp == 0 {
			return nil, nil
		} else if cmp > 0 {
			return nil, ErrHeaderTraversalAheadOfProvider
		}
	}

	nextHeight := bigint.Zero
	if f.lastTraversedHeader != nil {
		nextHeight = new(big.Int).Add(big.NewInt(int64(f.lastTraversedHeader.Height())), bigint.One)
	}

	endHeight = bigint.Clamp(nextHeight, endHeight, maxSize)
	headers, err := f.BlockHeadersByRange(ctx, nextHeight, endHeight)
	if err != nil {
		return nil, fmt.Errorf("error querying blocks by range: %w", err)
	}
	if len(headers) == 0 {
		return nil, nil
	}
	err = f.checkHeaderListByHash(f.lastTraversedHeader, headers)
	if err != nil {
		log.Error("next headers check blockList by hash", "error", err)
		return nil, err
	}

	numHeaders := len(headers)
	if numHeaders == 0 {
		return nil, nil
	} else if f.lastTraversedHeader != nil && headers[0].LastBlockID.Hash.String() != f.lastTraversedHeader.Hash().String() {
		log.Error("Err header traversal and provider mismatched state", "parentHash = ", headers[0].LastBlockID.Hash.String(), "hash", f.lastTraversedHeader.Hash().String())
		return nil, ErrHeaderTraversalAndProviderMismatchedState
	}
	f.lastTraversedHeader = headers[numHeaders-1]
	return headers, nil
}

func (f *CelestiaHeaderTraversal) checkHeaderListByHash(dbLatestHeader *header.ExtendedHeader, headerList []*header.ExtendedHeader) error {
	if len(headerList) == 0 {
		return nil
	}
	if len(headerList) == 1 {
		return nil
	}
	// check input and db
	// input first ParentHash = dbLatestHeader.Hash
	if dbLatestHeader != nil && headerList[0].LastBlockID.Hash.String() != dbLatestHeader.Hash().String() {
		log.Error("check header list by hash", "parentHash = ", headerList[0].LastBlockID.Hash.String(), "hash", dbLatestHeader.Hash().String())
		return ErrHeaderTraversalCheckHeaderByHashDelDbData
	}

	// check input
	for i := 1; i < len(headerList); i++ {
		if headerList[i].LastBlockID.Hash.String() != headerList[i-1].Hash().String() {
			return fmt.Errorf("check header list by hash: block parent hash not equal parent block hash")
		}
	}
	return nil
}

func (f *CelestiaHeaderTraversal) ChangeLastTraversedHeaderByDelAfter(dbLatestHeader *header.ExtendedHeader) {
	f.lastTraversedHeader = dbLatestHeader
}

func (f *CelestiaHeaderTraversal) BlockHeadersByRange(ctx context.Context, nextHeight *big.Int, endHeight *big.Int) ([]*header.ExtendedHeader, error) {
	var headers []*header.ExtendedHeader
	if nextHeight.Cmp(endHeight) == 0 {
		nextHeader, err := f.celestiaClient.Header.GetByHeight(ctx, endHeight.Uint64())
		if err != nil {
			return nil, fmt.Errorf("failed to get celestia header, height = %v , err = %v", nextHeight, err)
		}
		headers = append(headers, nextHeader)
		return headers, nil
	}

	nextHeader, err := f.celestiaClient.Header.GetByHeight(ctx, nextHeight.Uint64()-1)
	if err != nil {
		return nil, fmt.Errorf("failed to get celestia header, height = %v , err = %v", nextHeight, err)
	}
	headers, err = f.celestiaClient.Header.GetRangeByHeight(ctx, nextHeader, endHeight.Uint64())
	if err != nil {
		return nil, fmt.Errorf("failed to get celestia header by range, next_height = %v , end_height = %v , err = %v", nextHeight, endHeight, err)
	}
	return headers, nil
}
