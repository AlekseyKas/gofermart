package main

import (
	"context"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"go.gofermart/cmd/server/handlers"
	"go.gofermart/cmd/server/storage"
	"go.gofermart/internal/app"
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
	go app.WaitSignals(cancel, wg)

	r := chi.NewRouter()
	r.Route("/", handlers.Router)
	go http.ListenAndServe(args.Address, r)

	wg.Wait()
}
