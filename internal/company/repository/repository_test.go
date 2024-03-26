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

func (suite *CompanyRepositoryTestSuite) TestFindCompanyById() {
    t := suite.T()
    companyId := "244c423a-930f-42a3-837f-c99102d27339"

    result, err := suite.repository.FindById(companyId)
    assert.Nil(t, err)
    assert.Equal(t, "Test co.1", result.Name)
    assert.Equal(t, "testco1@gmail.com", result.Email)
    assert.Equal(t, "1234-1234", result.Phone)
}

func TestEventsRepository(t *testing.T) {
	suite.Run(t, new(CompanyRepositoryTestSuite))
}
