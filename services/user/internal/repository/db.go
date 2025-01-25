package repository

import "gorm.io/gorm"

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func NewDatabaseConnection(config DatabaseConfig) (*gorm.DB, error)
func MigrateDatabase(db *gorm.DB) error
func CloseConnection(db *gorm.DB) error
