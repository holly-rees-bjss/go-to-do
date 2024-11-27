package middleware

import (
	"net/http"
	"time"
	"todo_app/internal/models"
)

// takes a store and returns a middleware function
func CheckOverdue(store models.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var overdue []models.Todo
			now := time.Now()
			for _, todo := range store.GetTodos() {
				if !todo.DueDate.IsZero() && now.After(todo.DueDate) {
					overdue = append(overdue, todo)
				}
			}
			store.SetOverdueList(overdue)
			next.ServeHTTP(w, r)
		})
	}
}
