package entity

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID    uuid.UUID
	Name  string
	Email string
	Phone string
    CreatedAt time.Time
}
