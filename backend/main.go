package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {

	http.HandleFunc("/", greet)

	httpsServer := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	log.Printf("Starting HTTPS server on %s", ":8080")
	if err := httpsServer.ListenAndServeTLS("./cert.pem", "./key.pem"); err != http.ErrServerClosed {
		log.Fatalf("HTTPS server failed: %v", err)
	}
}
