package lists

import "github.com/WiviWonderWoman/DevUx/Go/tasks"

type TodoList struct {
	ToDo    []tasks.Task
	Archive []tasks.Task
}

func NewToDoList() (TodoList, error) {
	return TodoList{
		ToDo:    []tasks.Task{},
		Archive: []tasks.Task{},
	}, nil
}
