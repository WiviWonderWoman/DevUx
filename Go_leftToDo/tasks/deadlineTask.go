package tasks

type DeadlineTask struct {
	TaskType    string
	Description string
	Done        string
	DaysLeft    int
	// SubTask     []SimpleTask
}

func (d DeadlineTask) SetTask(taskType string, desc string, done string, days int, subTasks []SimpleTask) (task, error) {

	dt := task{
		TaskType:    "D",
		Description: desc,
		Done:        done,
		DaysLeft:    days,
		// SubTask:     subTasks,
	}
	return dt, nil
}

func (d DeadlineTask) Task() (task, error) {
	t := task{}
	return t, nil
}
