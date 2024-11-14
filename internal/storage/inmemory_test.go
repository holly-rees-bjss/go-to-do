package storage

import (
	"slices"
	"testing"
	"todo_app/internal/models"
)

func TestGetToDos(t *testing.T) {
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
		{Task: "Task 3", Completed: false},
	}}

	expected := []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
		{Task: "Task 3", Completed: false},
	}

	actual := store.GetTodos()

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestAddToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
		{Task: "Task 3", Completed: false},
	}}

	newToDo := models.ToDo{Task: "Task 4", Completed: false}

	expected := []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
		{Task: "Task 3", Completed: false},
		{Task: "Task 4", Completed: false},
	}

	store.Add(newToDo)
	actual := store.Todos

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestMarkComplete(t *testing.T) {
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
		{Task: "Task 3", Completed: false},
	}}
	expected := []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
		{Task: "Task 3", Completed: true},
	}

	store.MarkComplete(3)
	actual := store.Todos

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDeleteToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
		{Task: "Task 3", Completed: false},
	}}
	expected := []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
	}

	store.Delete(3)
	actual := store.Todos

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestEditToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "feed the cat", Completed: false},
	}}
	expected := []models.ToDo{
		{Task: "feed the dog", Completed: false},
	}

	store.EditToDo(1, "feed the dog")
	actual := store.Todos

	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestGetToDo(t *testing.T) {
	store := &Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
		{Task: "Task 3", Completed: false},
	}}

	expected := models.ToDo{Task: "Task 2", Completed: false}

	actual := store.GetToDo(2)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
