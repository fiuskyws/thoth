package main

import (
	"os"
	"os/signal"
	"syscall"

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

	<-exit
}
