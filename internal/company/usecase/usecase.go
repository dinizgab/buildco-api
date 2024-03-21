package company

import (
	company "github.com/dinizgab/buildco-api/internal/company/repository"
	"github.com/jackc/pgx/v5"
)

type CompanyUseCase interface {

}

type companyUseCaseImpl struct {
    repo company.CompanyRepository
}

func NewUsecase(db *pgx.Conn) CompanyUseCase {
    return &companyUseCaseImpl{
        repo: company.NewRepository(db),
    }
}
