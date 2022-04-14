package tasks

import (
	"fmt"
	"strings"
)

// Struct with embedded Task
type ChecklistTask struct {
	Task
}

// private Constructor function
func newChecklistTask(description string) (c ChecklistTask) {
	c.TaskType = "C"
	c.Done = false
	c.Description = description
	c.SubTask = []*SimpleTask{}
	return c
}

// Implement TaskRepository interface - adds input from user as description.
func (c ChecklistTask) Create() *ChecklistTask {
	var input string
	fmt.Printf("\nAnge Rubrik-uppgift:\n")
	fmt.Scanln(&input)

	checklist := newChecklistTask(strings.ToUpper(input))

	fmt.Println("Ange uppgifter, separerat med enter. [0] för att slutföra checklistan.")
	stopInput := false
	for !stopInput {
		var subInput string
		fmt.Scanln(&subInput)
		if subInput != "0" {
			sub := NewSimpleTask(subInput)
			checklist.SubTask = append(checklist.SubTask, &sub)
		} else if subInput == "0" {
			stopInput = true
		}
	}
	return &checklist
}
