package loggers

import (
	"log/slog"
	"os"
)

// NewSlogLogger configure and return a new slog logger
func NewSlogLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: getLogLevel(),
	}))
}

// getLogLevel return the slog log level based on the LOG_LEVEL environment variable
func getLogLevel() slog.Level {
	switch os.Getenv("LOG_LEVEL") {
	case "error":
		return slog.LevelError
	case "warn":
		return slog.LevelWarn
	case "info":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	default:
		return slog.LevelWarn
	}
}
