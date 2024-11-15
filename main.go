package main

import (
	"todo_app/internal/api"
	"todo_app/internal/models"
	"todo_app/internal/storage"
)

func main() {
	// select memory
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "make a list", Status: "Not Started"},
		{Task: "water plants", Status: "Completed"},
		{Task: "go outside", Status: "Not Started"},
		{Task: "touch grass", Status: "Not Started"},
		{Task: "learn go", Status: "In Progress"},
	}}

	// select app
	// app := cli.App{Store: store}
	app := api.App{Store: store}
	app.Run()
}
