package main

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/log"
)

var (
	GitCommit = ""
	GitDate   = ""
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
	app := newCli(GitCommit, GitDate)
	if err := app.RunContext(context.Background(), os.Args); err != nil {
		log.Error("application failed", "err", err)
		os.Exit(1)
	}
}
