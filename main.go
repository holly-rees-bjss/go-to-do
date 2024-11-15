package main

import (
	"os"
	"todo_app/internal/api"
	"todo_app/internal/cli"
	"todo_app/internal/models"
	"todo_app/internal/storage"
)

var app interface {
	Run()
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

	switch appType {
	case "cli":
		app = cli.App{Store: store}
	case "api":
		app = api.App{Store: store}
	}

	app.Run()
}
