package service

import (
	"github.com/google/uuid"
	"simple-postgres-client/entity"
	"simple-postgres-client/repository"
)

type CustomerService interface {
	FindById(id uuid.UUID) (*entity.Customer, error)
}

// CustomerServiceImpl I know it looks like a Java code, sorry
type CustomerServiceImpl struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerServiceImpl(customerRepository repository.CustomerRepository) *CustomerServiceImpl {
	return &CustomerServiceImpl{customerRepository: customerRepository}
}

func (c *CustomerServiceImpl) FindById(id uuid.UUID) (*entity.Customer, error) {
	return c.customerRepository.FindById(id)
}
