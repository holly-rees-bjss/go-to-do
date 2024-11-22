package models

type Store interface {
	GetTodos() []Todo
	GetArchive() []Todo
	Add(Todo)
	MarkComplete(int) error
	MarkInProgress(int) error
	MarkNotStarted(int) error
	Delete(int) error
	EditToDo(int, string) error
	GetToDo(int) Todo
	SetOverdueList([]Todo)
	GetOverdue() []Todo
}
