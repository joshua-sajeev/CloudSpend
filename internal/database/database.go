package database

import (
	"database/sql"
	"fmt"

	//for connecting to db
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "joshua"
	dbname   = "cloudspend"
)

var (
	DbClient *sql.DB
)

func Dbinit() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	DbClient, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	// defer dbClient.Close()
	fmt.Println((DbClient))
	fmt.Println("Successfully connected!")
}
