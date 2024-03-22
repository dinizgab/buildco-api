package handler

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dinizgab/buildco-api/internal/company/entity"
	"github.com/dinizgab/buildco-api/internal/company/repository"
	"github.com/dinizgab/buildco-api/internal/company/usecase"
)

type API struct {
	logger  *slog.Logger
	usecase usecase.CompanyUseCase
}

func New(logger *slog.Logger, db *sql.DB) *API {
    repo := repository.NewRepository(db)

	return &API{
		logger:  logger,
		usecase: usecase.NewUsecase(repo),
	}
}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {
    var company *entity.Company

    err := json.NewDecoder(r.Body).Decode(&company)
    if err != nil {
        api.logger.Error("error: ", slog.Any("error", err))
        return
    }

    createdCompany, err := api.usecase.Create(company)
    if err != nil {
        api.logger.Error("error: ", slog.Any("error", err))
        return
    }

    err = json.NewEncoder(w).Encode(createdCompany)
    if err != nil {
        api.logger.Error("error: ", slog.Any("error", err))
        return
    }
}
