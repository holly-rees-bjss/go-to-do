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
		{Task: "Task 1", Completed: false},
		{Task: "Task 2", Completed: false},
	}}

	serv := &Server{store}
	handler := http.HandlerFunc(serv.GetTodos)

	request, _ := http.NewRequest(http.MethodGet, "/api/todos", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	t.Run("GET /api/todos returns a JSON list of Todos", func(t *testing.T) {

		expected := `[{"Task":"Task 1","Completed":false},{"Task":"Task 2","Completed":false}]`
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
	todo := models.ToDo{Task: "Task 1"}
	body, _ := json.Marshal(todo)

	request, _ := http.NewRequest(http.MethodPost, "/api/todo", strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	server := &Server{store}
	handler := http.HandlerFunc(server.PostTodo)
	handler.ServeHTTP(response, request)

	t.Run("POST /api/todo returns newly posted Todo", func(t *testing.T) {

		expected := `{"Task":"Task 1","Completed":false}`
		actual := strings.TrimSpace(response.Body.String())

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("POST /api/todo adds todo to storage", func(t *testing.T) {

		expected := []models.ToDo{
			{Task: "Task 1", Completed: false},
		}
		actual := store.Todos

		if !slices.Equal(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("POST /api/todo returns a 200 status code when successful", func(t *testing.T) {

		expected := http.StatusCreated
		actual := response.Code

		if actual != expected {
			t.Errorf("returned wrong status code: got %v expected %v", actual, expected)
		}
	})
}
