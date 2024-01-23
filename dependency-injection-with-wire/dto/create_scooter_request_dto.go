package dto

type CreateScooterRequestDTO struct {
	SerialNumber string `json:"serialNumber"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
}
