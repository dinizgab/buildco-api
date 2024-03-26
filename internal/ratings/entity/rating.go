package entity

import (
	"github.com/dinizgab/buildco-api/internal/company/entity"
	"github.com/google/uuid"
)

type Rating struct {
	ID      uuid.UUID
	Grade   int
	Comment string
	Company *entity.Company
}
