package company

import (
	"log/slog"
	"net/http"

	company "github.com/dinizgab/buildco-api/internal/company/usecase"
	"github.com/jackc/pgx/v5"
)

type API struct {
	logger  *slog.Logger
	usecase company.CompanyUseCase
}

func New(logger *slog.Logger, db *pgx.Conn) *API {
	return &API{
		logger:  logger,
		usecase: company.NewUsecase(db),
	}
}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {

}
