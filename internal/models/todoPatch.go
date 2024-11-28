package models

type TodoPatch struct {
	Status string `json:"status,omitempty"`
	Task   string `json:"task,omitempty"`
}
