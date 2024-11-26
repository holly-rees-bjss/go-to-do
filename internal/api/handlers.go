package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"todo_app/internal/models"
)

type Server struct {
	Store models.Store
}

type TodoPatch struct {
	Status string
}

func (s *Server) GetTodos(writer http.ResponseWriter, request *http.Request) {
	logger := request.Context().Value("logger").(*slog.Logger)
	logger.Info("GET request")
	listType := request.URL.Query().Get("list")

	var todos []models.Todo
	switch listType {
	case "archive":
		todos = s.Store.GetArchive()
	case "overdue":
		todos = s.Store.GetOverdue()
	default:
		todos = s.Store.GetTodos()
	}

	json.NewEncoder(writer).Encode(todos)
}

func (s *Server) PostTodo(writer http.ResponseWriter, request *http.Request) {
	logger := request.Context().Value("logger").(*slog.Logger)
	logger.Info("POST todo request")

	body, _ := io.ReadAll(request.Body)
	var todo models.Todo
	_ = json.Unmarshal(body, &todo)

	logger.Debug("Parsed request", "todo", todo)

	s.Store.Add(todo)
	writer.WriteHeader(http.StatusCreated)

	json.NewEncoder(writer).Encode(todo)
}

func (s *Server) PatchTodoStatus(writer http.ResponseWriter, request *http.Request) {
	logger := request.Context().Value("logger").(*slog.Logger)
	logger.Info("PATCH todo request")

	body, _ := io.ReadAll(request.Body)
	var patch TodoPatch
	_ = json.Unmarshal(body, &patch)

	logger.Debug("Parsed request", "patch", patch)

	index := request.URL.Path[len("/api/todo/"):]
	i, _ := strconv.Atoi(index)

	switch {
	case patch.Status == "Completed":
		s.Store.MarkComplete(i)
	case patch.Status == "In Progress":
		s.Store.MarkInProgress(i)
	}
	json.NewEncoder(writer).Encode(s.Store.GetToDo(i))

}

func (s *Server) DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	logger := request.Context().Value("logger").(*slog.Logger)
	logger.Info("DELETE request: " + request.URL.Path)

	index := request.URL.Path[len("/api/todo/"):]
	i, _ := strconv.Atoi(index)
	s.Store.Delete(i)
}
