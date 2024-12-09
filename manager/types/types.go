package types

import (
	"context"
	"math/big"
)

type SignMsgRequest struct {
	BlockNumber *big.Int `json:"block_number"`
	TxHash      []byte   `json:"tx_hash"`
	TxType      string   `json:"tx_type"`
}

type NodeSignRequest struct {
	Timestamp   int64       `json:"timestamp"`
	Nodes       []string    `json:"nodes"`
	RequestBody interface{} `json:"request_body"`
}

type SignMsgResponse struct {
	Signature     []byte   `json:"signature"`
	G2Point       []byte   `json:"g2_point"`
	L2BlockNumber *big.Int `json:"l2_block_number"`
	Vote          uint8    `json:"vote"`
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
