package tasks

import "fmt"

type Task struct {
	TaskType    string
	Description string
	Done        bool
	DaysLeft    int
	SubTask     []*SimpleTask
	// MarkAsDone  func(t Task) Task
}

type TaskRepository interface {
	ShowTask(index int)
	ShowSubTask(outer int, inner int)
	MarkAsDone(t Task) Task
	Create(description string)
}

func (t Task) ShowTask(index int) {
	if !t.Done {
		fmt.Println("[ ]\t", index, "\t", t.Description)

	} else if t.Done {
		fmt.Println("[X]\t", index, "\t", t.Description)
	}

}

func (t Task) MarkAsDone(task Task) Task {
	fmt.Println("DONE, TASK", task)
	task.Done = !task.Done
	return task
}
