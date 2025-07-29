package node

import (
	"errors"
	"fmt"

	"github.com/Manta-Network/manta-fp-aggregator/manager/types"

	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

func (n *Node) ProcessMessage() {
	n.log.Info("process websocket message")
	defer n.wg.Done()
	reqChan := make(chan tmtypes.RPCRequest)
	stopChan := make(chan struct{})
	if err := n.wsClient.RegisterResChannel(reqChan, stopChan); err != nil {
		n.log.Error("failed to register request channel with websocket client", "err", err)
		return
	}

	go func() {
		defer func() {
			close(stopChan)
		}()
		for {
			select {
			case rpcReq := <-reqChan:
				reqId := rpcReq.ID.(tdtypes.JSONRPCStringID).String()
				n.log.Info(fmt.Sprintf("receive request method : %s", rpcReq.Method), "reqId", reqId)
				if rpcReq.Method == types.SignMsgBatch.String() {
					if err := n.writeChan(n.signRequestChan, rpcReq); err != nil {
						n.log.Error("failed to write msg to sign channel,channel blocked ", "err", err)
					}
				} else {
					n.log.Error(fmt.Sprintf("unknown rpc request method : %s ", rpcReq.Method))
				}
			case <-n.done:
				return
			}

		}
	}()
}

func (n *Node) writeChan(cache chan tdtypes.RPCRequest, msg tdtypes.RPCRequest) error {
	select {
	case cache <- msg:
		if msg.Method == types.SignMsgBatch.String() {
			n.log.Info("write msg to sign channel successfully")
		}
		return nil
	default:
		return errors.New(msg.Method + " channel blocked,can not write!")
	}
}
