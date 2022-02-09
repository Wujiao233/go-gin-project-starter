package main

import (
	"fmt"
	router "git.caimi-inc.com/stanlee/approval-chain-go/src"
	"git.caimi-inc.com/stanlee/approval-chain-go/src/utils/config"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	g errgroup.Group
)

func main() {
	r := router.InitRouter()

	g.Go(func() error {
		return r.Run(fmt.Sprintf(":%s", config.Conf.Server.ServerPort))
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
