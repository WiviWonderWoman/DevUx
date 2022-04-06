package tasks

import (
	"fmt"
	"strings"
)

type ChecklistTask struct {
	Task
}

func newChecklistTask(description string) *ChecklistTask {
	c := ChecklistTask{}
	c.TaskType = "C"
	c.Done = false
	c.Description = description
	c.SubTask = []*SimpleTask{}
	return &c
}

func (c ChecklistTask) Create() *ChecklistTask {
	fmt.Printf("\nAnge Rubrik-uppgift:\n")
	var input string
	fmt.Scanln(&input)
	checklist := newChecklistTask(strings.ToUpper(input))
	stopInput := false
	fmt.Println("Ange uppgifter, separerat med enter. [0] för att slutföra checklistan.")
	for !stopInput {
		var subInput string
		fmt.Scanln(&subInput)
		if subInput != "0" {
			sub := NewSimpleTask(subInput)
			checklist.SubTask = append(checklist.SubTask, sub)
		} else if subInput == "0" {
			stopInput = true
		}
	}
	return checklist
}
