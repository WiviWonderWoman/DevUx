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
	MarkAsDone(t Task) Task
	Create(description string)
}

func (t Task) ShowTask(index int) {
	if !t.Done && t.TaskType == "C" {
		fmt.Println(" - \t\t", index, "\t\t", t.Description)
	} else if !t.Done && t.TaskType == "S" {
		fmt.Println("[ ]\t\t", index, "\t\t", t.Description)
	} else if t.Done {
		fmt.Println("[X]\t\t", index, "\t\t", t.Description)
	}
}

func (t Task) MarkAsDone(task Task) Task {
	fmt.Println("DONE, TASK", task)
	task.Done = !task.Done
	return task
}
