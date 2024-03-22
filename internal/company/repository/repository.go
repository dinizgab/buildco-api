package company

import (
	"database/sql"
    _ "embed"

	company "github.com/dinizgab/buildco-api/internal/company/entity"
)

var (
    //go:embed sql/create_new_company.sql
    queryCreateNewCompany string
)

type CompanyRepository interface {
    Create(*company.Company) (*company.Company, error)
}

type companyRepositoryImpl struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) CompanyRepository {
	return &companyRepositoryImpl{
		DB: db,
	}
}

func (repo *companyRepositoryImpl) Create(company *company.Company) (*company.Company, error) {
    args := []interface{}{company.Name, company.Email, company.Phone}
     
    err := repo.DB.QueryRow(queryCreateNewCompany, args...).Scan(&company.ID, &company.Name, &company.Email, &company.Phone)
    if err != nil {
        return nil, err
    }

    return company, nil
}
