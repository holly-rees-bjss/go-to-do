package models

type TodoPatch struct {
	Status string `json:"task,omitempty"`
	Task   string `json:"status,omitempty"`
}
