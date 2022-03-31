package tasks

import "fmt"

type TaskWrapper struct {
	simpleTask  *SimpleTask
	TaskType    string
	Description string
	Done        bool
	DaysLeft    int
	SubTask     []Task
}

func NewTaskWrapper(simpleTask *SimpleTask) *TaskWrapper {
	return &TaskWrapper{
		simpleTask: simpleTask,
	}
}

func (tw *TaskWrapper) SetTask(desc string, days int, subTasks []Task) (SimpleTask, error) {
	return tw.simpleTask.SetTask(desc, days, subTasks)
}

func (tw *TaskWrapper) Create() (SimpleTask, error) {
	return tw.simpleTask.Create()
}

func (tw *TaskWrapper) ShowTask(task TaskWrapper, index int) {
	if task.Done == false {
		fmt.Println("[ ]\t", index, "\t", task.Description)
	} else {
		fmt.Println("[X]\t", index, "\t", task.Description)
	}
}

func (tw *TaskWrapper) MarkAsDone() {
	tw.simpleTask.MarkAsDone()
}
