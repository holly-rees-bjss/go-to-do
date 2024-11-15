package models

import "time"

type ToDo struct {
	Task        string
	Status      string
	DueDate     time.Time //optional
	LastUpdated time.Time
}

func NewToDo(task string, dueDate time.Time) ToDo {
	return ToDo{
		Task:        task,
		Status:      "Not Started",
		DueDate:     dueDate,
		LastUpdated: time.Now(),
	}
}
