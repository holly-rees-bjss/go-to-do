package logger

import (
	"flag"
	"log/slog"
	"os"
)

func InitializeLogger() *slog.Logger {
	var logLevel = flag.String("loglevel", "error", "set log level (debug, info, warn, error)")

	flag.Parse()

	var level slog.Level
	switch *logLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelError
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	return logger
}
