## Gracious Shutdown
```go
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
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	idleConnsClosed := make(chan struct{})
	signalCh := make(chan os.Signal, 1)

	// Listen for SIGINT (Ctrl+C), SIGTSTP (Ctrl+Z), and SIGTERM
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTSTP, syscall.SIGTERM)

	go func() {
		<-signalCh
		fmt.Println("\nReceived shutdown signal, stopping server...")

		// Graceful shutdown
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}

		close(idleConnsClosed)
	}()

	log.Println("Server started on :8080")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe error: %v", err)
	}

	<-idleConnsClosed
	fmt.Println("Server has shut down gracefully.")
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mi chiami Geusue! The time is: %s", time.Now().Format(time.RFC1123))
}

```
```
```
