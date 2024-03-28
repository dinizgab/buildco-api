package entity

import (
	"time"

	"github.com/dinizgab/buildco-api/internal/ratings/entity"
	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Phone     string           `json:"phone_number"`
	Ratings   []*entity.Rating `json:"ratings"`
	CreatedAt time.Time        `json:"created_at"`
}
