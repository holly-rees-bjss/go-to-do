package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func InitializeLogger(logLevel *string, optionalLogFile ...*os.File) *slog.Logger {

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

	var loggerOutput *os.File

	if len(optionalLogFile) > 0 && optionalLogFile[0] != nil {
		loggerOutput = optionalLogFile[0]
	} else {
		loggerOutput = os.Stderr
	}

	logger := slog.New(slog.NewJSONHandler(loggerOutput, &slog.HandlerOptions{
		Level: level,
	}))

	return logger
}

func CreateLogsFile() *os.File {
	logsFile, err := os.Create("cli_logs.log")
	if err != nil {
		fmt.Println("Error creating logs file")
		panic(err)
	}
	return logsFile
}
