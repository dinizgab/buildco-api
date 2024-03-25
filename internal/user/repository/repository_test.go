package repository

import (
	"context"
	"log"
	"testing"

	"github.com/dinizgab/buildco-api/internal/user/entity"
	"github.com/dinizgab/buildco-api/utils/containers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UsersRepositoryTestSuite struct {
	suite.Suite
	pgContainer *containers.PostgresContainer
	repo        UsersRepository
	ctx         context.Context
}

func (suite *UsersRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := containers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.pgContainer = pgContainer
	suite.repo = NewRepository(suite.pgContainer.DBConn)
}

func (suite *UsersRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating container: %s", err)
	}
}

func (suite *UsersRepositoryTestSuite) TestCreateNewUser() {
	t := suite.T()
	user := &entity.User{
		Name:     "Gabriel",
		UserName: "dinizgab",
		Email:    "gabriel@gmail.com",
		Password: "gabriel123",
	}

	newUser, err := suite.repo.Create(user)

	assert.Nil(t, err)
	assert.NotNil(t, newUser.ID)
	assert.Equal(t, "Gabriel", newUser.Name)
	assert.Equal(t, "dinizgab", newUser.UserName)
	assert.Equal(t, "gabriel@gmail.com", newUser.Email)
}

func TestUsersRepository(t *testing.T) {
	suite.Run(t, new(UsersRepositoryTestSuite))
}
