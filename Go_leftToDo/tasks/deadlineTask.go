package tasks

type DeadlineTask struct {
	TaskType    string
	Description string
	Done        bool
	DaysLeft    int
}

func NewDeadlineTask(desc string, done bool, days int) (DeadlineTask, error) {
	return DeadlineTask{
		TaskType:    "D",
		Description: desc,
		Done:        done,
		DaysLeft:    days,
	}, nil
}

func (d DeadlineTask) SetTask(taskType string, desc string, done bool, days int, subTasks []SimpleTask) (Task, error) {

	dt := Task{
		taskType:    "D",
		description: desc,
		done:        done,
		daysLeft:    days,
		// SubTask:     subTasks,
	}
	return dt, nil
}

func (d DeadlineTask) Task() ([]Task, error) {
	t := []Task{}
	return t, nil
}
