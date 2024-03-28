package usecase

import (
	"github.com/dinizgab/buildco-api/internal/company/entity"
	"github.com/dinizgab/buildco-api/internal/company/repository"
)

type CompanyUseCase interface {
	Create(*entity.Company) (*entity.Company, error)
	FindAll() ([]*entity.Company, error)
	FindById(string) (*entity.Company, error)
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

func (uc *companyUseCaseImpl) FindAll() ([]*entity.Company, error) {
	companies, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (uc *companyUseCaseImpl) FindById(id string) (*entity.Company, error) {
	resultCompany, err := uc.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return resultCompany, nil
}
