package tasks

import "fmt"

type ChecklistTask struct {
	Task
	// tr TaskRepository
}

func NewChecklistTask(description string) *ChecklistTask {
	return &ChecklistTask{
		Task{
			TaskType:    "C",
			Done:        false,
			Description: description,
			SubTask:     []SimpleTask{},
		},
	}
}

func (c ChecklistTask) Create() {
	fmt.Println("Ange Rubrik-uppgift:")
	//TODO:  input = Console.ReadLine();
	input := ""
	c.Description = input
	fmt.Println("Ange uppgifter, separerat med enter.\n\n[0] för att slutföra checklistan.")
	//TODO:  input = Console.ReadLine();
	stopInput := false
	for !stopInput {
		if input == "0" {
			stopInput = true
			return
		}
		// input = Console.ReadLine();
		sub := NewSimpleTask(input)
		c.SubTask = append(c.SubTask, *sub)
	}
}
