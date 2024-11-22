package main

import (
	"io"
	"log/slog"
	"os"
	"strings"
	"todo_app/internal/api"
	"todo_app/internal/cli"
	l "todo_app/internal/logger"
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

	var logger *slog.Logger

	if len(os.Args) > 2 && strings.Contains(os.Args[1], "log") {
		// Enable logging
		logger = l.InitializeLogger()
	} else {
		// Disable logging by setting a handler that discards logs
		logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	}

	// select app
	appType := os.Args[len(os.Args)-1]

	switch appType {
	case "cli":
		app = cli.App{Store: store, Logger: logger}
	case "api":
		app = api.App{Store: store, Logger: logger}
	}
	app.Run()
}
