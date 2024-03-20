package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/dinizgab/buildco-api/internal/application"
	"github.com/go-chi/chi/v5"
)

func main() {
	app := application.New()
	router := chi.NewRouter()

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", app.Port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := server.ListenAndServe()
	app.Logger.Error("error while setting up server", slog.Any("error", err))
}
