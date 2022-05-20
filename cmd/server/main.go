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
	"go.gofermart/internal/config"
)

func main() {
	wg := &sync.WaitGroup{}
	_, cancel := context.WithCancel(context.Background())

	//load env vars
	env, err := config.LoadConfig()
	if err != nil {
		logrus.Error("Error parse env: ", err)
	}
	wg.Add(1)
	//wait signal SIGTERM
	go waitSignals(cancel, wg)

	r := chi.NewRouter()
	r.Route("/", handlers.Router)
	go http.ListenAndServe(env.Address, r)

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
