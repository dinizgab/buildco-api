package entity

import (
	"github.com/google/uuid"
)

type Rating struct {
	ID      uuid.UUID
	Grade   int
	Comment string
}
