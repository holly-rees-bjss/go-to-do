### Phase 1 - A self-contained single-user CLI To Do Store

>> - chapter 13 - the standard library for structured logging, time, io and Friends

- [x] CLI works directly with the In-Memory Data Store.
- [x] Allows the user to Add, and Remove a To Do Item with an optional due datetime.
- [x] Allows the user to mark a To Do Item as Not Started, In Progress or Completed along with the datetime when the To Do Item was updated.
- [x] The To Do Store automatically moves Complete Items to an Archive list.
- [x] The To Do Store automatically moves Overdue Items (current datetime > due datetime) to an Overdue list.
- [x] Allows the user to list all or some To Do Items.  For a reduced list this can be based on Status, or list (Archive, Overdue).
- [x] Include unit and integration tests validating the behaviour of the To Do Store.
- [x] Use [log/slog] to capture logging information (and save to a file), provide a [flag] to set the logging level


### Phase 2 - A client-server CLI persistent To Do Store

>> - chapter 13 - the standard library for structured encoding/json, net/http
>> - chapter 14 - context
>> - test chapter - fuzzing technique

- [ ] Use [net/http] to wrap the Data Store with the [V1 REST API]
        [x] update PatchTodo for different status' (complete, in progress)
        [x] get todos to return archived with ?list=archive
        [ ] get todos to return overdue
        [ ] check overdue middleware function to automatically add overdue item to overdue
- [ ] Use http server middleware and the [context] package to add a [github.com/google/uuid] TraceID which should be including in [log/slog] traceability of calls through the solution.
- [ ] Update the CLI App to use the REST API.
- [ ] Add an JSON Data Store and use a startup [flag] value to tell the server which data store to use.
- [ ] Include a fuzzing test to validate that the REST `POST /todo` can handle malformed values.

### Phase 3 - A multi-user Web App To Do Store

>> - chapter 12 - concurrency 
>> - chapter 11 - embed

- [ ] Add a [V2 REST API](./to-do-app-api-v2.yaml) that supports multiple users.
- [ ] Add a Web App using [html/template] that uses the [V2 REST API](./to-do-app-api-v2.yaml) that supports multiple users
- [ ] The Web App should use [embed] to support a single file deployment of the web server
- [ ] Add multi-user support to the CLI App
- [ ] Use the CSP pattern to support concurrent reads and concurrent safe write.
- [ ] Use Parallel tests to validate that the solution is concurrent safe.
- [ ] Use Parallel test to validate the solution is concurrent safe across multiple users.
- [ ] Include a benchmark test to measure the performance of the multi-user To Do Store.

#### Stretch goal

There is no need to implement authentication however, as a stretch goal, use [Okta] and [OAuth] 

### Phase 4 - Replace backend data store with a DB

- [ ] Add an additional data store implementation using an external DB (PostgreSQL)