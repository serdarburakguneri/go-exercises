package repository

import (
	"github.com/google/uuid"
	"simple-postgres-client/database"
	"simple-postgres-client/entity"
)

//I think interfaces deserve to be in a separate file, but let's keep it simple for now

type CustomerRepository interface {
	FindById(id uuid.UUID) (*entity.Customer, error)
}

type CustomerRepositoryImpl struct {
	dbClient database.DBClient
}

func NewCustomerRepositoryImpl(dbClient database.DBClient) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{dbClient: dbClient}
}

func (r *CustomerRepositoryImpl) FindById(id uuid.UUID) (*entity.Customer, error) {

	db := r.dbClient.OpenConnection()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM customer WHERE id = $1", id)

	if row == nil {
		return nil, nil
	}

	var name string
	var fetchedId uuid.UUID

	err := row.Scan(&fetchedId, &name)
	if err != nil {
		return nil, err
	}

	return &entity.Customer{
		Id:   id,
		Name: name,
	}, nil
}
