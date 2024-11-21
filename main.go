package main

import (
	"flag"
	"log/slog"
	"os"
	"todo_app/internal/api"
	"todo_app/internal/cli"
	"todo_app/internal/models"
	"todo_app/internal/storage"
)

var app interface {
	Run()
}

type App struct {
	Store  models.Store
	Logger *slog.Logger
}

func main() {
	// select memory
	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "make a list", Status: "Not Started"},
		{Task: "water plants", Status: "Completed"},
		{Task: "go outside", Status: "Not Started"},
		{Task: "touch grass", Status: "Not Started"},
		{Task: "learn go", Status: "In Progress"},
	}}

	// select app
	appType := os.Args[1]

	logger := initializeLogger()

	switch appType {
	case "cli":
		app = cli.App{Store: store, Logger: logger}
	case "api":
		app = api.App{Store: store, Logger: logger}
	}

	app.Run()
}

func initializeLogger() *slog.Logger {
	var logLevel = flag.String("loglevel", "info", "set log level (debug, info, warn, error)")

	flag.Parse()

	// Set up the logger
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
		level = slog.LevelInfo
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	return logger
}
