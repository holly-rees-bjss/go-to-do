package models

import "time"

type Todo struct {
	Task        string
	Status      string
	DueDate     time.Time //optional
	LastUpdated time.Time
}

func NewToDo(task string, dueDate time.Time) Todo {
	return Todo{
		Task:        task,
		Status:      "Not Started",
		DueDate:     dueDate,
		LastUpdated: time.Now(),
	}
}
