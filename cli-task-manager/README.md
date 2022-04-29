# CLI Task Manager
Build a CLI application that can be used to manage your TODOs in the terminal.

```
Task is a CLI task manager

Usage:
  task [command]

Available Commands:
  add         Add a task to the list of tasks
  completion  Generate the autocompletion script for the specified shell
  do          marks a task as complete
  help        Help about any command
  list        lists all of our incomplete tasks

Flags:
  -h, --help   help for task

Use "task [command] --help" for more information about a command.
```

## Usage
  `make task action=<list/add/do> value=<"add some task">`
### Flags
  1. `action` the action to perform - **list / add / do**
  2. `value` the value to perform the action with. 
  