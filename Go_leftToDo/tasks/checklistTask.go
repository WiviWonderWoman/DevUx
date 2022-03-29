package tasks

type ChecklistTask struct {
	TaskType    string
	Description string
	Done        bool
	SubTask     []SimpleTask
}

func NewChecklistTask(desc string, done bool, days int, subTasks []SimpleTask) (ChecklistTask, error) {
	return ChecklistTask{
		TaskType:    "D",
		Description: desc,
		Done:        done,
		SubTask:     subTasks,
	}, nil
}

func (c ChecklistTask) SetTask(taskType string, desc string, done bool, days int, subTasks []SimpleTask) (Task, error) {

	ct := Task{
		taskType:    "C",
		description: desc,
		done:        done,
		// DaysLeft:    days,
		subTask: subTasks,
	}
	return ct, nil
}

func (c ChecklistTask) Task() ([]Task, error) {
	t := []Task{}
	return t, nil
}
