package usecase

import (
	"github.com/dinizgab/buildco-api/internal/company/entity"
	"github.com/dinizgab/buildco-api/internal/company/repository"
)

type CompanyUseCase interface {
	Create(*entity.Company) (*entity.Company, error)
}

type companyUseCaseImpl struct {
	repo repository.CompanyRepository
}

func NewUsecase(repo repository.CompanyRepository) CompanyUseCase {
	return &companyUseCaseImpl{
		repo: repo,
	}
}

func (uc *companyUseCaseImpl) Create(company *entity.Company) (*entity.Company, error) {
	newCompany, err := uc.repo.Create(company)
	if err != nil {
		return nil, err
	}

	return newCompany, nil
}
