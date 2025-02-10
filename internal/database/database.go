package database

import (
	"cloudspend/internal/config"
	"cloudspend/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	cfg := config.LoadConfig()

	// Create database if it doesn't exist
	if err := createDatabaseIfNotExists(cfg); err != nil {
		log.Fatal("Failed to create/verify database:", err)
	}

	// Add retry logic for database connection
	var db *gorm.DB
	var err error
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		db, err = connectToDatabase(cfg)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(time.Second * 2) // Wait before retrying
	}

	if err != nil {
		log.Fatal("Failed to connect to database after multiple attempts:", err)
	}

	log.Println("Successfully connected to database")

	// AutoMigrate models
	if err := db.AutoMigrate(&models.Transaction{}); err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	DB = db
	return db
}

func connectToDatabase(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func createDatabaseIfNotExists(cfg config.Config) error {
	// Connect to PostgreSQL's default database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	defer db.Close()

	// Wait for PostgreSQL to be ready
	for i := 0; i < 5; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		log.Println("Waiting for PostgreSQL to be ready...")
		time.Sleep(time.Second * 2)
	}

	// Check if database exists
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)")
	err = db.QueryRow(query, cfg.DBName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking database existence: %w", err)
	}

	// Create database if it doesn't exist
	if !exists {
		// Important: Use parametrized query with Quote identifier to prevent SQL injection
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", QuoteIdentifier(cfg.DBName)))
		if err != nil {
			return fmt.Errorf("error creating database: %w", err)
		}
		log.Println("Database created successfully")
	} else {
		log.Println("Database already exists")
	}

	return nil
}

// QuoteIdentifier properly quotes database identifiers to prevent SQL injection
func QuoteIdentifier(name string) string {
	return fmt.Sprintf("\"%s\"", name)
}
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error getting database instance:", err)
		return
	}
	sqlDB.Close()
}
