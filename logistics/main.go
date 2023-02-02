package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nandohos/go-warehouse-manager/logistics/loaders"
)

func main() {
	log.Println("Starting Logistics Service")

	// Initialize data sources
	ds, err := loaders.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize data sources: %v\n", err)
	}

	router, err := loaders.Inject(ds)
	if err != nil {
		log.Fatalf("Failed to inject data sources: %v\n", err)
	}

	srv := &http.Server{
		Addr:    ":8181",
		Handler: router,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v\n", err)
	}
}
