package controller

import (
	"dependency-injection-with-wire/dto"
	"dependency-injection-with-wire/service"
	"fmt"
)

type ScooterController struct {
	scooterService service.ScooterService
}

func NewScooterController(scooterService service.ScooterService) *ScooterController {
	return &ScooterController{scooterService: scooterService}
}

func (c *ScooterController) Register(request dto.CreateScooterRequestDTO) {
	fmt.Println("Received register request")
	//scooter := c.scooterService.Register(request)
	//fmt.Println(scooter)
	fmt.Println("Returned registered scooter object")
}
