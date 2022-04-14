package tasks

import (
	"fmt"
)

type SimpleTask struct {
	Task
	MarkAsDoneFn func(simple SimpleTask) SimpleTask
}

func NewSimpleTask(description string) *SimpleTask {
	s := SimpleTask{}
	s.TaskType = "S"
	s.Done = false
	s.Description = description
	return &s
}

func (s SimpleTask) Create() *SimpleTask {
	fmt.Printf("\nAnge uppgift:\n")
	var input string
	fmt.Scanln(&input)
	simple := NewSimpleTask(input)
	return simple
}

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
