package cli

import (
	"bytes"
	"os"
	"slices"
	"testing"
	"todo_app/internal/models"
	"todo_app/internal/storage"
)

func TestCliListToDos(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: true},
	}}
	app := App{Store: store}
	expected := "1. Task 1 [Completed: false]\n2. Task 2 [Completed: true]\n"
	actual := CaptureOutputOf(app.ListToDos)

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

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

func CaptureOutputOf(function func()) string {
	var buffer bytes.Buffer
	originalOutputSetting := os.Stdout
	readEnd, writeEnd, _ := os.Pipe()
	os.Stdout = writeEnd

	function()

	writeEnd.Close()
	os.Stdout = originalOutputSetting

	buffer.ReadFrom(readEnd)
	return buffer.String()
}
