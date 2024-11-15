package storage

import (
	"fmt"
	"time"
	"todo_app/internal/models"
)

type Inmemory struct {
	Todos   []models.Todo
	Archive []models.Todo
	Overdue []models.Todo
}

func (m *Inmemory) GetTodos() []models.Todo {
	m.CheckAnyOverdue()
	return m.Todos
}

func (m *Inmemory) GetArchive() []models.Todo {
	m.CheckAnyOverdue()
	return m.Archive
}

func (m *Inmemory) Add(newToDo models.Todo) {
	m.CheckAnyOverdue()
	m.Todos = append(m.Todos, newToDo)
}

func (m *Inmemory) updateStatus(num int, status string) error {
	m.CheckAnyOverdue()
	i := num - 1
	if i < 0 || i >= len(m.Todos) {
		return fmt.Errorf("invalid task number")
	}

	m.Todos[i].Status = status
	m.Todos[i].LastUpdated = time.Now()
	return nil
}

func (m *Inmemory) MarkComplete(num int) (err error) {
	m.CheckAnyOverdue()
	m.updateStatus(num, "Completed")
	m.Archive = append(m.Archive, m.Todos[num-1])
	return nil
}

func (m *Inmemory) MarkNotStarted(num int) (err error) {
	m.CheckAnyOverdue()
	return m.updateStatus(num, "Not Started")
}

func (m *Inmemory) MarkInProgress(num int) (err error) {
	m.CheckAnyOverdue()
	return m.updateStatus(num, "In Progress")
}

func (m *Inmemory) Delete(num int) (err error) {
	i := num - 1
	if i < 0 || i >= len(m.Todos) {
		return fmt.Errorf("invalid number task")
	}
	m.Todos = append(m.Todos[:i], m.Todos[i+1:]...)
	m.CheckAnyOverdue()
	return nil
}

func (m *Inmemory) EditToDo(num int, edit string) error {
	i := num - 1
	if i < 0 || i >= len(m.Todos) {
		return fmt.Errorf("invalid number task")
	}

	m.Todos[i].Task = edit
	m.CheckAnyOverdue()
	return nil
}

func (m *Inmemory) GetToDo(num int) models.Todo {
	i := num - 1
	m.CheckAnyOverdue()
	return m.Todos[i]
}

func (m *Inmemory) CheckAnyOverdue() {
	var overdue []models.Todo
	now := time.Now()
	for _, todo := range m.Todos {
		if now.After(todo.DueDate) {
			overdue = append(overdue, todo)
		}
	}
	m.Overdue = overdue
}
