package repository

import (
	"database/sql"
	_ "embed"

	"github.com/dinizgab/buildco-api/internal/user/entity"
)

var (
	//go:embed sql/create_new_user.sql
	queryCreateNewUser string
)

type UsersRepository interface {
	Create(*entity.User) (*entity.User, error)
}

type usersRepositoryImpl struct {
	DB *sql.DB
}

func NewRepository(DB *sql.DB) UsersRepository {
	return &usersRepositoryImpl{
		DB: DB,
	}
}

func (repo *usersRepositoryImpl) Create(user *entity.User) (*entity.User, error) {
	newUser := new(entity.User)
	args := []interface{}{user.Name, user.UserName, user.Email, user.Password}

	err := repo.DB.QueryRow(queryCreateNewUser, args...).Scan(&newUser.ID, &newUser.Name, &newUser.UserName, &newUser.Email)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
