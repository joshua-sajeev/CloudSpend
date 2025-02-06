package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

type Service interface {
	Close() error
}

type service struct {
	db *sql.DB
}

func New() Service {
	// Construct connection string using environment variables
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	)

	// Open database connection
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Return database service
	return &service{db: db}
}

func (s *service) Close() error {
	log.Printf("Disconnected from database")
	return s.db.Close()
}
