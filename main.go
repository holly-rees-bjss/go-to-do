package main

import (
	"todo_app/internal/api"
	"todo_app/internal/models"
	"todo_app/internal/storage"
)

func main() {
	// select memory
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "make a list", Completed: true},
		{Task: "water plants", Completed: false},
		{Task: "go outside", Completed: false},
		{Task: "touch grass", Completed: true},
		{Task: "learn go", Completed: false},
	}}

	// select app
	// app := cli.App{Store: store}
	app := api.App{Store: store}
	app.Run()
}
