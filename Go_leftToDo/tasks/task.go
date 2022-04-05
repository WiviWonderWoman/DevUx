package tasks

import "fmt"

type Task struct {
	TaskType    string
	Description string
	Done        bool
	DaysLeft    int
	SubTask     []SimpleTask
}

type TaskRepository interface {
	ShowTask(index int)
	ShowSubTask(index int)
	MarkAsDone()
	Create()
}

func (t Task) ShowTask(index int) {
	if !t.Done {
		fmt.Println("[ ]\t", index, "\t", t.Description)
	} else {
		fmt.Println("[X]\t", index, "\t", t.Description)
	}
}

func (t Task) ShowSubTask(outer int) {
	inner := 0
	for _, sub := range t.SubTask {
		inner++
		if !t.Done {
			fmt.Println("[ ]\t", outer, " - ", inner, "\t", sub.Description)
			return
		}
		fmt.Println("[X]\t", outer, " - ", inner, "\t", sub.Description)
	}
}

func (t Task) MarkAsDone() {
	t.Done = !t.Done
}
