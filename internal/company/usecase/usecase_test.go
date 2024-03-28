package usecase

import (
	"testing"

	"github.com/dinizgab/buildco-api/internal/company/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CompanyRepositoryMock struct {
	mock.Mock
}

func (repo *CompanyRepositoryMock) Create(company *entity.Company) (*entity.Company, error) {
	args := repo.Called(company)

	return args.Get(0).(*entity.Company), nil
}

func (repo *CompanyRepositoryMock) FindAll() ([]*entity.Company, error) {
    args := repo.Called()

	return args.Get(0).([]*entity.Company), nil
}

func (repo *CompanyRepositoryMock) FindById(id string) (*entity.Company, error) {
	args := repo.Called(id)

	return args.Get(0).(*entity.Company), nil
}

type CompanyUsecaseTestSuite struct {
	suite.Suite
	usecase CompanyUseCase
	repo    *CompanyRepositoryMock
}

func (suite *CompanyUsecaseTestSuite) SetupSuite() {
	repo := &CompanyRepositoryMock{}

	suite.usecase = NewUsecase(repo)
	suite.repo = repo
}

func (suite *CompanyUsecaseTestSuite) TestCreateCompany() {
	t := suite.T()

	newCompany := &entity.Company{
		Name:  "Build Co. 1",
		Email: "buildco1@gmail.com",
		Phone: "837990-2345",
	}
	newId := uuid.New()

	expected := &entity.Company{ID: newId, Name: "Culto 1", Email: "buildco1@gmail.com", Phone: "837990-2345"}
	suite.repo.Mock.On("Create", newCompany).Return(expected)

	result, err := suite.usecase.Create(newCompany)

	assert.Nil(t, err)
	assert.Equal(t, newId, result.ID)
	assert.Equal(t, expected.Name, result.Name)
	assert.Equal(t, expected.Email, result.Email)
	assert.Equal(t, expected.Phone, result.Phone)
}

func (suite *CompanyUsecaseTestSuite) TestFindAll() {
	t := suite.T()

	expected := []*entity.Company{{Name: "Co2"}, {Name: "Co1"}}
	suite.repo.Mock.On("FindAll").Return(expected)

	result, err := suite.usecase.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
}

func (suite *CompanyUsecaseTestSuite) TestFindById() {
	t := suite.T()
	companyId := "8db46e78-bf5b-46fb-8768-7e1fc457e5a7"
	parsedCompanyId, _ := uuid.Parse(companyId)

	expected := &entity.Company{ID: parsedCompanyId, Name: "Buildco 1", Email: "buildco1@gmail.com", Phone: "1234-1234"}
	suite.repo.Mock.On("FindById", companyId).Return(expected)

	result, err := suite.usecase.FindById(companyId)

	assert.Nil(t, err)
	assert.Equal(t, parsedCompanyId, result.ID)
	assert.Equal(t, "Buildco 1", result.Name)
}

func TestCompanyUseCase(t *testing.T) {
	suite.Run(t, new(CompanyUsecaseTestSuite))
}
