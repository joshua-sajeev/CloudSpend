package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"cloudspend/internal/database"
	"cloudspend/internal/server"
)

func main() {
	// Initialize Database
	db := database.ConnectDB()
	defer database.CloseDB(db)

	// Start HTTP server
	router := server.SetupRouter(db)
	serverAddr := ":8080"

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	go func() {
		log.Println("Server running on", serverAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-quit
	log.Println("Shutting down server...")

	// Give 5 seconds to close existing connections
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
