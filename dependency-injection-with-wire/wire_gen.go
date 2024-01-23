// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"dependency-injection-with-wire/controller"
	"dependency-injection-with-wire/service"
)

// Injectors from wire.go:

func InitializeScooterController() *controller.ScooterController {
	scooterServiceNewImplementation := service.NewScooterServiceNewImplementation()
	scooterController := controller.NewScooterController(scooterServiceNewImplementation)
	return scooterController
}