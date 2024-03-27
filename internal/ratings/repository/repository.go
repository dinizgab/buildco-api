package repository

import (
	"database/sql"
	_ "embed"

	rating "github.com/dinizgab/buildco-api/internal/ratings/entity"
	"github.com/google/uuid"
)

var (
	//go:embed sql/create_new_rating.sql
	queryCreateNewRating string
)

type RatingsRepository interface {
	Create(uuid.UUID, *rating.Rating) (*rating.Rating, error)
}

type ratingsRepositoryImpl struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) RatingsRepository {
	return &ratingsRepositoryImpl{
		DB: db,
	}
}

func (repo *ratingsRepositoryImpl) Create(companyId uuid.UUID, rating *rating.Rating) (*rating.Rating, error) {
	args := []interface{}{rating.Grade, rating.Comment, companyId}

	err := repo.DB.QueryRow(queryCreateNewRating, args...).Scan(&rating.ID, &rating.Grade, &rating.Comment)
	if err != nil {
		return nil, err
	}

	return rating, nil
}
