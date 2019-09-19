package graceful

import (
	"github.com/guobingithub/graceful-shut/logger"
	"os"
	"os/signal"
	"time"
)

const (
	GracefulShut_PendingTime = 5
)

func StartGracefulShutServer(sigSvr *SignalServer) {
	logger.Info("gracefulShutServer running...")
	sigSvr.ListenOn(os.Interrupt)
	sigSvr.ListenOn(os.Kill)

	c := make(chan os.Signal)
	signal.Notify(c)
	sig := <-c

	sigSvr.GracefulHandle(sig)

	logger.Info("gracefulShutServer, all shutdown, prepare to stop...")
	<-time.After(time.Second * GracefulShut_PendingTime)
	logger.Info("gracefulShutServer, all shutdown, succeed to stop...")
	os.Exit(1)
}
