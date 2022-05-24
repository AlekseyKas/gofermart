package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

func WaitSignals(cancel context.CancelFunc, wg *sync.WaitGroup) {
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	for {
		sig := <-terminate
		switch sig {
		case os.Interrupt:
			cancel()
			wg.Done()
			logrus.Info("Terminate signal OS!")
			return
		}
	}
}
