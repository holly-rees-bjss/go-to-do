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
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Completed"},
	}}
	app := App{Store: store}
	expected := "1. Task 1 [Status: Not Started]\n2. Task 2 [Status: Completed]\n"
	actual := CaptureOutputOf(app.ListToDos)

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCliHandleAddToDo(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
	}}
	app := App{Store: store}

	expected := []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}

	app.HandleAdd("add Task 2")
	actual := app.Store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCliHandleMarkComplete(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
	}}
	app := App{Store: store}

	expected := []models.ToDo{
		{Task: "Task 1", Status: "Completed"},
	}

	app.HandleMarkComplete("complete 1")
	actual := app.Store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCliHandleDelete(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}
	app := App{Store: store}

	expected := []models.ToDo{
		{Task: "Task 2", Status: "Not Started"},
	}

	app.HandleDelete("delete 1")
	actual := app.Store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCliHandleEdit(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}
	app := App{Store: store}

	expected := []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Edited task", Status: "Not Started"},
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
