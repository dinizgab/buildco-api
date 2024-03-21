package company

import (
	"database/sql"

	company "github.com/dinizgab/buildco-api/internal/company/repository"
)

type CompanyUseCase interface {

}

type companyUseCaseImpl struct {
    repo company.CompanyRepository
}

func NewUsecase(db *sql.DB) CompanyUseCase {
    return &companyUseCaseImpl{
        repo: company.NewRepository(db),
    }
}
