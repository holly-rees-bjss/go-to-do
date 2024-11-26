package json

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"todo_app/internal/models"
)

type JsonStore struct {
	FilePath string
	Todos    []models.Todo
	Archive  []models.Todo
	Overdue  []models.Todo
}

func NewJSONStore(filePath string) (*JsonStore, error) {
	store := &JsonStore{FilePath: filePath}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return nil, fmt.Errorf("failed to initialize file: %v", err)
		}
	}

	err := store.load()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *JsonStore) load() error {
	file, err := os.Open(s.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s.Todos)
	if err != nil {
		return err
	}
	return nil
}

func (s *JsonStore) save() error {
	file, err := os.Create(s.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(s.Todos)
	if err != nil {
		return err
	}
	return nil
}

func (m *JsonStore) GetTodos() []models.Todo {
	return m.Todos
}

func (m *JsonStore) GetArchive() []models.Todo {
	return m.Archive
}

func (m *JsonStore) GetOverdue() []models.Todo {
	return m.Overdue
}

func (m *JsonStore) Add(newToDo models.Todo) {
	m.Todos = append(m.Todos, newToDo)
	m.save()
}

func (m *JsonStore) updateStatus(num int, status string) error {
	i := num - 1
	if i < 0 || i >= len(m.Todos) {
		return fmt.Errorf("invalid task number")
	}

	m.Todos[i].Status = status
	m.Todos[i].LastUpdated = time.Now()
	return m.save()
}

func (m *JsonStore) MarkComplete(num int) (err error) {
	m.updateStatus(num, "Completed")
	m.Archive = append(m.Archive, m.Todos[num-1])
	return m.save()
}

func (m *JsonStore) MarkNotStarted(num int) (err error) {
	m.updateStatus(num, "Not Started")
	return m.save()
}

func (m *JsonStore) MarkInProgress(num int) (err error) {
	m.updateStatus(num, "In Progress")
	return m.save()
}

func (m *JsonStore) Delete(num int) (err error) {
	i := num - 1
	if i < 0 || i >= len(m.Todos) {
		return fmt.Errorf("invalid number task")
	}
	m.Todos = append(m.Todos[:i], m.Todos[i+1:]...)
	return m.save()
}

func (m *JsonStore) EditToDo(num int, edit string) error {
	i := num - 1
	if i < 0 || i >= len(m.Todos) {
		return fmt.Errorf("invalid number task")
	}

	m.Todos[i].Task = edit
	return m.save()
}

func (m *JsonStore) GetToDo(num int) models.Todo {
	i := num - 1
	return m.Todos[i]
}

func (m *JsonStore) SetOverdueList(overdue []models.Todo) {
	m.Overdue = overdue
	m.save()
}
