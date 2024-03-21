package company

import (
	"github.com/jackc/pgx/v5"
)

type CompanyRepository interface {

}

type companyRepositoryImpl struct {
    DB *pgx.Conn
}

func NewRepository(db *pgx.Conn) CompanyRepository {
    return &companyRepositoryImpl{
        DB: db,
    }
}
