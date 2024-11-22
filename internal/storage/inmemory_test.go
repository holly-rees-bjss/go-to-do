package storage

import (
	"slices"
	"testing"
	"time"
	"todo_app/internal/models"
)

func TestGetToDos(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}

	expected := []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}

	actual := store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAddToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}

	newToDo := models.Todo{Task: "Task 4", Status: "Not Started"}

	expected := []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
		{Task: "Task 4", Status: "Not Started"},
	}

	store.Add(newToDo)
	actual := store.Todos

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestMarkComplete(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}
	expected := "Completed"

	store.MarkComplete(3)
	actual := store.Todos[2].Status

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestMarkNotStarted(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Completed"},
	}}
	expected := "Not Started"

	store.MarkNotStarted(3)
	actual := store.Todos[2].Status

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestMarkInProgress(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Completed"},
	}}
	expected := "In Progress"

	store.MarkInProgress(3)
	actual := store.Todos[2].Status

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestDeleteToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}
	expected := []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}

	store.Delete(3)
	actual := store.Todos

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestEditToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "feed the cat", Status: "Not Started"},
	}}
	expected := []models.Todo{
		{Task: "feed the dog", Status: "Not Started"},
	}

	store.EditToDo(1, "feed the dog")
	actual := store.Todos

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestGetToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}

	expected := models.Todo{Task: "Task 2", Status: "Not Started"}

	actual := store.GetToDo(2)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestLastUpdatedUpdatesAfterStatusUpdate(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	initialTime := store.Todos[0].LastUpdated

	time.Sleep(1 * time.Second)

	store.MarkInProgress(1)

	updatedTime := store.Todos[0].LastUpdated

	if !updatedTime.After(initialTime) {
		t.Errorf("expected LastUpdated to be after %v, got %v", initialTime, updatedTime)
	}
}

func TestCompleteItemsAutomaticallyMovedToArchiveList(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	store.MarkComplete(2)

	expected := "Task 2"
	actual := store.Archive[0].Task

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSetOverdueList(t *testing.T) {
	pastDueDate := time.Now().Add(-24 * time.Hour)

	overdueList := []models.Todo{
		models.NewToDo("Task 1", pastDueDate),
	}
	store := &Inmemory{}

	store.SetOverdueList(overdueList)

	expected := "Task 1"
	actual := store.Overdue[0].Task

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetArchive(t *testing.T) {
	store := &Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}
	store.MarkComplete(2)
	expected := 1
	actual := len(store.GetArchive())
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetOverdue(t *testing.T) {
	pastDueDate := time.Now().Add(-24 * time.Hour)
	overdueList := []models.Todo{
		models.NewToDo("Task 1", pastDueDate),
	}
	store := &Inmemory{}
	store.SetOverdueList(overdueList)

	expected := 1
	actual := len(store.GetOverdue())
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
