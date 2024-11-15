package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"todo_app/internal/models"
)

type App struct {
	Store models.Store
}

func (a App) Run() {
	fmt.Println("Welcome to your CLI to-do app!\nHere's your To-Do list:")
	a.HandleList()

appLoop:
	for {
		fmt.Println("What would you like to do? (add [task] [optional: due date], list, complete [task number], in progress [task number], delete [task number], edit [task number] [new desciption], exit)")
		fmt.Println("Example: add finish to-do app 22-11-2024")
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

		case "in progress":
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
	}

}

func (a App) HandleList(input ...string) {
	parts := strings.Split(input[0], " ")
	selected := strings.TrimSpace(parts[1])

	switch selected {
	case "archive":

		for i, todo := range a.Store.GetArchive() {
			taskNum := strconv.Itoa(1 + i)
			fmt.Println(taskNum + ". " + todo.Task + " [Status: " + todo.Status + "]")
		}
	default:
		todos := a.Store.GetTodos()

		for i, todo := range todos {
			taskNum := strconv.Itoa(1 + i)
			fmt.Println(taskNum + ". " + todo.Task + " [Status: " + todo.Status + "]")
		}
	}

}

func (a App) HandleAdd(input string) {
	parts := strings.Split(input, " ")
	task := strings.Join(parts[1:len(parts)-1], " ")

	layout := "02-01-2006"
	dueDate, err := time.Parse(layout, parts[len(parts)-1])

	var toDo models.Todo

	if err != nil {
		task = strings.Join(parts[1:], " ")
		toDo = models.Todo{Task: task, Status: "Not Started"}
	} else {

		toDo = models.NewToDo(task, dueDate)
	}

	a.Store.Add(toDo)
}

func (a App) HandleMarkComplete(input string) {
	taskNumber, err := strconv.Atoi(input[9:])
	if err != nil {
		fmt.Println("please enter valid task number ie for task 1 'complete 1'")
	}

	err = a.Store.MarkComplete(taskNumber)
	if err != nil {
		fmt.Println("Couldn't mark complete: ", err)
	}
}

func (a App) HandleMarkInProgress(input string) {
	taskNumber, err := strconv.Atoi(input[12:])
	if err != nil {
		fmt.Println("please enter valid task number ie for task 1 'complete 1'")
	}

	err = a.Store.MarkInProgress(taskNumber)
	if err != nil {
		fmt.Println("Couldn't mark complete: ", err)
	}
}

func (a App) HandleDelete(input string) {
	taskNumber, err := strconv.Atoi(input[7:])
	if err != nil {
		fmt.Println("please enter valid task number ie for task 1 'complete 1'")
	}

	err = a.Store.Delete(taskNumber)
	if err != nil {
		fmt.Println("Couldn't delete: ", err)
	}
}

func (a App) HandleEdit(input string) {
	args := strings.Split(input, " ")
	taskNumber, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("please enter valid task number ie for task 1 'complete 1'")
	}
	editedTask := strings.Join(args[2:], " ")

	err = a.Store.EditToDo(taskNumber, editedTask)
	if err != nil {
		fmt.Println("Couldn't edit: ", err)
	}
}
