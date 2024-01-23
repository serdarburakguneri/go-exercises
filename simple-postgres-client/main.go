package main

import (
	"github.com/google/uuid"
	"log"
	"simple-postgres-client/controller"
	"simple-postgres-client/database"
	"simple-postgres-client/repository"
	"simple-postgres-client/service"
)

func populateDB(dbClient *database.PostgresClient) {
	db := dbClient.OpenConnection()
	defer db.Close()

	// Create a new table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS customer (
			id UUID PRIMARY KEY,
			name VARCHAR(255)		
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
			log.Fatal(err)
		}
	}()

	_, err = tx.Exec("INSERT INTO customer (id, name) VALUES ($1, $2)", "56ff36c1-b913-45f6-8fc3-d238569b6995", "serdar")
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	postgresClient := &database.PostgresClient{}

	//populateDB(postgresClient)

	customerRepo := repository.NewCustomerRepositoryImpl(postgresClient)
	customerService := service.NewCustomerServiceImpl(customerRepo)
	customerController := controller.NewCustomerController(customerService)

	customerController.GetById(uuid.MustParse("56ff36c1-b913-45f6-8fc3-d238569b6995"))

}
