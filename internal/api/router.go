package api

import (
	"fmt"
	"net/http"
	v1 "todo_app/internal/api/v1"
	"todo_app/internal/models"
)

type App struct {
	Store models.Store
}

func (a App) Run() {
	router := setUpRouter(a.Store)
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func setUpRouter(s models.Store) *http.ServeMux {
	router := http.NewServeMux()
	server := &v1.Server{Store: s}
	router.HandleFunc("GET /api/todos", server.GetTodos)
	router.HandleFunc("POST /api/todo", server.PostTodo)
	router.HandleFunc("PATCH /api/todo/", server.PatchTodoStatus)
	router.HandleFunc("DELETE /api/todo/", server.DeleteTodo)
	return router
}
