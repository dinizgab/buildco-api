package repository

import (
	"database/sql"
	_ "embed"

	company "github.com/dinizgab/buildco-api/internal/company/entity"
	rating "github.com/dinizgab/buildco-api/internal/ratings/entity"
)

var (
	//go:embed sql/create_new_company.sql
	queryCreateNewCompany string
	//go:embed sql/find_company_by_id.sql
	queryFindById string
)

type CompanyRepository interface {
	Create(*company.Company) (*company.Company, error)
	FindById(id string) (*company.Company, error)
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

func (repo *companyRepositoryImpl) FindById(id string) (*company.Company, error) {
	company := &company.Company{}

	rows, err := repo.DB.Query(queryFindById, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var grade sql.NullInt16
		var comment sql.NullString

		err = rows.Scan(
			&company.ID,
			&company.Name,
			&company.Phone,
			&company.Email,
			&company.CreatedAt,
			&grade,
			&comment,
		)
		if err != nil {
			return nil, err
		}

		if grade.Valid && comment.Valid {
			rating := &rating.Rating{
				Grade:   int(grade.Int16),
				Comment: comment.String,
			}

			company.Ratings = append(company.Ratings, rating)
		}
	}

	return company, nil
}
