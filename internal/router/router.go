package router

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	company "github.com/dinizgab/buildco-api/internal/company/handler"
	rating "github.com/dinizgab/buildco-api/internal/ratings/handler"
	requestlog "github.com/dinizgab/buildco-api/internal/router/middleware/request_log"
	"github.com/go-chi/chi/v5"
)

func New(logger *slog.Logger, db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		ratingAPI := rating.New(logger, db)
		companyAPI := company.New(logger, db)

		r.Method(http.MethodPost, "/company", requestlog.NewHandler(companyAPI.Create, logger))
		r.Method(http.MethodGet, "/company", requestlog.NewHandler(companyAPI.FindAll, logger))
		r.Method(http.MethodGet, "/company/{id}", requestlog.NewHandler(companyAPI.FindById, logger))

		r.Method(http.MethodPost, "/rating/{id}", requestlog.NewHandler(ratingAPI.Create, logger))
	})

	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	return router
}
