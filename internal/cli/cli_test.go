package cli

import (
	"bytes"
	"log/slog"
	"os"
	"slices"
	"testing"
	"time"
	"todo_app/internal/models"
	"todo_app/internal/storage"
)

func setUpAppForTest(store models.Store) App {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	return App{Store: store, Logger: logger}
}

func TestCliHandleListAll(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Completed"},
	}}
	app := setUpAppForTest(store)
	expected := "1. Task 1 [Status: Not Started]\n2. Task 2 [Status: Completed]\n"
	actual := CaptureOutputOf(app.HandleList, "list all")

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCliHandleListAllIfTodoHasDueDate(t *testing.T) {

	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started", DueDate: time.Date(2124, time.November, 22, 0, 0, 0, 0, time.UTC)},
		{Task: "Task 2", Status: "Completed"},
	}}
	app := setUpAppForTest(store)

	expected := "1. Task 1 [Status: Not Started] [Due: 22-11-2124]\n2. Task 2 [Status: Completed]\n"
	actual := CaptureOutputOf(app.HandleList, "list all")

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCliHandleListArchive(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}
	app := setUpAppForTest(store)
	store.MarkComplete(2)
	expected := "1. Task 2 [Status: Completed]\n"
	actual := CaptureOutputOf(app.HandleList, "list archive")

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCliHandleAddTodo(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
	}}
	app := setUpAppForTest(store)

	expected := 2

	app.HandleAdd("add Task 2")
	actual := len(app.Store.GetTodos())

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	addedTask := store.Todos[1].Task
	expectedTask := "Task 2"

	if addedTask != "Task 2" {
		t.Errorf("Expected %v, got %v", expectedTask, addedTask)
	}

}

func TestHandleAddTodoWithDueDate(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
	}}
	app := setUpAppForTest(store)

	dueDate := time.Now().Add(24 * time.Hour)
	formattedDueDate := dueDate.Format("02-01-2006")

	expected := formattedDueDate

	app.HandleAdd("add Task 2 due " + formattedDueDate)

	actual := store.Todos[1].DueDate.Format("02-01-2006")

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}

}

func TestCliHandleMarkComplete(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
	}}
	app := setUpAppForTest(store)

	expected := "Completed"

	app.HandleMarkComplete("complete 1")
	actual := app.Store.GetTodos()[0].Status

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCliHandleInProgress(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
	}}
	app := setUpAppForTest(store)

	expected := "In Progress"

	app.HandleMarkInProgress("in progress 1")
	actual := app.Store.GetTodos()[0].Status

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCliHandleDelete(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}
	app := setUpAppForTest(store)

	expected := []models.Todo{
		{Task: "Task 2", Status: "Not Started"},
	}

	app.HandleDelete("delete 1")
	actual := app.Store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCliHandleEdit(t *testing.T) {
	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}
	app := setUpAppForTest(store)

	expected := []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Edited task", Status: "Not Started"},
	}

	app.HandleEdit("edit 2 Edited task")
	actual := app.Store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func CaptureOutputOf(function func(...string), parameter string) string {
	var buffer bytes.Buffer
	originalOutputSetting := os.Stdout
	readEnd, writeEnd, _ := os.Pipe()
	os.Stdout = writeEnd

	function(parameter)

	writeEnd.Close()
	os.Stdout = originalOutputSetting

	buffer.ReadFrom(readEnd)
	return buffer.String()
}
