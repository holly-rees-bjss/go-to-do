package storage

import (
	"slices"
	"testing"
	"todo_app/internal/models"
)

func TestGetToDos(t *testing.T) {
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}

	expected := []models.ToDo{
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
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}

	newToDo := models.ToDo{Task: "Task 4", Status: "Not Started"}

	expected := []models.ToDo{
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
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}
	expected := []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Completed"},
	}

	store.MarkComplete(3)
	actual := store.Todos

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDeleteToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}
	expected := []models.ToDo{
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
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "feed the cat", Status: "Not Started"},
	}}
	expected := []models.ToDo{
		{Task: "feed the dog", Status: "Not Started"},
	}

	store.EditToDo(1, "feed the dog")
	actual := store.Todos

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestGetToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
		{Task: "Task 3", Status: "Not Started"},
	}}

	expected := models.ToDo{Task: "Task 2", Status: "Not Started"}

	actual := store.GetToDo(2)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
