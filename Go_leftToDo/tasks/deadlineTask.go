package tasks

import "fmt"

type DeadlineTask struct {
	TaskType    string
	Description string
	Done        bool
	DaysLeft    int
}

func NewDeadlineTask() DeadlineTask {
	return DeadlineTask{
		TaskType: "D",
		Done:     false,
	}
}

func NewDeadlineTaskWrapper() *TaskWrapper {
	return &TaskWrapper{
		DeadlineTask: NewDeadlineTask(),
	}
}

// func (d *DeadlineTask) SetTask(desc string, days int, subTasks []SimpleTask) (TaskWrapper, error) {
// 	tw := NewDeadlineTaskWrapper()
// 	tw.DeadlineTask.Description = desc
// 	tw.DeadlineTask.DaysLeft = days
// 	return *tw, nil
// }

func (d *DeadlineTask) Create() (TaskWrapper, error) {
	tw := NewDeadlineTaskWrapper()
	fmt.Println("Ange uppgift:")
	input := "" //TO-DO: read input
	tw.DeadlineTask.Description = input
	fmt.Println("Ange dagar till deadline:")
	inputInt := 0 //TO-DO: read input
	tw.DeadlineTask.DaysLeft = inputInt
	return *tw, nil
}

func (d *DeadlineTask) ShowTask(task TaskWrapper, index int) {
	panic("not implemented") // TODO: Implement
}

func (d *DeadlineTask) ShowSubTask(task TaskWrapper, outer int, inner int) {
	panic("not implemented") // TODO: Implement
}

func (d *DeadlineTask) MarkAsDone() {
	panic("not implemented") // TODO: Implement
}
