package main

import (
	"fmt"
	"github.com/guobingithub/graceful-shut/graceful"
	"github.com/guobingithub/graceful-shut/logger"
	"github.com/guobingithub/graceful-shut/server"
)

func main() {
	logger.Error(fmt.Sprintf("graceful-shut starting..."))
	sigSvr := graceful.NewSignal()

	//gRPC
	go server.StartGrpcServer(sigSvr)

	logger.Info("graceful shutdown start...")
	graceful.StartGracefulShutServer(sigSvr)
}
