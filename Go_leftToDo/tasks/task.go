package tasks

import "fmt"

type Task struct {
	TaskType    string
	description string
	Done        bool
	daysLeft    int
	SubTask     []Task
}

type TaskRepository interface {
	SetTask(desc string, days int, subTasks []Task) (Task, error)
	Task() ([]Task, error)
	ShowTask(task Task, index int)
}

func (t Task) ShowTask(task Task, index int) {
	if task.Done == false {
		fmt.Println("[ ]\t", index, "\t", task.description)
	} else {
		fmt.Println("[X]\t", index, "\t", task.description)
	}
}

func (t Task) ShowSubTask(task Task, outer int, inner int) {
	if task.Done == false {
		fmt.Println("[ ]\t", outer, " - ", inner, "\t", task.description)
	} else {
		fmt.Println("[X]\t", outer, " - ", inner, "\t", task.description)
	}
}

func (t Task) MarkAsDone() {
	t.Done = !t.Done
}

func (t Task) Create(tt string) {
	t.TaskType = tt
	fmt.Println("Ange uppgift:")
	//TODO: read input
	input := ""
	t.description = input
	inputInt := 0
	fmt.Println("Ange dagar till deadline:")
	t.daysLeft = inputInt
}
