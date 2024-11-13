package api

import (
	"net/http"
	v1 "todo_app/internal/api/v1"
	"todo_app/internal/models"
)

func setUpRouter(s models.Store) *http.ServeMux {
	router := http.NewServeMux()
	server := &v1.Server{Store: s}
	router.HandleFunc("GET /api/todos", server.GetTodos)
	return router
}
