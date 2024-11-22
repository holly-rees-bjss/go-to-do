# Go To-Do App

This is my Go To-Do app.

## CLI Application

The CLI application can be run by using the command `go run main.go cli` or `./main cli` (after building). <br>
As explained in the start menu, you can navigate by typing the following commands:

* add [task] [optional: due date]
* list [all/archive/overdue]
* complete [task number]
* in progress [task number]
* delete [task number]
* edit [task number] [new desciption]
* exit

When todo items are marked as complete, they are automatically moved to the Archive list (which you can view with `list archive`). Items that are overdue are also automatically added to the Overdue list (the program checks for any overdue items every time a command is executed).

You can run with logs and specify logging level: `go run main.go -loglevel=debug cli`
