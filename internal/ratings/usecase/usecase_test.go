package usecase

import (
	"testing"

	company "github.com/dinizgab/buildco-api/internal/company/entity"
	rating "github.com/dinizgab/buildco-api/internal/ratings/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RatingsRepositoryMock struct {
	mock.Mock
}

func (repo *RatingsRepositoryMock) Create(companyId uuid.UUID, newRating *rating.Rating) (*rating.Rating, error) {
	args := repo.Called(companyId, newRating)

	return args.Get(0).(*rating.Rating), nil
}

type RatingsUseCaseTestSuite struct {
	suite.Suite
	usecase RatingsUsecase
	repo    *RatingsRepositoryMock
}

func (suite *RatingsUseCaseTestSuite) SetupSuite() {
	repo := &RatingsRepositoryMock{}

	suite.usecase = NewUsecase(repo)
	suite.repo = repo
}

func (suite *RatingsUseCaseTestSuite) TestCreateRating() {
	t := suite.T()
	newRating := &rating.Rating{
		Grade:   3,
		Comment: "Test comment 1",
	}
	newId := uuid.New()
	companyId := uuid.NewString()
	parsedCompanyId, _ := uuid.Parse(companyId)

	expected := &rating.Rating{ID: newId, Grade: 3, Comment: "Test comment 1", Company: &company.Company{ID: parsedCompanyId}}
	suite.repo.Mock.On("Create", parsedCompanyId, newRating).Return(expected)

	newRating, err := suite.usecase.Create(companyId, newRating)
	assert.Nil(t, err)
	assert.Equal(t, 3, newRating.Grade)
	assert.Equal(t, "Test comment 1", newRating.Comment)
	assert.Equal(t, parsedCompanyId, newRating.Company.ID)
}

func (suite *RatingsUseCaseTestSuite) TestCreateRatingBiggestGrade() {
	t := suite.T()
	newRating := &rating.Rating{
		Grade:   5,
		Comment: "Test comment 2",
	}
	newId := uuid.New()
	companyId := uuid.NewString()
	parsedCompanyId, _ := uuid.Parse(companyId)

	expected := &rating.Rating{ID: newId, Grade: 5, Comment: "Test comment 2", Company: &company.Company{ID: parsedCompanyId}}
	suite.repo.Mock.On("Create", parsedCompanyId, newRating).Return(expected)

	newRating, err := suite.usecase.Create(companyId, newRating)
	assert.Nil(t, err)
	assert.Equal(t, 5, newRating.Grade)
	assert.Equal(t, "Test comment 2", newRating.Comment)
	assert.Equal(t, parsedCompanyId, newRating.Company.ID)
}

func (suite *RatingsUseCaseTestSuite) TestCreateRatingSmallestGrade() {
	t := suite.T()
	newRating := &rating.Rating{
		Grade:   1,
		Comment: "Test comment 3",
	}
	newId := uuid.New()
	companyId := uuid.NewString()
	parsedCompanyId, _ := uuid.Parse(companyId)

	expected := &rating.Rating{ID: newId, Grade: 1, Comment: "Test comment 3", Company: &company.Company{ID: parsedCompanyId}}
	suite.repo.Mock.On("Create", parsedCompanyId, newRating).Return(expected)

	newRating, err := suite.usecase.Create(companyId, newRating)
	assert.Nil(t, err)
	assert.Equal(t, 1, newRating.Grade)
	assert.Equal(t, "Test comment 3", newRating.Comment)
	assert.Equal(t, parsedCompanyId, newRating.Company.ID)
}

func (suite *RatingsUseCaseTestSuite) TestCreateRatingEqualsZero() {
    t := suite.T()
	companyId := uuid.NewString()
    newRating := &rating.Rating{
        Grade: 0,
        Comment: "test",
    }
    
    newRating, err := suite.usecase.Create(companyId, newRating)

    assert.Nil(t, newRating)
    assert.NotNil(t, err)
    assert.Equal(t, "Invalid grade value: 0", err.Error())
}

func (suite *RatingsUseCaseTestSuite) TestCreateRatingAboveMax() {
    t := suite.T()
	companyId := uuid.NewString()
    newRating := &rating.Rating{
        Grade: 6,
        Comment: "test",
    }
    
    newRating, err := suite.usecase.Create(companyId, newRating)

    assert.Nil(t, newRating)
    assert.NotNil(t, err)
    assert.Equal(t, "Invalid grade value: 6", err.Error())
}

func (suite *RatingsUseCaseTestSuite) TestCreateRatingEmptyComment() {
    t := suite.T()
	companyId := uuid.NewString()
    newRating := &rating.Rating{
        Grade: 4,
    }
    
    newRating, err := suite.usecase.Create(companyId, newRating)

    assert.Nil(t, newRating)
    assert.NotNil(t, err)
    assert.Equal(t, "Comment must not be empty!", err.Error())
}

func TestRatingsUseCase(t *testing.T) {
	suite.Run(t, new(RatingsUseCaseTestSuite))
}
