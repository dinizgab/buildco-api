package repository

import (
	"database/sql"
	_ "embed"

	"github.com/dinizgab/buildco-api/internal/company/entity"
)

var (
	//go:embed sql/create_new_company.sql
	queryCreateNewCompany string
    //go:embed sql/find_company_by_id.sql
    queryFindById string
)

type CompanyRepository interface {
	Create(*entity.Company) (*entity.Company, error)
    FindById(id string) (*entity.Company, error)
}

type companyRepositoryImpl struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) CompanyRepository {
	return &companyRepositoryImpl{
		DB: db,
	}
}

func (repo *companyRepositoryImpl) Create(company *entity.Company) (*entity.Company, error) {
	args := []interface{}{company.Name, company.Email, company.Phone}

	err := repo.DB.QueryRow(queryCreateNewCompany, args...).Scan(&company.ID, &company.Name, &company.Email, &company.Phone)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (repo *companyRepositoryImpl) FindById(id string) (*entity.Company, error) {
    company := new(entity.Company)

    err := repo.DB.QueryRow(queryFindById, id).Scan(&company.ID, &company.Name, &company.Email, &company.Phone, &company.CreatedAt)
    if err != nil {
        return nil, err
    }

    // TODO - Get company ratings

    return company, nil
}
