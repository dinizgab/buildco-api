package router

import (
	"log/slog"
	"net/http"

	company "github.com/dinizgab/buildco-api/internal/company/handler"
	requestlog "github.com/dinizgab/buildco-api/internal/router/middleware/request_log"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func New(logger *slog.Logger, db *pgx.Conn) *chi.Mux {
    router := chi.NewRouter()

    router.Route("/v1", func(r chi.Router) {
        
        companyAPI := company.New(logger, db)
        r.Method(http.MethodPost, "/company", requestlog.NewHandler(companyAPI.Create, logger))
    })
    
    return router
}
