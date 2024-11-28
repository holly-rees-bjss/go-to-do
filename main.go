package main

import (
	"flag"
	"log/slog"
	"os"
	"todo_app/internal/api"
	v2 "todo_app/internal/cli/v2"
	l "todo_app/internal/logger"
	"todo_app/internal/models"
	"todo_app/internal/storage"
	j "todo_app/internal/storage/json"
)

var app interface {
	Run()
}

type App struct {
	Store  models.Store
	Logger *slog.Logger
}

func main() {

	var storeType = flag.String("store", "memory", "set data store type (memory, json)")
	var logLevel = flag.String("loglevel", "error", "set log level (debug, info, warn, error)")

	flag.Parse()

	var store models.Store

	switch *storeType {
	case "inmemory":
		store = &storage.Inmemory{Todos: []models.Todo{
			{Task: "make a list", Status: "Not Started"},
			{Task: "water plants", Status: "Completed"},
			{Task: "go outside", Status: "Not Started"},
			{Task: "touch grass", Status: "Not Started"},
			{Task: "learn go", Status: "In Progress"},
		}}
	case "json":
		store, _ = j.NewJSONStore("./json_store.json")
	}

	appType := os.Args[len(os.Args)-1]

	switch appType {
	case "cli":
		logsFile := l.CreateLogsFile()
		defer logsFile.Close()
		logger := l.InitializeLogger(logLevel, logsFile)
		//app = v1.App{Store: store, Logger: logger}
		app = v2.App{Logger: logger}
	case "api":
		logger := l.InitializeLogger(logLevel)
		app = api.App{Store: store, Logger: logger}
	}
	app.Run()
}
