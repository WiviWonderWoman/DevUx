package tasks

type SimpleTask struct {
	TaskType    string
	Description string
	Done        string
	// DaysLeft    int
	// SubTask     []SimpleTask
}

func (s SimpleTask) SetTask(taskType string, desc string, done string, days int, subTasks []SimpleTask) (task, error) {

	st := task{
		TaskType:    "S",
		Description: desc,
		Done:        done,
		// DaysLeft:    days,
		// SubTask:     subTasks,
	}
	return st, nil
}

func (s SimpleTask) Task() (task, error) {
	t := task{}
	return t, nil
}
