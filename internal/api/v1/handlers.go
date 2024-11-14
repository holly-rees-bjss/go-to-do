package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"todo_app/internal/models"
)

type Server struct {
	Store models.Store
}

type StatusPatch struct {
	Completed bool
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

func (s *Server) PatchTodoStatus(writer http.ResponseWriter, request *http.Request) {
	body, _ := io.ReadAll(request.Body)
	var patch StatusPatch
	_ = json.Unmarshal(body, &patch)
	if patch.Completed {
		index := request.URL.Path[len("/api/todo/"):]
		i, _ := strconv.Atoi(index)
		s.Store.MarkComplete(i)

		json.NewEncoder(writer).Encode(s.Store.GetToDo(i))
	}
}

func (s *Server) DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	index := request.URL.Path[len("/api/todo/"):]
	i, _ := strconv.Atoi(index)
	s.Store.Delete(i)
}
