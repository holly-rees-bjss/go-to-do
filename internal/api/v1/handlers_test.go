package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
	"todo_app/internal/models"
	"todo_app/internal/storage"
)

func TestGetHandler(t *testing.T) {

	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	serv := &Server{store}
	handler := http.HandlerFunc(serv.GetTodos)

	request, _ := http.NewRequest(http.MethodGet, "/api/todos", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	t.Run("GET /api/todos returns a JSON list of Todos", func(t *testing.T) {

		expected := `[{"Task":"Task 1","Status":"Not Started"},{"Task":"Task 2","Status":"Not Started"}]`
		actual := strings.TrimSpace(response.Body.String())

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("GET /api/todos returns a 200 status code when successful", func(t *testing.T) {

		expected := http.StatusOK
		actual := response.Code

		if actual != expected {
			t.Errorf("returned wrong status code: got %v expected %v", actual, expected)
		}
	})

}

func TestPostTodoHandler(t *testing.T) {
	store := &storage.Inmemory{}
	todo := models.ToDo{Task: "Task 1", Status: "Not Started"}
	body, _ := json.Marshal(todo)

	request, _ := http.NewRequest(http.MethodPost, "/api/todo", strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	server := &Server{store}
	handler := http.HandlerFunc(server.PostTodo)
	handler.ServeHTTP(response, request)

	t.Run("POST /api/todo returns newly posted Todo", func(t *testing.T) {

		expected := `{"Task":"Task 1","Status":"Not Started"}`
		actual := strings.TrimSpace(response.Body.String())

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("POST /api/todo adds todo to storage", func(t *testing.T) {

		expected := []models.ToDo{
			{Task: "Task 1", Status: "Not Started"},
		}
		actual := store.Todos

		if !slices.Equal(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("POST /api/todo returns a 201 status code when successful", func(t *testing.T) {

		expected := http.StatusCreated
		actual := response.Code

		if actual != expected {
			t.Errorf("returned wrong status code: got %v expected %v", actual, expected)
		}
	})
}

func TestPatchTodoStatusCompletedHandler(t *testing.T) {

	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	patch := StatusPatch{Completed: true}
	body, _ := json.Marshal(patch)
	request, _ := http.NewRequest(http.MethodPatch, "/api/todo/1", strings.NewReader(string(body)))
	response := httptest.NewRecorder()

	serv := &Server{store}
	handler := http.HandlerFunc(serv.PatchTodoStatus)

	handler.ServeHTTP(response, request)

	t.Run("PATCH /api/todo/{i} returns the patched todo", func(t *testing.T) {

		expected := `{"Task":"Task 1","Status":"Completed"}`
		actual := strings.TrimSpace(response.Body.String())

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("PATCH /api/todo/{i} returns a 200 status code when successful", func(t *testing.T) {

		expected := http.StatusOK
		actual := response.Code

		if actual != expected {
			t.Errorf("returned wrong status code: got %v expected %v", actual, expected)
		}
	})
}

func TestDeleteTodoHandler(t *testing.T) {

	store := &storage.Inmemory{Todos: []models.ToDo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	request, _ := http.NewRequest(http.MethodDelete, "/api/todo/1", nil)
	response := httptest.NewRecorder()

	serv := &Server{store}
	handler := http.HandlerFunc(serv.DeleteTodo)

	handler.ServeHTTP(response, request)

	t.Run("DELETE /api/todo/{i} deletes todo", func(t *testing.T) {

		expectedLen := 1
		actualLen := len(store.Todos)
		expected := models.ToDo{Task: "Task 2", Status: "Not Started"}
		actual := store.Todos[0]

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
		if actualLen != expectedLen {
			t.Errorf("Expected length %v, got %v", expected, actual)
		}
	})

	t.Run("DELETE /api/todo/{i} returns a 200 status code when successful", func(t *testing.T) {

		expected := http.StatusOK
		actual := response.Code

		if actual != expected {
			t.Errorf("returned wrong status code: got %v expected %v", actual, expected)
		}
	})
}
