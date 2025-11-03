package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aboagye-Dacosta/shopBackend/logger"
)

func (app *application) Serve() {
	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.port),
		Handler:      app.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Channel to listen for interrupt or terminate signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Server started on PORT: %d\n", app.port)
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := serve.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v\n", err)
	}

	// Close the logger file
	if err := logger.Close(); err != nil {
		log.Printf("Error closing log file: %v\n", err)
	}

	log.Println("Server stopped")
}
