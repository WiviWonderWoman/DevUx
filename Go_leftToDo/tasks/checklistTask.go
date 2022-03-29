package tasks

type ChecklistTask struct {
	TaskType    string
	Description string
	Done        string
	// DaysLeft    int
	SubTask []SimpleTask
}

func (c ChecklistTask) SetTask(taskType string, desc string, done string, days int, subTasks []SimpleTask) (task, error) {

	ct := task{
		TaskType:    "C",
		Description: desc,
		Done:        done,
		// DaysLeft:    days,
		SubTask: subTasks,
	}
	return ct, nil
}

func (c ChecklistTask) Task() (task, error) {
	t := task{}
	return t, nil
}
