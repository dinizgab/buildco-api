package usecase

import (
	"testing"

	"github.com/dinizgab/buildco-api/internal/user/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (repo *UserRepositoryMock) Create(user *entity.User) (*entity.User, error) {
	args := repo.Called(user)

	return args.Get(0).(*entity.User), nil
}

type UsersUseCaseTestSuite struct {
	suite.Suite
	usecase UsersUseCase
	repo    *UserRepositoryMock
}

func (suite *UsersUseCaseTestSuite) SetupSuite() {
	repo := &UserRepositoryMock{}

	suite.repo = repo
	suite.usecase = NewUsecase(repo)
}

func (suite *UsersUseCaseTestSuite) TestCreateNewUser() {
	t := suite.T()
	newId := uuid.New()
	user := &entity.User{
		Name:     "Gabriel",
		UserName: "dinizgab",
		Email:    "gabriel@gmail.com",
		Password: "gabriel123",
	}

	expected := &entity.User{ID: newId, Name: "Gabriel", UserName: "dinizgab", Email: "gabriel@gmail.com"}
	suite.repo.Mock.On("Create", user).Return(expected)

	newUser, err := suite.usecase.Create(user)
	assert.Nil(t, err)
	assert.NotNil(t, newUser.ID)
	assert.Equal(t, "Gabriel", newUser.Name)
	assert.Equal(t, "dinizgab", newUser.UserName)
	assert.Equal(t, "gabriel@gmail.com", newUser.Email)
}

func (suite *UsersUseCaseTestSuite) TestCreateNewUserEmptyName() {
	t := suite.T()
	user := &entity.User{
		Name: "",
	}

	newUser, err := suite.usecase.Create(user)

	assert.Nil(t, newUser)
	assert.NotNil(t, err)
	assert.Equal(t, "User's name must not be empty!", err.Error())
}

func (suite *UsersUseCaseTestSuite) TestCreateNewUserEmptyUserName() {
	t := suite.T()
	user := &entity.User{
		Name:     "Gabriel",
		UserName: "",
	}

	newUser, err := suite.usecase.Create(user)

	assert.Nil(t, newUser)
	assert.NotNil(t, err)
	assert.Equal(t, "User's username must not be empty!", err.Error())
}

func (suite *UsersUseCaseTestSuite) TestCreateNewUserEmptyEmail() {
	t := suite.T()
	user := &entity.User{
		Name:     "Gabriel",
		UserName: "dinizgab",
		Email:    "",
	}

	newUser, err := suite.usecase.Create(user)

	assert.Nil(t, newUser)
	assert.NotNil(t, err)
	assert.Equal(t, "User's email must not be empty!", err.Error())
}

func (suite *UsersUseCaseTestSuite) TestCreateNewUserEmptyPassword() {
	t := suite.T()
	user := &entity.User{
		Name:     "Gabriel",
		UserName: "dinizgab",
		Email:    "gabriel@gmail.com",
		Password: "",
	}

	newUser, err := suite.usecase.Create(user)

	assert.Nil(t, newUser)
	assert.NotNil(t, err)
	assert.Equal(t, "User's password must not be empty!", err.Error())
}

func TestUsersUseCase(t *testing.T) {
	suite.Run(t, new(UsersUseCaseTestSuite))
}
