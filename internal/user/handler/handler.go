package handler

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dinizgab/buildco-api/internal/helpers"
	"github.com/dinizgab/buildco-api/internal/user/entity"
	"github.com/dinizgab/buildco-api/internal/user/repository"
	"github.com/dinizgab/buildco-api/internal/user/usecase"
)

type API struct {
	logger  *slog.Logger
	usecase usecase.UsersUseCase
}

func New(logger *slog.Logger, db *sql.DB) *API {
	repo := repository.NewRepository(db)

	return &API{
		logger:  logger,
		usecase: usecase.NewUsecase(repo),
	}
}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {
	var user *entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}

	createdUser, err := api.usecase.Create(user)
	if err != nil {
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}

	err = json.NewEncoder(w).Encode(createdUser)
	if err != nil {
		api.logger.Error("Something went wrong:", slog.Any("error", err))
		helpers.ServerError(w)

		return
	}
}
