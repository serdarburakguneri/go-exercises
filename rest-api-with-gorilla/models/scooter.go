package models

import "github.com/google/uuid"

type Scooter struct {
	Id           uuid.UUID `json:"id"`
	SerialNumber string    `json:"serialNumber"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	BatteryLevel float32   `json:"batteryLevel"`
}
