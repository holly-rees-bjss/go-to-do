package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"todo_app/internal/models"
)

type Server struct {
	Store models.Store
}

func (s *Server) GetTodos(writer http.ResponseWriter, request *http.Request) {
	todos := s.Store.GetTodos()

	json.NewEncoder(writer).Encode(todos)
}

func (s *Server) PostTodo(writer http.ResponseWriter, request *http.Request) {
	body, _ := io.ReadAll(request.Body)
	var todo models.ToDo
	_ = json.Unmarshal(body, &todo)

	s.Store.Add(todo)
	writer.WriteHeader(http.StatusCreated)

	json.NewEncoder(writer).Encode(todo)
}
