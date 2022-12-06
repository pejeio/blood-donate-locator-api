package types

import (
	"time"

	"github.com/google/uuid"
)

type Coordinates struct {
	Latitude  float32 `json:"lat" validate:"number"`
	Longitude float32 `json:"lng" validate:"number"`
}

type Address struct {
	Street     string `gorm:"index" json:"street,omitempty"`
	Number     int32  `json:"number,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country,omitempty"`
}

type Location struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string      `gorm:"index;not null" json:"name"`
	Address     Address     `gorm:"embedded" json:"address,omitempty"`
	Coordinates Coordinates `gorm:"embedded" json:"coordinates"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type CreateLocationRequest struct {
	Name        string      `json:"name" validate:"required"`
	Coordinates Coordinates `json:"coordinates" validate:"required"`
	Address     Address     `json:"address"`
}

type PaginationRequest struct {
	Page  string `query:"page"`
	Limit string `query:"limit"`
}
