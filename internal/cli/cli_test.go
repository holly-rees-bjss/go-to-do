package cli

import (
	"slices"
	"testing"
	"todo_app/internal/models"
	"todo_app/internal/storage"
)

// test the printing of cli app?
// func TestListToDos(t *testing.T) {
// }

func TestCliHandleAddToDo(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
	}}
	app := App{Store: store}

	expected := []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
	}

	app.HandleAdd("add Task 2")
	actual := app.Store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCliHandleMarkComplete(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
	}}
	app := App{Store: store}

	expected := []models.ToDo{
		{Task: "Task 1", Completed: true},
	}

	app.HandleMarkComplete("complete 1")
	actual := app.Store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCliHandleDelete(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
	}}
	app := App{Store: store}

	expected := []models.ToDo{
		{Task: "Task 2", Completed: false},
	}

	app.HandleDelete("delete 1")
	actual := app.Store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCliHandleEdit(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
	}}
	app := App{Store: store}

	expected := []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Edited task", Completed: false},
	}

	app.HandleEdit("edit 2 Edited task")
	actual := app.Store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
