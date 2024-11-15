### Phase 1 - A self-contained single-user CLI To Do Store

- [x] CLI works directly with the In-Memory Data Store.
- [x] Allows the user to Add, and Remove a To Do Item with an optional due datetime.
- [x] Allows the user to mark a To Do Item as Not Started, In Progress or Completed along with the datetime when the To Do Item was updated.
- [x] The To Do Store automatically moves Complete Items to an Archive list.
- [x] The To Do Store automatically moves Overdue Items (current datetime > due datetime) to an Overdue list.
- [x] Allows the user to list all or some To Do Items.  For a reduced list this can be based on Status, or list (Archive, Overdue).
- [ ] Include unit and integration tests validating the behaviour of the To Do Store.
- [ ] Use [log/slog] to capture logging information, provide a [flag] to set the logging level