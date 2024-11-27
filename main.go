package main

import (
	"flag"
	"log/slog"
	"os"
	"todo_app/internal/api"
	"todo_app/internal/cli"
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

	// select memory
	var storeType = flag.String("store", "memory", "set data store type (memory, json)")
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

	// select app - TO DO update to use flags instead of os.Args
	appType := os.Args[len(os.Args)-1]

	switch appType {
	case "cli":
		logsFile := l.CreateLogsFile()
		defer logsFile.Close()
		logger := l.InitializeLogger(logsFile)
		app = cli.App{Store: store, Logger: logger}
	case "api":
		logger := l.InitializeLogger()
		app = api.App{Store: store, Logger: logger}
	}
	app.Run()
}
