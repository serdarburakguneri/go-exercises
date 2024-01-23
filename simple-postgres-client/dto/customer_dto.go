package dto

import "github.com/google/uuid"

type CustomerDTO struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
