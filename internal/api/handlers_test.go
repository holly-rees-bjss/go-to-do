package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
	"time"
	"todo_app/internal/api/middleware"
	"todo_app/internal/models"
	"todo_app/internal/storage"

	"github.com/google/go-cmp/cmp"
)

func TestGetHandler(t *testing.T) {

	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	serv := &Server{Store: store}
	handler := http.HandlerFunc(serv.GetTodos)

	t.Run("GET /api/todos returns a JSON list of Todos", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/api/todos", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		expected := []models.Todo{
			{Task: "Task 1", Status: "Not Started"},
			{Task: "Task 2", Status: "Not Started"},
		}

		var actual []models.Todo
		if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
			t.Fatal(err)
		}

		for i, todo := range actual {
			// ignores DueDate and LastUpdated time.Time fields - could add mock to test for time
			if todo.Task != expected[i].Task || todo.Status != expected[i].Status {
				t.Errorf("expected %v, got %v", expected[i], todo)
			}
		}
	})

	t.Run("GET /api/todos returns a 200 status code when successful", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/api/todos", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		expected := http.StatusOK
		actual := response.Code

		if actual != expected {
			t.Errorf("returned wrong status code: got %v expected %v", actual, expected)
		}
	})

}

func TestPostTodoHandler(t *testing.T) {
	store := &storage.Inmemory{}
	todo := models.Todo{Task: "Task 1", Status: "Not Started"}
	body, _ := json.Marshal(todo)

	request, _ := http.NewRequest(http.MethodPost, "/api/todo", strings.NewReader(string(body)))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	server := &Server{Store: store}
	handler := http.HandlerFunc(server.PostTodo)
	handler.ServeHTTP(response, request)

	t.Run("POST /api/todo returns newly posted Todo", func(t *testing.T) {

		expected := models.Todo{Task: "Task 1", Status: "Not Started"}

		var actual models.Todo
		if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
			t.Fatal(err)
		}

		// ignores DueDate and LastUpdated time.Time fields - could add mock to test for time
		if actual.Task != expected.Task || actual.Status != expected.Status {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("POST /api/todo adds todo to storage", func(t *testing.T) {

		expected := []models.Todo{
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

func TestPatchTodoCompletedHandler(t *testing.T) {

	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	patch := models.TodoPatch{Status: "Completed"}
	body, _ := json.Marshal(patch)
	request, _ := http.NewRequest(http.MethodPatch, "/api/todo/1", strings.NewReader(string(body)))
	response := httptest.NewRecorder()

	serv := &Server{Store: store}
	handler := http.HandlerFunc(serv.PatchTodo)

	handler.ServeHTTP(response, request)

	t.Run("PATCH /api/todo/{i} returns the patched todo", func(t *testing.T) {

		expected := models.Todo{Task: "Task 1", Status: "Completed"}

		var actual models.Todo
		if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
			t.Fatal(err)
		}

		// ignores DueDate and LastUpdated time.Time fields - could add mock to test for time
		if actual.Task != expected.Task || actual.Status != expected.Status {
			t.Errorf("expected %v, got %v", expected, actual)
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

func TestPatchCompletedAddsTodoToArchive(t *testing.T) {

	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	patch := models.TodoPatch{Status: "Completed"}
	body, _ := json.Marshal(patch)
	request, _ := http.NewRequest(http.MethodPatch, "/api/todo/2", strings.NewReader(string(body)))
	response := httptest.NewRecorder()

	serv := &Server{Store: store}
	handler := http.HandlerFunc(serv.PatchTodo)

	handler.ServeHTTP(response, request)

	t.Run("GET /api/todos?list=archive returns a JSON list of archived Todos", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/api/todos?list=archive", nil)
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(serv.GetTodos)

		handler.ServeHTTP(response, request)
		expected := []models.Todo{
			{Task: "Task 2", Status: "Completed"},
		}

		var actual []models.Todo
		if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
			t.Fatal(err)
		}

		if len(actual) != len(expected) {
			t.Fatalf("expected %d todos, got %d", len(expected), len(actual))
		}

		for i, todo := range actual {
			// ignores DueDate and LastUpdated time.Time fields - could add mock to test for time
			if todo.Task != expected[i].Task || todo.Status != expected[i].Status {
				t.Errorf("expected %v, got %v", expected[i], todo)
			}
		}
	})
}

func TestPatchTodoInProgressHandler(t *testing.T) {

	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	patch := models.TodoPatch{Status: "In Progress"}
	body, _ := json.Marshal(patch)
	request, _ := http.NewRequest(http.MethodPatch, "/api/todo/1", strings.NewReader(string(body)))
	response := httptest.NewRecorder()

	serv := &Server{Store: store}
	handler := http.HandlerFunc(serv.PatchTodo)

	handler.ServeHTTP(response, request)

	t.Run("PATCH /api/todo/{i} returns the patched todo", func(t *testing.T) {

		expected := models.Todo{Task: "Task 1", Status: "In Progress"}

		var actual models.Todo
		if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
			t.Fatal(err)
		}

		// ignores DueDate and LastUpdated time.Time fields - could add mock to test for time
		if actual.Task != expected.Task || actual.Status != expected.Status {
			t.Errorf("expected %v, got %v", expected, actual)
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

	store := &storage.Inmemory{Todos: []models.Todo{
		{Task: "Task 1", Status: "Not Started"},
		{Task: "Task 2", Status: "Not Started"},
	}}

	request, _ := http.NewRequest(http.MethodDelete, "/api/todo/1", nil)
	response := httptest.NewRecorder()

	serv := &Server{Store: store}
	handler := http.HandlerFunc(serv.DeleteTodo)

	handler.ServeHTTP(response, request)

	t.Run("DELETE /api/todo/{i} deletes todo", func(t *testing.T) {

		expectedLen := 1
		actualLen := len(store.Todos)
		expected := models.Todo{Task: "Task 2", Status: "Not Started"}
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

func TestGetOverdueTodos(t *testing.T) {
	pastDueDate := time.Date(2024, time.November, 21, 0, 0, 0, 0, time.UTC)
	futureDueDate := time.Now().Add(24 * time.Hour)

	store := &storage.Inmemory{Todos: []models.Todo{
		models.NewToDo("Task 1", pastDueDate),
		models.NewToDo("Task 2", futureDueDate),
	}}

	server := &Server{Store: store}
	checkOverdueMiddleware := middleware.CheckOverdue(store)
	handler := checkOverdueMiddleware(http.HandlerFunc(server.GetTodos))

	t.Run("GET /api/todos?list=overdue returns a JSON list of overdue Todos", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/todos?list=overdue", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		expected := []models.Todo{
			{Task: "Task 1", Status: "Not Started", DueDate: pastDueDate},
		}

		var actual []models.Todo
		if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
			t.Fatal(err)
		}

		if len(actual) != len(expected) {
			t.Fatalf("expected %d todos, got %d", len(expected), len(actual))
		}

		for i, todo := range actual {
			if todo.Task != expected[i].Task || todo.Status != expected[i].Status || !todo.DueDate.Equal(expected[i].DueDate) {
				t.Errorf("expected %v, got %v", expected[i], todo)
			}
		}
	})

	t.Run("GET /api/todos?status=overdue returns a 200 status code when successful", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/todos?status=overdue", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		expected := http.StatusOK
		actual := response.Code

		if actual != expected {
			t.Errorf("returned wrong status code: got %v expected %v", actual, expected)
		}
	})
}

func FuzzPostTodo(f *testing.F) {
	// seed corpus - sets of sample data
	testcases := [][]byte{
		[]byte(`{"Task":"Test Todo 1","Status":"Not Started"}`),
		[]byte(`{"Task":"","Status":""}`),
		[]byte(`{}`),
		[]byte(`{"Task":"Test Todo 2","Status":"In Progress"}`),
		[]byte(`{"Task":""}`),
		[]byte(`{"Task":"","Status":"Completed"}`),
		[]byte(`{"todo":""}`),
		[]byte(`{"Task":"Invalid JSON"`),
		[]byte(`{"Task":123, "Status":true}`),
		[]byte(`[{"Task":"Array Task", "Status":"Not Started"}]`),
		[]byte(`{"Task":"Special \n Characters", "Status":"In Progress"}`),
		[]byte(`{"Task":null, "Status":null}`),
		[]byte(`{"invalidJSON asgdfkdgasjfldhalfdsgJAFDHSFL;DGSFJDHGSLFA;DSFDJSGFHDSpecialSFJKDGHSDGKS", "Status":"In Progress"}`),
	}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, in []byte) {
		req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(in))
		req.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		store := &storage.Inmemory{Todos: []models.Todo{}}
		server := &Server{Store: store}
		server.PostTodo(response, req)

		if status := response.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		var todo models.Todo
		err := json.Unmarshal(response.Body.Bytes(), &todo)
		if err != nil {
			t.Errorf("could not unmarshal response body: %v", err)
		}

		i := len(store.Todos)
		storedTodo := store.GetToDo(i)
		if err != nil {
			t.Errorf("could not get stored todo: %v", err)
		}

		if diff := cmp.Diff(todo, storedTodo); diff != "" {
			t.Errorf("stored todo mismatch (-want +got):\n%s", diff)
		}

		fmt.Println("Stored todo: ", storedTodo)
	})
}
