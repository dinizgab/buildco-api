package repository

import (
	"context"
	"log"
	"testing"

	rating "github.com/dinizgab/buildco-api/internal/ratings/entity"
	"github.com/dinizgab/buildco-api/utils/containers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RatingsRepositoryTestSuite struct {
	suite.Suite
	pgContainer containers.PostgresContainer
	repository  RatingsRepository
	ctx         context.Context
}

func (suite *RatingsRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := containers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.pgContainer = *pgContainer
	suite.repository = NewRepository(pgContainer.DBConn)
}

func (suite *RatingsRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating container: %s", err)
	}
}

func (suite *RatingsRepositoryTestSuite) TestCreateNewRating() {
	t := suite.T()
	rating := &rating.Rating{
		Grade:   3,
		Comment: "Test rating",
	}

	uuid, _ := uuid.Parse("124f7323-ee68-4eb6-9509-84eb966cc5cf")

	newRating, err := suite.repository.Create(uuid, rating)

	assert.Nil(t, err)
	assert.NotNil(t, newRating.ID)
	assert.Equal(t, 3, newRating.Grade)
	assert.Equal(t, "Test rating", newRating.Comment)
}

func TestRatingsRepository(t *testing.T) {
	suite.Run(t, new(RatingsRepositoryTestSuite))
}
