package api_calls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"todo_app/internal/models"
)

const baseUrl = "http://localhost:8080/api/"

func Add(todo models.Todo) {
	url := baseUrl + "todo"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		slog.Error("Error marshalling JSON", "error", err)
		fmt.Println("Couldn't Add your todo, incorrect format")
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Error("Error creating request", "error", err)
		fmt.Println("Couldn't Add your todo")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error creating request", "error", err)
		fmt.Println("Couldn't Add your todo")
	}
	defer resp.Body.Close()
}

func GetAll() []models.Todo {
	url := baseUrl + "todos"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Couldn't get all todos")
		slog.Error("Error reading response body", "error", err)
	}

	var todos []models.Todo
	err = json.Unmarshal(body, &todos)
	if err != nil {
		fmt.Println("Couldn't read all todos")
		slog.Error("Error unmarshalling response body", "error", err)
	}
	return todos
}

func GetList(list string) []models.Todo {
	url := fmt.Sprintf("%stodos?list=%s", baseUrl, list)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Couldn't get all todos")
		slog.Error("Error reading response body", "error", err)
	}

	var todos []models.Todo
	err = json.Unmarshal(body, &todos)
	if err != nil {
		fmt.Println("Couldn't read all todos")
		slog.Error("Error unmarshalling response body", "error", err)
	}
	return todos
}

func PatchTodo(taskNumber int, patch models.TodoPatch) error {
	url := fmt.Sprintf("%stodo/%d", baseUrl, taskNumber)

	jsonData, err := json.Marshal(patch)
	if err != nil {
		slog.Error("Error marshalling JSON", "error", err)
		return fmt.Errorf("error marshalling JSON")
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Error("Error creating patch request", "error", err)
		return fmt.Errorf("error creating patch request")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error making patch request", "error", err)
		return fmt.Errorf("error making patch request")
	}
	defer resp.Body.Close()
	return nil
}

func Delete(taskNumber int) error {
	url := fmt.Sprintf("%stodo/%d", baseUrl, taskNumber)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		slog.Error("Error creating delete request", "error", err)
		return fmt.Errorf("error creating delete request")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error making delete request", "error", err)
		return fmt.Errorf("error making delete request")
	}
	defer resp.Body.Close()
	return nil
}
