package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"todo_app/internal/api/middleware"
	"todo_app/internal/models"

	"github.com/rs/cors"
)

type App struct {
	Store  models.Store
	Logger *slog.Logger
}

func (a App) Run() {
	a.Logger.Info("Server start up")
	router := a.setUpRouter()

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func (a *App) setUpRouter() http.Handler {
	router := http.NewServeMux() // request router
	server := &Server{Store: a.Store}
	router.HandleFunc("GET /api/todos", server.GetTodos)
	router.HandleFunc("POST /api/todo", server.PostTodo)
	router.HandleFunc("PATCH /api/todo/", server.PatchTodoStatus)
	router.HandleFunc("DELETE /api/todo/", server.DeleteTodo)

	checkOverdueMiddleware := middleware.CheckOverdue(a.Store)
	wrappedRouter := middleware.TraceIDMiddleware(a.Logger, checkOverdueMiddleware(router))
	return wrappedRouter
}
