package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Close() error
}

type service struct {
	db *sql.DB
}

var (
	dbInstance *service
	once       sync.Once
)

func New() Service {
	once.Do(func() {
		// Load .env file explicitly
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: No .env file found or unable to load it")
		}

		// Read environment variables
		dbUser := os.Getenv("DBUSER")
		dbPassword := os.Getenv("PASSWORD")
		dbHost := os.Getenv("HOST")
		dbPort := os.Getenv("DBPORT")
		dbName := os.Getenv("DBNAME")

		// Debugging: Print environment variables to check
		fmt.Printf("DB Config: USER=%s, PASSWORD=%s, HOST=%s, DBPORT=%s, DBNAME=%s\n",
			dbUser, dbPassword, dbHost, dbPort, dbName)

		// Construct connection string
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbUser, dbPassword, dbHost, dbPort, dbName,
		)

		// Open database connection
		db, err := sql.Open("pgx", connStr)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Assign the instance
		dbInstance = &service{db: db}
	})

	return dbInstance
}
func AddTransaction(category, title, transactionType string, amount float64, date time.Time) error {
	query := `INSERT INTO transactions (category, title, type, amount, date) VALUES ($1, $2, $3, $4, $5)`

	_, err := dbInstance.db.Exec(query, category, title, transactionType, amount, date)
	if err != nil {
		log.Printf("Failed to insert transaction: %s", err)
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	log.Println("Transaction inserted successfully")
	return nil
}

func (s *service) Close() error {
	log.Printf("Disconnected from database")
	return s.db.Close()
}
