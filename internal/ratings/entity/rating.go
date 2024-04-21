package entity

import (
	"time"

	"github.com/google/uuid"
)

type Rating struct {
	ID        uuid.UUID `json:"id"`
	Grade     int       `json:"grade"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"commentDate"`
}
