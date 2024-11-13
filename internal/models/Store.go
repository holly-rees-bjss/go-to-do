package models

type Store interface {
	GetTodos() []ToDo
	Add(ToDo)
	MarkComplete(int) error
	Delete(int) error
	EditToDo(int, string) error
}
