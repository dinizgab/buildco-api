package entity

import (
	"time"

	"github.com/dinizgab/buildco-api/internal/ratings/entity"
	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Phone     string
	Ratings   []*entity.Rating
	CreatedAt time.Time
}
