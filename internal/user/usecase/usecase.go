package usecase

import (
	"errors"

	"github.com/dinizgab/buildco-api/internal/user/entity"
	"github.com/dinizgab/buildco-api/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
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
	if len(user.Name) == 0 {
		return nil, errors.New("User's name must not be empty!")
	}

	if len(user.UserName) == 0 {
		return nil, errors.New("User's username must not be empty!")
	}

	if len(user.Email) == 0 {
		return nil, errors.New("User's email must not be empty!")
	}

	if len(user.Password) == 0 {
		return nil, errors.New("User's password must not be empty!")
	}
    
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    if err != nil {
        return nil, err
    }
    user.Password = string(hashedPassword)

	newUser, err := uc.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
