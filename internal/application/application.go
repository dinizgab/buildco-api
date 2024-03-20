package application

import (
	"log"
	"log/slog"
	"os"
)

type Application struct {
	Port    int
	Env     string
	Version string
	Logger  *log.Logger
}

func New() *Application {
	return &Application{
		Port:    8080,
		Env:     "development",
		Version: "1.0.0",
		Logger:  slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}
