package service

import "dependency-injection-with-wire/dto"

type ScooterService interface {
	Register(request dto.CreateScooterRequestDTO) dto.ScooterDTO
}

type ScooterServiceNewImplementation struct {
}

func NewScooterServiceNewImplementation() *ScooterServiceNewImplementation {
	return &ScooterServiceNewImplementation{}
}

func (*ScooterServiceNewImplementation) Register(request dto.CreateScooterRequestDTO) dto.ScooterDTO {
	return dto.ScooterDTO{SerialNumber: request.SerialNumber}
}

type ScooterServiceLegacyImplementation struct {
}

func (*ScooterServiceLegacyImplementation) Register(request dto.CreateScooterRequestDTO) dto.ScooterDTO {
	return dto.ScooterDTO{SerialNumber: request.SerialNumber}
}
