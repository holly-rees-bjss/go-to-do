package json

import (
	"encoding/json"
	"os"
	"slices"
	"testing"
	"time"
	"todo_app/internal/models"
)

const testFilePath = "./test_todos.json"

func setupJSONStore() (*JsonStore, error) {
	return NewJSONStore(testFilePath)
}

func teardownJSONStore() {
	os.Remove(testFilePath)
}

func TestAddAndGetTodos(t *testing.T) {
	s, err := setupJSONStore()
	if err != nil {
		t.Fatalf("failed to setup JSON store: %v", err)
	}
	defer teardownJSONStore()

	todo := models.Todo{Task: "Task 1", Status: "Not Started"}
	s.Add(todo)
	todos := s.GetTodos()

	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}

	expected := "Task 1"
	actual := todos[0].Task

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestSaveAfterAdd(t *testing.T) {
	s, err := setupJSONStore()
	if err != nil {
		t.Fatalf("failed to setup JSON store: %v", err)
	}
	defer teardownJSONStore()

	todo := models.Todo{Task: "Task 1", Status: "Not Started"}
	s.Add(todo)

	file, err := os.Open(testFilePath)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	var todos []models.Todo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&todos)
	if err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}

	expected := "Task 1"
	actual := todos[0].Task

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func TestLoadFromJsonFile(t *testing.T) {

	s, err := setupJSONStore()
	if err != nil {
		t.Fatalf("failed to setup JSON store: %v", err)
	}
	initialTodos := []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}
	s.Todos = initialTodos
	s.save()

	defer teardownJSONStore()

	t.Run("loads from a file of todos", func(t *testing.T) {

		todos := s.GetTodos()

		if len(todos) != 2 {
			t.Fatalf("expected 2 todos, got %d", len(todos))
		}

		expected := "Task 1"
		actual := todos[0].Task

		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	})
	t.Run("adds to a file of todos", func(t *testing.T) {

		todo := models.Todo{Task: "Task 1", Status: "Not Started"}
		s.Add(todo)
		todos := s.GetTodos()

		if len(todos) != 3 {
			t.Fatalf("expected 2 todos, got %d", len(todos))
		}

		expected := "Task 1"
		actual := todos[0].Task

		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	})
}

func TestMarkComplete(t *testing.T) {
	s, err := setupJSONStore()
	if err != nil {
		t.Fatalf("failed to setup JSON store: %v", err)
	}
	defer teardownJSONStore()

	todo := models.Todo{Task: "Task 1", Status: "Not Started"}
	s.Add(todo)
	s.MarkComplete(1)

	expected := "Completed"

	actual := s.Todos[0].Status

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestMarkNotStarted(t *testing.T) {
	defer teardownJSONStore()
	store := &JsonStore{Todos: []models.Todo{
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
	store := &JsonStore{Todos: []models.Todo{
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
	store := &JsonStore{Todos: []models.Todo{
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
	store := &JsonStore{Todos: []models.Todo{
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
	store := &JsonStore{Todos: []models.Todo{
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
	store := &JsonStore{Todos: []models.Todo{
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
	store := &JsonStore{Todos: []models.Todo{
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
	store := &JsonStore{}

	store.SetOverdueList(overdueList)

	expected := "Task 1"
	actual := store.Overdue[0].Task

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetArchive(t *testing.T) {
	store := &JsonStore{Todos: []models.Todo{
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
	store := &JsonStore{}
	store.SetOverdueList(overdueList)

	expected := 1
	actual := len(store.GetOverdue())
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
