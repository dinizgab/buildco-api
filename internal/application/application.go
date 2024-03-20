package application

import (
	"log/slog"
	"os"
)

type Application struct {
	Port    int
	Env     string
	Version string
	Logger  *slog.Logger
}

func New() *Application {
	return &Application{
		Port:    8080,
		Env:     "development",
		Version: "1.0.0",
		Logger:  slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}
