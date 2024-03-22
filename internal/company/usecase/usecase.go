package usecase 

import (
	"database/sql"

	"github.com/dinizgab/buildco-api/internal/company/repository"
	"github.com/dinizgab/buildco-api/internal/company/entity"
)

type CompanyUseCase interface {
    Create(*entity.Company) (*entity.Company, error)
}

type companyUseCaseImpl struct {
    repo repository.CompanyRepository
}

func NewUsecase(db *sql.DB) CompanyUseCase {
    return &companyUseCaseImpl{
        repo: repository.NewRepository(db),
    }
}

func (uc *companyUseCaseImpl) Create(company *entity.Company) (*entity.Company, error) {
    return nil, nil
}
