# Go To-Do App

This is my Go To-Do app.

## Phase 1: CLI Application

The CLI application can be run by using the command `go run main.go cli` or `./main cli` (after building). <br>
As explained in the start menu, you can navigate by typing the following commands:

* add [task] [optional: due date]       
    - eg. `add finish to-do app due 29-11-2024` will add 'finish to-do app' with that due date
* list [all/archive/overdue]            
    - eg. `list overdue` will list all overdue items
* complete [task number]                
    - eg. `complete 2` will mark Task 2 as Complete
* in progress [task number]             
    - eg. `progress 1` will mark Task 1 as In Progress
* delete [task number]                  
    - eg. `delete 1` will delete Task 1
* edit [task number] [new desciption]   
    - eg. `edit 3 learn go` will update Task 3 to be 'learn go'
* exit                                  
    - eg. `exit` will exit the app

When todo items are marked as complete, they are automatically moved to the Archive list (which you can view with `list archive`). Items that are overdue are also automatically added to the Overdue list (the program checks for any overdue items every time a command is executed).

The CLI app is run with the default logging level as Error, but can specify logging level (debug, info, warn, error) using the logLevel flag eg. `go run main.go -loglevel=debug cli`. Logs for the CLI app are stored in file `cli_logs.log`, storing logs for the most recent run of the application.

## Phase 2: A client-server CLI persistent To Do Store

The API application can be run by using the command `go run main.go api` or `./main api` (after building), and will run on localhost port 8080. <br>
This API has the following endpoints:
* GET /api/todos will return JSON of all todos. You can also add optional query parameters to get the Overdue or Archive lists.
    - /api/todos?list=archive
    - /api/todos?list=overdue
* POST /api/todo to post a new todo
* PATCH /api/todo/{i} to update a specific todo (Task, Status or both) 
* DELETE /api/todo/{i} to delete a specific todo

As with the CLI app, when a todo is marked complete (via a PATCH request), it is automatically added to the Archive list. This API also uses a middleware that checks for any overdue todos every request call, so any overdue todos are automatically added to the Overdue list. 

The API app is also run with the default logging level as Error, but again you can specify logging level (debug, info, warn, error) using the logLevel flag eg. `go run main.go -loglevel=debug api`. 

This phase of the app also introduces persistent data storage with a json store. You can specify which data store you'd like using the `-store` flag and specifying `inmemory` or `json` (eg. `-store=json`). This is also true of the version 1 CLI app, eg. `go run main.go -store=json cli` will run the CLI app with json store. However the version 2 CLI app wraps the API so that relies on whichever data store you select when you run the API app, and thus cannot be run independently - you must start up the API first.
    
## Phase 3: A web app that wraps the API

This React app consumes the API. To use, first make sure to start the api `go run main.go api` (this defaults to using the JSON memory store), to run the server locally. 
To run the React app locally, you will need Node.js installed. Install the required dependencies with `npm i`, navigate to the `web_app` directory and launch using `npm run dev`.
This web app allows you to view all the todos, mark as complete or in progress, delete a todo and add a new todo. Data is persistent and saved to the json store.