package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"go.gofermart/cmd/server/handlers"
	"go.gofermart/cmd/server/storage"
	"go.gofermart/internal/config"
)

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	args, err := config.TerminateFlags()
	if err != nil {
		logrus.Error("Error setting args: ", err)
	}
	storage.IDB = &storage.DB
	storage.IDB.InitDB(ctx, args.DatabaseURL)
	wg.Add(1)
	go waitSignals(cancel, wg)

	r := chi.NewRouter()
	r.Route("/", handlers.Router)
	go http.ListenAndServe(args.Address, r)

	wg.Wait()
}

func waitSignals(cancel context.CancelFunc, wg *sync.WaitGroup) {
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
