package usecase

import (
	"github.com/dinizgab/buildco-api/internal/user/entity"
	"github.com/dinizgab/buildco-api/internal/user/repository"
)

type UsersUseCase interface {
	Create(*entity.User) (*entity.User, error)
}

type usersUsecaseImpl struct {
	repo repository.UsersRepository
}

func NewUsecase(repo repository.UsersRepository) UsersUseCase {
	return &usersUsecaseImpl{
		repo: repo,
	}
}

func (uc *usersUsecaseImpl) Create(user *entity.User) (*entity.User, error) {
	newUser, err := uc.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
