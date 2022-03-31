package tasks

import "fmt"

type Task struct {
	TaskType    string
	Description string
	Done        bool
	DaysLeft    int
	SubTask     []Task
}

type TaskRepository interface {
	SetTask(desc string, days int, subTasks []Task) (Task, error)
	Create() (Task, error)
	ShowTask(task Task, index int)
	ShowSubTask(task Task, outer int, inner int)
	MarkAsDone()
}

func (t Task) ShowTask(task Task, index int) {
	if task.Done == false {
		fmt.Println("[ ]\t", index, "\t", task.Description)
	} else {
		fmt.Println("[X]\t", index, "\t", task.Description)
	}
}

func (t Task) ShowSubTask(task Task, outer int, inner int) {
	if task.Done == false {
		fmt.Println("[ ]\t", outer, " - ", inner, "\t", task.Description)
	} else {
		fmt.Println("[X]\t", outer, " - ", inner, "\t", task.Description)
	}
}

func (t Task) MarkAsDone() {
	t.Done = !t.Done
}

func (t Task) Create() (Task, error) {
	// t.TaskType = tt
	fmt.Println("Ange uppgift:")
	input := "" //TODO: read input
	t.Description = input
	inputInt := 0 //TODO: read input
	fmt.Println("Ange dagar till deadline:")
	t.DaysLeft = inputInt
	return t, nil
}
