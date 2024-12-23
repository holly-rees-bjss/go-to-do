package v2

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
	api "todo_app/internal/cli/api_calls"
	"todo_app/internal/models"
)

type App struct {
	Logger *slog.Logger
}

func (a App) Run() {
	a.Logger.Info("CLI App starting")

	fmt.Println("Welcome to your CLI to-do app!\nHere's your To-Do list:")
	a.HandleList("list")

appLoop:
	for {
		fmt.Println("What would you like to do? (add [task] [optional: due date], list, complete [task number], in progress [task number], delete [task number], edit [task number] [new desciption], exit)")
		fmt.Println("Example: add finish to-do app due 29-11-2024")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		command := strings.Split(input, " ")[0]

		switch command {
		case "add":
			a.HandleAdd(input)

		case "list":
			a.HandleList(input)

		case "complete":
			a.HandleMarkComplete(input)

		case "progress":
			a.HandleMarkInProgress(input)

		case "delete":
			a.HandleDelete(input)

		case "edit":
			a.HandleEdit(input)

		case "exit":
			fmt.Println("Exiting...")
			break appLoop

		default:
			fmt.Println("your command was: " + command)
			fmt.Println("Please enter a valid command:")
		}
		fmt.Println("")
	}

}

func (a App) HandleList(input ...string) {
	a.Logger.Info("handling list")

	parts := strings.Split(input[0], " ")
	var selected = ""
	if len(parts) > 1 {
		selected = strings.TrimSpace(parts[1])
	}

	a.Logger.Debug("handle list",
		"command line input", parts,
		"selected", selected)

	switch selected {
	case "archive":
		archivedTodos := api.GetList("archive")
		if len(archivedTodos) == 0 {
			fmt.Println("There are no archived tasks.")
		}
		for i, todo := range archivedTodos {
			taskNum := strconv.Itoa(1 + i)
			fmt.Println(taskNum + ". " + todo.Task + " [Status: " + todo.Status + "]")
		}
	case "overdue":
		overdueTodos := api.GetList("overdue")
		if len(overdueTodos) == 0 {
			fmt.Println("There are no overdue tasks.")
		}
		for i, todo := range overdueTodos {
			taskNum := strconv.Itoa(1 + i)
			str := taskNum + ". " + todo.Task + " [Status: " + todo.Status + "]"
			if !todo.DueDate.IsZero() {
				str += " [Due: " + todo.DueDate.Format("02-01-2006") + "]"
			}
			fmt.Println(str)
		}
	default:
		todos := api.GetAll()

		for i, todo := range todos {
			taskNum := strconv.Itoa(1 + i)
			str := taskNum + ". " + todo.Task + " [Status: " + todo.Status + "]"
			if !todo.DueDate.IsZero() {
				str += " [Due: " + todo.DueDate.Format("02-01-2006") + "]"
			}
			fmt.Println(str)
		}
	}

}

func (a App) HandleAdd(input string) {
	a.Logger.Info("Handling add toDo", "input", input)

	parts := strings.Split(input, " ")

	var task string
	var toDo models.Todo

	if strings.Contains(input, "due") {
		layout := "02-01-2006"
		task = strings.Join(parts[1:len(parts)-2], " ")
		dueDate, err := time.Parse(layout, parts[len(parts)-1])
		if err != nil {
			fmt.Println("Couldn't add this, due date needs to be in the format dd-mm-yyy - please try again!")
			a.Logger.Error("Couldn't parse date", "error:", err)
			return
		}
		toDo = models.NewToDo(task, dueDate)
		a.Logger.Info("Todo with due date created", "toDo", toDo)
	} else {
		task = strings.Join(parts[1:], " ")
		toDo = models.Todo{Task: task, Status: "Not Started"}
		a.Logger.Info("Todo without due date created", "toDo", toDo)
	}

	api.Add(toDo)

	a.Logger.Info("toDo added to store", "toDo", toDo)
	fmt.Printf("Added %v to ToDo List.\n", task)
}

func (a App) HandleMarkComplete(input string) {
	a.Logger.Info("handling mark complete", "input", input)

	taskNumber, err := strconv.Atoi(input[9:])
	if err != nil {
		fmt.Println("please enter valid task number ie for task 1 'complete 1'")
		a.Logger.Error("invalid task number", "error:", err)
	}

	patch := models.TodoPatch{Status: "Completed"}
	err = api.PatchTodo(taskNumber, patch)
	if err != nil {
		fmt.Println("Oops! We couldn't mark that task as complete.")
		a.Logger.Error("Couldn't mark complete", "error:", err)
	}
	fmt.Printf("Marked task number %v as Complete.\n", taskNumber)
}

func (a App) HandleMarkInProgress(input string) {
	a.Logger.Info("handling mark in progress", "input", input)

	taskNumber, err := strconv.Atoi(input[9:])
	if err != nil {
		fmt.Println("please enter valid task number ie for task 1 'complete 1'")
		a.Logger.Error("invalid task number", "error:", err)
	}

	patch := models.TodoPatch{Status: "In Progress"}
	err = api.PatchTodo(taskNumber, patch)
	if err != nil {
		fmt.Println("Oops! We couldn't mark that task as in progress.")
		a.Logger.Error("Couldn't mark todo as in progress", "error", err)
	}
	fmt.Printf("Marked task number %v as In Progress.\n", taskNumber)
}

func (a App) HandleDelete(input string) {
	a.Logger.Info("handling delete", "input", input)

	taskNumber, err := strconv.Atoi(input[7:])
	if err != nil {
		fmt.Println("please enter valid task number ie for task 1 'complete 1'")
		a.Logger.Error("invalid task number", "error:", err)
	}

	err = api.Delete(taskNumber)
	if err != nil {
		fmt.Println("Oops! We couldn't delete that todo.")
		a.Logger.Error("Couldn't delete", "error", err)
	}
	fmt.Printf("Deleted task number %v.\n", taskNumber)
}

func (a App) HandleEdit(input string) {
	a.Logger.Info("handling edit", "input", input)

	args := strings.Split(input, " ")
	taskNumber, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("please enter valid task number ie for task 1 'complete 1'")
		a.Logger.Error("invalid task number", "error:", err)
	}
	editedTask := strings.Join(args[2:], " ")

	patch := models.TodoPatch{Task: editedTask}
	err = api.PatchTodo(taskNumber, patch)
	if err != nil {
		fmt.Println("Oops! We couldn't delete that todo.")
		a.Logger.Error("Couldn't edit", "error", err)
	}
	fmt.Printf("Edited task number %v.\n", taskNumber)
}
