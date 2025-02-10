package handlers

import (
	"encoding/json"
	"net/http"
)

// Transaction represents a dummy transaction structure
type Transaction struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

// CreateTransaction responds with a dummy transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	dummyTransaction := Transaction{
		ID:          1,
		Amount:      100.50,
		Description: "Test Transaction",
		Category:    "Food",
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dummyTransaction)
}

// GetTransactions responds with a list of dummy transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	dummyTransactions := []Transaction{
		{ID: 1, Amount: 100.50, Description: "Test Transaction 1", Category: "Food"},
		{ID: 2, Amount: 200.75, Description: "Test Transaction 2", Category: "Transport"},
		{ID: 3, Amount: 50.00, Description: "Test Transaction 3", Category: "Entertainment"},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dummyTransactions)
}
