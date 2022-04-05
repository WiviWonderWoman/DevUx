package tasks

import "fmt"

type SimpleTask struct {
	Task
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

func (s SimpleTask) ShowTask(index int) {
	if !s.Done {
		fmt.Println("[ ]\t", index, "\t", s.Description)
	} else if s.Done {
		fmt.Println("[X]\t", index, "\t", s.Description)
	}

}

func (s SimpleTask) ShowSubTask(outer int, inner int) {

	if !s.Done {
		fmt.Println("\t[ ]\t", outer, " - ", inner, "\t", s.Description)
	} else if s.Done {
		fmt.Println("\t[X]\t", outer, " - ", inner, "\t", s.Description)
	}

}

func MarkAsDone(s SimpleTask) {
	s.Done = !s.Done
}
