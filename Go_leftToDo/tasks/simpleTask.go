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

func (s SimpleTask) SetTask(desc string, days int, subTasks []Task) (Task, error) {

	st := Task{
		TaskType:    "S",
		description: desc,
		Done:        false,
		// DaysLeft:    days,
		// SubTask:     subTasks,
	}
	return st, nil
}

func (s SimpleTask) Task() ([]Task, error) {
	t := []Task{}
	return t, nil
}
