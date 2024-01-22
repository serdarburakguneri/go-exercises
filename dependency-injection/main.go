package main

import "fmt"

//customer_service.go

type CustomerService interface {
	save()
}

//customer_service_impl.go

type CustomerServiceImpl struct {
}

func (*CustomerServiceImpl) save() {
	fmt.Println("CustomerServiceImpl.save() called")
}

//legacy_customer_service_impl.go

type LegacyCustomerServiceImpl struct {
}

func (*LegacyCustomerServiceImpl) save() {
	fmt.Println("LegacyCustomerServiceImpl.save() called")
}

//customer_handler.go

type CustomerHandler struct {
	customerService CustomerService
}

func NewCustomerHandler(service CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: service,
	}
}

func (handler *CustomerHandler) save() {
	handler.customerService.save()
}

//main.go

func main() {
	customerService := &CustomerServiceImpl{}
	customerHandler1 := NewCustomerHandler(customerService)
	customerHandler1.save()

	legacyCustomerService := &LegacyCustomerServiceImpl{}
	customerHandler2 := NewCustomerHandler(legacyCustomerService)
	customerHandler2.save()
}
