package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo_app/internal/models"
)

type App struct {
	Store models.Store
}

func (a App) Run() {
	fmt.Println("Welcome to your CLI to-do app!\nHere's your To-Do list:")
	a.ListToDos()

appLoop:
	for {
		fmt.Println("What would you like to do? (add, list, complete [task number], delete [task number], edit [task number] [new desciption], exit)")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		command := strings.Split(input, " ")[0]

		switch command {
		case "add":
			a.HandleAdd(input)

		case "list":
			a.ListToDos()

		case "complete":
			a.HandleMarkComplete(input)

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

func (a App) ListToDos() {
	todos := a.Store.GetTodos()

	for i, todo := range todos {
		taskNum := strconv.Itoa(1 + i)
		fmt.Println(taskNum + ". " + todo.Task + " [Completed: " + strconv.FormatBool(todo.Completed) + "]")
	}
}

func (a App) HandleAdd(input string) {
	task := input[4:]
	toDo := models.ToDo{Task: task, Completed: false}
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
