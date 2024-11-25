# Go To-Do App

This is my Go To-Do app.

## CLI Application

The CLI application can be run by using the command `go run main.go cli` or `./main cli` (after building). <br>
As explained in the start menu, you can navigate by typing the following commands:

* add [task] [optional: due date]       
    - eg. `add finish to-do app due 29-11-2024` will add 'finish to-do app' with that due date
* list [all/archive/overdue]            
    - eg. `list overdue` will list all overdue items
* complete [task number]                
    - eg. `complete 2` will mark Task 2 as Complete
* in progress [task number]             
    - eg. `in progress 1` will mark Task 1 as In Progress
* delete [task number]                  
    - eg. `delete 1` will delete Task 1
* edit [task number] [new desciption]   
    - eg. `edit 3 learn go` will update Task 3 to be 'learn go'
* exit                                  
    - eg. `exit` will exit the app

When todo items are marked as complete, they are automatically moved to the Archive list (which you can view with `list archive`). Items that are overdue are also automatically added to the Overdue list (the program checks for any overdue items every time a command is executed).

The CLI app is run with the default logging level as Error, but can specify logging level (debug, info, warn, error) using the logLevel flag eg. `go run main.go -loglevel=debug cli`. Logs for the CLI app are stored in file `cli_logs.log`, storing logs for the most recent run of the application.

## A client-server CLI persistent To Do Store

The API application can be run by using the command `go run main.go api` or `./main api` (after building), and will run on localhost port 8080. <br>
This API has the following endpoints:
* /api/todos will return JSON of all todos. You can also add optional query parameters to get the Overdue or Archive lists.
    - /api/todos/archive
    - /api/todos/overdue
    