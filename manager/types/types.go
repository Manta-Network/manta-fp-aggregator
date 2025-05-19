package types

import (
	"context"
	"math/big"

	"github.com/Manta-Network/manta-fp-aggregator/sign"
)

type SignMsgRequest struct {
	StartTimestamp uint64 `json:"start_timestamp"`
	EndTimestamp   uint64 `json:"end_timestamp"`
	L1BlockNumber  uint64 `json:"l1_block_number"`
	L2BlockNumber  uint64 `json:"l2_block_number"`
	StateRoot      string `json:"state_root"`
}

type StakerDelegationRequest struct {
	Address string `json:"address"`
}

type StakerDetailsRequest struct {
	BatchId uint64 `json:"batch_id"`
}

type NodeSignRequest struct {
	Timestamp   int64       `json:"timestamp"`
	Nodes       []string    `json:"nodes"`
	RequestBody interface{} `json:"request_body"`
}

type SignMsgResponse struct {
	Signature       []byte `json:"signature"`
	G2Point         []byte `json:"g2_point"`
	NonSignerPubkey []byte `json:"non_signer_pubkey"`
	Vote            uint8  `json:"vote"`
}

type SignResult struct {
	Signature        *sign.G1Point   `json:"signature"`
	G2Point          *sign.G2Point   `json:"g2_point"`
	NonSignerPubkeys []*sign.G1Point `json:"non_signer_pubkeys"`
}

type VoteStateRoot struct {
	StartTimestamp         uint64 `json:"start_timestamp"`
	EndTimestamp           uint64 `json:"end_timestamp"`
	L1BlockNumber          uint64 `json:"l1_block_number"`
	L1BlockHash            string `json:"l1_block_hash"`
	L2BlockNumber          uint64 `json:"l2_block_number"`
	StateRoot              string `json:"state_root"`
	BabylonHeight          uint64 `json:"babylon_height"`
	BabylonFpSignCache     map[string]string
	SymbioticFpSignList    []string
	TotalSymbioticFpStaked *big.Int
}

type Method string

const (
	SignMsgBatch Method = "signMsgBatch"
)

func (m Method) String() string {
	return string(m)
}

// Context ---------------------------------------------
type Context struct {
	ctx            context.Context
	requestId      string
	availableNodes []string
	approvers      []string
	unApprovers    []string
	electionId     uint64
	stateBatchRoot [32]byte
}

func NewContext() Context {
	return Context{
		ctx: context.Background(),
	}
}

func (c Context) RequestId() string {
	return c.requestId
}

func (c Context) AvailableNodes() []string {
	return c.availableNodes
}
func (c Context) Approvers() []string {
	return c.approvers
}

func (c Context) UnApprovers() []string {
	return c.unApprovers
}

func (c Context) StateBatchRoot() [32]byte {
	return c.stateBatchRoot
}

func (c Context) WithRequestId(requestId string) Context {
	c.requestId = requestId
	return c
}

func (c Context) WithAvailableNodes(nodes []string) Context {
	c.availableNodes = nodes
	return c
}

func (c Context) WithApprovers(nodes []string) Context {
	c.approvers = nodes
	return c
}

func (c Context) WithUnApprovers(nodes []string) Context {
	c.unApprovers = nodes
	return c
}

func (c Context) WithStateBatchRoot(stateBatchRoot [32]byte) Context {
	c.stateBatchRoot = stateBatchRoot
	return c
}
