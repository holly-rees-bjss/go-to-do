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

// this tests the conversion from string to ToDo
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

// ask alan/oliver about this test - just testing what's already been tested in inmemory test?
// func TestCliMarkComplete(t *testing.T) {
// 	store := &storage.Inmemory{Todos: []models.ToDo{
// 		{Task: "Task 1", Completed: false},
// 	}}
// 	app := App{Store: store}

// 	expected := []models.ToDo{
// 		{Task: "Task 1", Completed: true},
// 	}

// 	app.MarkComplete(1)
// 	actual := app.Store.GetTodos()

// 	if !slices.Equal(actual, expected) {
// 		t.Errorf("Expected %v, got %v", expected, actual)
// 	}
// }

// tests for CLI interface? mock?
