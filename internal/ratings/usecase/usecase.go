package usecase

import (
	"errors"
	"fmt"

	"github.com/dinizgab/buildco-api/internal/ratings/entity"
	"github.com/dinizgab/buildco-api/internal/ratings/repository"
	"github.com/google/uuid"
)

type RatingsUsecase interface {
	Create(string, *entity.Rating) (*entity.Rating, error)
}

type ratingsUseCaseImpl struct {
	repo repository.RatingsRepository
}

func NewUsecase(repo repository.RatingsRepository) RatingsUsecase {
	return &ratingsUseCaseImpl{
		repo: repo,
	}
}

func (uc *ratingsUseCaseImpl) Create(companyId string, rating *entity.Rating) (*entity.Rating, error) {
	parsedCompanyId, err := uuid.Parse(companyId)
	if err != nil {
		return nil, err
	}

	if rating.Grade < 1 || rating.Grade > 5 {
		return nil, errors.New(fmt.Sprintf("Invalid grade value: %d", rating.Grade))
	}

	if len(rating.Comment) == 0 {
		return nil, errors.New("Comment must not be empty!")
	}

	newRating, err := uc.repo.Create(parsedCompanyId, rating)
	if err != nil {
		return nil, err
	}

	return newRating, nil
}
