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

func (d DeadlineTask) SetTask(desc string, days int, subTasks []Task) (Task, error) {

	dt := Task{
		TaskType:    "D",
		description: desc,
		Done:        false,
		daysLeft:    days,
		// SubTask:     subTasks,
	}
	return dt, nil
}

func (d DeadlineTask) Task() ([]Task, error) {
	t := []Task{}
	return t, nil
}
