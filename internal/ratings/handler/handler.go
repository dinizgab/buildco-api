package handler

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dinizgab/buildco-api/internal/helpers"
	"github.com/dinizgab/buildco-api/internal/ratings/entity"
	"github.com/dinizgab/buildco-api/internal/ratings/repository"
	"github.com/dinizgab/buildco-api/internal/ratings/usecase"
	"github.com/go-chi/chi/v5"
)

type API struct {
	logger  *slog.Logger
	usecase usecase.RatingsUsecase
}

func New(logger *slog.Logger, db *sql.DB) *API {
	repo := repository.NewRepository(db)

	return &API{
		logger:  logger,
		usecase: usecase.NewUsecase(repo),
	}
}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {
    companyId := chi.URLParam(r, "id")
	var rating *entity.Rating

	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		api.logger.Error("Something went wrong!", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}

    newRating, err := api.usecase.Create(companyId, rating)
	if err != nil {
		api.logger.Error("Something went wrong!", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}

	err = json.NewEncoder(w).Encode(newRating)
	if err != nil {
		api.logger.Error("Something went wrong!", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}
}
