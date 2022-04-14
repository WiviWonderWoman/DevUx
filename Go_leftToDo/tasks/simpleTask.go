package tasks

import "fmt"

// Struct with embedded Task
type SimpleTask struct {
	Task
	MarkAsDoneFn func(simple SimpleTask) SimpleTask
}

// Public constructor function
func NewSimpleTask(description string) (s SimpleTask) {
	s.TaskType = "S"
	s.Done = false
	s.Description = description
	return s
}

// Implement TaskRepository interface Create() - adds input from user as description.
func (s SimpleTask) Create() *SimpleTask {
	var input string
	fmt.Printf("\nAnge uppgift:\n")
	fmt.Scanln(&input)

	simple := NewSimpleTask(input)
	return &simple
}

// Displays SimpleTask in ChecklistTasks SubTask
func (s *SimpleTask) ShowSubTask(outer int, inner int) {
	if !s.Done {
		fmt.Println("[ ]\t", outer, " - ", inner, "\t", s.Description)
	} else if s.Done {
		fmt.Println("[X]\t", outer, " - ", inner, "\t", s.Description)
	}
}

// Implement TaskRepository interface MarkAsDone()
func (s *SimpleTask) MarkAsDone(simple SimpleTask) SimpleTask {
	return simple.MarkAsDoneFn(simple)
}

func (s SimpleTask) CreateMarkAsDone(simple SimpleTask) SimpleTask {
	simple.MarkAsDoneFn = func(simple SimpleTask) SimpleTask {
		return simple
	}
	return simple
}
