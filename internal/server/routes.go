package server

import (
	"cloudspend-backend/internal/database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", s.ProcessTransaction)
	// Wrap the mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) ProcessTransaction(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract form values
	transactionType := r.FormValue("transactionType")
	category := r.FormValue("category")
	title := r.FormValue("title")
	amountStr := r.FormValue("amount")
	dateStr := r.FormValue("date") // Ensure your form includes a date field

	// Convert amount to float64
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount format", http.StatusBadRequest)
		return
	}

	// Parse date
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Log received data
	log.Printf("Received: Type=%s, Category=%s, Title=%s, Amount=%.2f, Date=%s",
		transactionType, category, title, amount, date.Format("2006-01-02"))

	database.AddTransaction(category, title, transactionType, amount, date)
	// Prepare JSON response
	resp := map[string]interface{}{
		"message":         "Transaction recorded successfully!",
		"transactionType": transactionType,
		"category":        category,
		"title":           title,
		"amount":          amount,
		"date":            date.Format("2006-01-02"),
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
