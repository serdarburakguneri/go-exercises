package main

import "dependency-injection-with-wire/dto"

func main() {
	scooterController := InitializeScooterController()
	request := dto.CreateScooterRequestDTO{
		SerialNumber: "1212ABCD",
	}
	scooterController.Register(request)
}
