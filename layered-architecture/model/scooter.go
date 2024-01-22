package models

import "github.com/google/uuid"

type Scooter struct {
	Id           uuid.UUID
	SerialNumber string
	Brand        string
	Model        string
	BatteryLevel float32
}
