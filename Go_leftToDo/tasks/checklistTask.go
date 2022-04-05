package tasks

import "fmt"

type ChecklistTask struct {
	Task
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

// func (c ChecklistTask) ShowTask(outer int) {
// 	if !c.Done {
// 		fmt.Println("[ ]\t", outer, "\t", c.Description)
// 		return
// 	}
// 	fmt.Println("[X]\t", outer, "\t", c.Description)
// 	inner := 0
// 	for _, sub := range c.SubTask {
// 		inner++
// 		sub.ShowSubTask(outer, inner)
// 	}
// }
