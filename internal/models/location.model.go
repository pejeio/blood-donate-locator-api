package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
