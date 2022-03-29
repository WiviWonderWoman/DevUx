package tasks

type SimpleTask struct {
	TaskType    string
	Description string
	Done        bool
}

func NewSimpleTask(desc string, done bool) (SimpleTask, error) {
	return SimpleTask{
		TaskType:    "S",
		Description: desc,
		Done:        done,
	}, nil
}

func (s SimpleTask) SetTask(taskType string, desc string, done bool, days int, subTasks []SimpleTask) (Task, error) {

	st := Task{
		taskType:    "S",
		description: desc,
		done:        done,
		// DaysLeft:    days,
		// SubTask:     subTasks,
	}
	return st, nil
}

func (s SimpleTask) Task() ([]Task, error) {
	t := []Task{}
	return t, nil
}
