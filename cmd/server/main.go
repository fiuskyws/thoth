package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fiuskyws/thoth/src/manager"
	"github.com/fiuskyws/thoth/src/server"
	"go.uber.org/zap"
)

var (
	port = uint(8090)
)

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// TODO:
	// 	- Remove this implementation, MUST NOT use `ReplaceGlobals`
	undo := zap.ReplaceGlobals(l)
	defer undo()

	exit := make(chan os.Signal, 1)

	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)

	m := manager.NewManager()
	g := server.NewGRPC(m, 8090)

	go func() {
		if err := g.Start(); err != nil {
			zap.L().Panic(err.Error())
		}
	}()

	<-exit

	if err := g.Close(); err != nil {
		zap.L().Panic(err.Error())
	}
}
