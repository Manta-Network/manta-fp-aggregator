package manager

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"testing"
)

func Test1(t *testing.T) {
	cli, err := client.NewClientFromNode("https://rpc-euphrates.devnet.babylonlabs.io:443")
	if err != nil {
		fmt.Printf("Error creating babylon client: %v", err)
	}
	height := int64(349588)
	block, err := cli.Block(context.Background(), &height)
	if err != nil {
		fmt.Println("failed to get babylon block", "err", err)
	}

	fmt.Println(block.Block.Height)
}
