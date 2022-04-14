package tasks

import (
	"fmt"
)

// struct with Task embedded
type SimpleTask struct {
	Task
	MarkAsDoneFn func(simple SimpleTask) SimpleTask
}

// Constructor function
func NewSimpleTask(description string) *SimpleTask {
	s := SimpleTask{}
	s.TaskType = "S"
	s.Done = false
	s.Description = description
	return &s
}

// Adds input from user as description.
func (s SimpleTask) Create() *SimpleTask {
	fmt.Printf("\nAnge uppgift:\n")
	var input string
	fmt.Scanln(&input)
	simple := NewSimpleTask(input)
	return simple
}

// Displays SimpleTask in ChecklistTasks subTask
func (s *SimpleTask) ShowSubTask(outer int, inner int) {
	if !s.Done {
		fmt.Println("[ ]\t", outer, " - ", inner, "\t", s.Description)
	} else if s.Done {
		fmt.Println("[X]\t", outer, " - ", inner, "\t", s.Description)
	}
}

func (s *SimpleTask) MarkAsDone(simple SimpleTask) SimpleTask {
	return simple.MarkAsDoneFn(simple)
}

func (s SimpleTask) CreateMarkAsDone(simple SimpleTask) SimpleTask {
	simple.MarkAsDoneFn = func(simple SimpleTask) SimpleTask {
		return simple
	}
	return simple
}
