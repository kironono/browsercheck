package server

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/kironono/browsercheck/internal/model"
)

func Run(config *model.Config) error {
	log.Printf("Config: %v\n", config)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr: config.ServerSettings.ListenAddress,
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	log.Printf("starting server at %s\n", config.ServerSettings.ListenAddress)
	server.ListenAndServe()

	return nil
}
