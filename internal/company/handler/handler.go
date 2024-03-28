package handler

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dinizgab/buildco-api/internal/company/entity"
	"github.com/dinizgab/buildco-api/internal/company/repository"
	"github.com/dinizgab/buildco-api/internal/company/usecase"
	"github.com/dinizgab/buildco-api/internal/helpers"
	"github.com/go-chi/chi/v5"
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
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}

	createdCompany, err := api.usecase.Create(company)
	if err != nil {
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}

	err = json.NewEncoder(w).Encode(createdCompany)
	if err != nil {
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}
}

func (api *API) FindAll(w http.ResponseWriter, r *http.Request) {
	companies, err := api.usecase.FindAll()
	if err != nil {
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}

	err = json.NewEncoder(w).Encode(companies)
	if err != nil {
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}
}

func (api *API) FindById(w http.ResponseWriter, r *http.Request) {
	companyId := chi.URLParam(r, "id")

	result, err := api.usecase.FindById(companyId)
	if err != nil {
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}
}
