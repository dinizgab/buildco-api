package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dinizgab/buildco-api/config"
	"github.com/dinizgab/buildco-api/internal/logger"
	"github.com/dinizgab/buildco-api/internal/router"
)

func main() {
	config := config.New()
    logger := logger.New(config.Server.Debug)
    router := router.New(logger)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      router,
		IdleTimeout:  config.Server.TimeoutIdle,
		ReadTimeout:  config.Server.TimeoutRead,
		WriteTimeout: config.Server.TimeoutWrite,
	}
    
    logger.Info(fmt.Sprintf("Running server in port %d", config.Server.Port))
	err := server.ListenAndServe()
	logger.Error("error while setting up server", slog.Any("error", err))
}
