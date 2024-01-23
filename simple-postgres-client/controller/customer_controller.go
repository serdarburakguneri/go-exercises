package controller

import (
	"fmt"
	"github.com/google/uuid"
	"simple-postgres-client/service"
)

type CustomerController struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

// route: rootUrl/customer/{id} GET

func (c *CustomerController) GetById(id uuid.UUID) {

	fmt.Println("REST Controller received GET request")

	customer, err := c.customerService.FindById(id)

	if err != nil {
		fmt.Println("ouch")
		fmt.Println(err)
	} else if customer == nil {
		fmt.Println("Not Found!")
	} else {
		fmt.Printf("Customer found: id %s name %s\n", id, customer.Name)
	}
}
