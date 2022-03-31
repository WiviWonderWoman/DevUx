package tasks

import "fmt"

type SimpleTask struct {
	TaskType    string
	Description string
	Done        bool
}

func NewSimpleTask() SimpleTask {
	return SimpleTask{
		TaskType: "S",
		Done:     false,
	}
}

func (s SimpleTask) SetTask(desc string, days int, subTasks []Task) (SimpleTask, error) {

	s.TaskType = "S"
	s.Description = desc
	s.Done = false

	return s, nil
}

func (s SimpleTask) Create() (SimpleTask, error) {
	fmt.Println("Ange uppgift:")
	input := "" //TODO: read input
	s.Description = input
	return s, nil
}

func (s SimpleTask) MarkAsDone() {
	s.Done = !s.Done
}
