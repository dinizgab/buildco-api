package repository

import (
	"context"
	"log"
	"testing"

	"github.com/dinizgab/buildco-api/internal/company/entity"
	"github.com/dinizgab/buildco-api/utils/containers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CompanyRepositoryTestSuite struct {
	suite.Suite
	pgContainer *containers.PostgresContainer
	repository  CompanyRepository
	ctx         context.Context
}

func (suite *CompanyRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := containers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.pgContainer = pgContainer
	suite.repository = NewRepository(pgContainer.DBConn)
}

func (suite *CompanyRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating container: %s", err)
	}
}

func (suite *CompanyRepositoryTestSuite) TestCreateNewCompany() {
	t := suite.T()
	newCompany := &entity.Company{
		Name:  "Build Co. 1",
		Email: "buildco1@gmail.com",
		Phone: "837990-2345",
	}

	result, err := suite.repository.Create(newCompany)
	assert.Nil(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Build Co. 1", result.Name)
	assert.Equal(t, "buildco1@gmail.com", result.Email)
	assert.Equal(t, "837990-2345", result.Phone)
}

func (suite *CompanyRepositoryTestSuite) TestFindByIdNoRatings() {
	t := suite.T()
	companyId := "8db46e78-bf5b-46fb-8768-7e1fc457e5a7"

	result, err := suite.repository.FindById(companyId)
	assert.Nil(t, err)
	assert.Equal(t, "Test co.1", result.Name)
	assert.Equal(t, "testco1@gmail.com", result.Email)
	assert.Equal(t, "1234-1234", result.Phone)
	assert.Equal(t, 0, len(result.Ratings))
}

func (suite *CompanyRepositoryTestSuite) TestFindByIdWithRatings() {
	t := suite.T()
	companyId := "124f7323-ee68-4eb6-9509-84eb966cc5cf"

	result, err := suite.repository.FindById(companyId)
	assert.Nil(t, err)
	assert.Equal(t, "Test co.2", result.Name)
	assert.Equal(t, "testco2@gmail.com", result.Email)
	assert.Equal(t, "4321-4321", result.Phone)
	assert.Equal(t, 3, len(result.Ratings))
}

func TestEventsRepository(t *testing.T) {
	suite.Run(t, new(CompanyRepositoryTestSuite))
}
