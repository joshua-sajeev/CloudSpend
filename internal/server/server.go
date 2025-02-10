package server

import (
	"cloudspend/internal/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// Transaction Routes
	router.HandleFunc("/transactions", handlers.CreateTransaction).Methods("GET")
	router.HandleFunc("/transactions", handlers.GetTransactions).Methods("POST")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	return router
}
