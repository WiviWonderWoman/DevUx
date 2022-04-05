package tasks

import "fmt"

type SimpleTask struct {
	Task
}

func NewSimpleTask(description string) *SimpleTask {
	return &SimpleTask{
		Task{
			TaskType:    "S",
			Description: description,
		},
	}
}

func (s SimpleTask) Create() {
	fmt.Println("Ange uppgift:")
	//TODO: read input from console
	input := ""
	s.Description = input

}
