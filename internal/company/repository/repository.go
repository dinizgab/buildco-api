package company

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

type CompanyRepository interface {

}

type companyRepositoryImpl struct {
    DB *sql.DB
}

func NewRepository(db *sql.DB) CompanyRepository {
    return &companyRepositoryImpl{
        DB: db,
    }
}
