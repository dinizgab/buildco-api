package company

import (
	"database/sql"
	"log/slog"
	"net/http"

	company "github.com/dinizgab/buildco-api/internal/company/usecase"
)

type API struct {
	logger  *slog.Logger
	usecase company.CompanyUseCase
}

func New(logger *slog.Logger, db *sql.DB) *API {
	return &API{
		logger:  logger,
		usecase: company.NewUsecase(db),
	}
}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {

}
