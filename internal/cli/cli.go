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
		fmt.Println("What would you like to do? (add, list, complete [task number], delete [task number], exit)")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		command := strings.Split(input, " ")[0]

		switch command {
		case "add":
			task := input[4:]
			a.Add(task)
		case "list":
			a.ListToDos()
		case "complete":
			taskNumber, err := strconv.Atoi(input[9:])
			if err != nil {
				fmt.Println("please enter valid task number ie for task 1 'complete 1'")
			}
			a.MarkComplete(taskNumber)
		case "delete":
			taskNumber, err := strconv.Atoi(input[7:])
			if err != nil {
				fmt.Println("please enter valid task number ie for task 1 'complete 1'")
			}
			a.Delete(taskNumber)
		case "exit":
			fmt.Println("Exiting...")
			break appLoop
		case "edit":
			args := strings.Split(input, " ")
			taskNumber, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("please enter valid task number ie for task 1 'complete 1'")
			}
			editedTask := strings.Join(args[2:], " ")
			a.Edit(taskNumber, editedTask)
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

func (a App) Add(task string) {
	toDo := models.ToDo{Task: task, Completed: false}
	a.Store.Add(toDo)
}

func (a App) MarkComplete(i int) {
	err := a.Store.MarkComplete(i)
	if err != nil {
		fmt.Println("Couldn't mark complete: ", err)
	}
}

func (a App) Delete(i int) {
	err := a.Store.Delete(i)
	if err != nil {
		fmt.Println("Couldn't delete: ", err)
	}
}

func (a App) Edit(i int, edit string) {
	err := a.Store.EditToDo(i, edit)
	if err != nil {
		fmt.Println("Couldn't edit: ", err)
	}
}
