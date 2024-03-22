package handler 

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/dinizgab/buildco-api/internal/company/usecase"
)

type API struct {
	logger  *slog.Logger
	usecase usecase.CompanyUseCase
}

func New(logger *slog.Logger, db *sql.DB) *API {
	return &API{
		logger:  logger,
		usecase: usecase.NewUsecase(db),
	}
}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {

}
