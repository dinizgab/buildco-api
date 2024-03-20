package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/dinizgab/buildco-api/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	config := config.New()
	router := chi.NewRouter()

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := server.ListenAndServe()
	app.Logger.Error("error while setting up server", slog.Any("error", err))
}
