package logger

import (
	"log/slog"
	"os"
)

func New(isDebug bool) *slog.Logger {
	logLevel := slog.LevelInfo
	if isDebug {
		logLevel = slog.LevelDebug
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
}
