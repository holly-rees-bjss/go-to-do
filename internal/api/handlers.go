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
	Status string `json:"Task,omitempty"`
	Task   string `json:"Status,omitempty"`
}

func (s *Server) GetTodos(writer http.ResponseWriter, request *http.Request) {
	slog.InfoContext(request.Context(), "GET Request")

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
	slog.InfoContext(request.Context(), "POST todo request")

	body, _ := io.ReadAll(request.Body)
	var todo models.Todo
	_ = json.Unmarshal(body, &todo)

	slog.DebugContext(request.Context(), "Parsed request", "todo", todo)

	s.Store.Add(todo)
	writer.WriteHeader(http.StatusCreated)

	json.NewEncoder(writer).Encode(todo)
}

func (s *Server) PatchTodo(writer http.ResponseWriter, request *http.Request) {
	slog.InfoContext(request.Context(), "PATCH todo request")

	body, _ := io.ReadAll(request.Body)
	var patch TodoPatch
	_ = json.Unmarshal(body, &patch)

	slog.DebugContext(request.Context(), "Parsed request", "patch", patch)

	index := request.URL.Path[len("/api/todo/"):]
	i, _ := strconv.Atoi(index)

	if patch.Task != "" {
		s.Store.EditToDo(i, patch.Task)
	}

	if patch.Status != "" {
		switch {
		case patch.Status == "Completed":
			s.Store.MarkComplete(i)
		case patch.Status == "In Progress":
			s.Store.MarkInProgress(i)
		}
	}
	json.NewEncoder(writer).Encode(s.Store.GetToDo(i))

}

func (s *Server) DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	slog.InfoContext(request.Context(), "DELETE todo request", "path", request.URL.Path)

	index := request.URL.Path[len("/api/todo/"):]
	i, _ := strconv.Atoi(index)
	s.Store.Delete(i)
}
