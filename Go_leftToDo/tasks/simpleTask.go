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

func (s SimpleTask) Create() *SimpleTask {
	fmt.Printf("\nAnge uppgift:\n")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("Du angav uppgift: %s\n", input)
	simple := NewSimpleTask(input)
	// s.Task.Description = input
	return simple
}
