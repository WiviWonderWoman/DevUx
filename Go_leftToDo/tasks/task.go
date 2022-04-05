package tasks

import "fmt"

type Task struct {
	TaskType    string
	Description string
	Done        bool
	DaysLeft    int
	SubTask     []*SimpleTask
}

type TaskRepository interface {
	ShowTask(index int)
	ShowSubTask(outer int, inner int)
	MarkAsDone()
	Create(description string)
}

func (t Task) ShowTask(index int) {
	if !t.Done {
		fmt.Println("[ ]\t", index, "\t", t.Description)

	} else if t.Done {
		fmt.Println("[X]\t", index, "\t", t.Description)
	}

}

// func (t Task) ShowSubTask(outer int, inner int) {

// 	if !t.Done {
// 		fmt.Println("\t[ ]\t", outer, " - ", inner, "\t", t.Description)
// 		return
// 	} else if t.Done {
// 	fmt.Println("\t[X]\t", outer, " - ", inner, "\t", t.Description)
// }
//
// }

func (t Task) MarkAsDone() {
	t.Done = !t.Done
}
