package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dinizgab/buildco-api/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	config := config.New()
	router := chi.NewRouter()

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      router,
		IdleTimeout:  config.Server.TimeoutIdle,
		ReadTimeout:  config.Server.TimeoutRead,
		WriteTimeout: config.Server.TimeoutWrite,
	}

	err := server.ListenAndServe()
	appLogger.Error("error while setting up server", slog.Any("error", err))
}
