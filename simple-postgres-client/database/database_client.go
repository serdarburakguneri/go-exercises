package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DBClient interface {
	OpenConnection() *sql.DB
}

type PostgresClient struct {
}

func (pc *PostgresClient) OpenConnection() *sql.DB {
	// imagine we get these credential from environment variables
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
