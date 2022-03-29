package tasks

type Task struct {
	taskType    string
	description string
	done        bool
	daysLeft    int
	subTask     []SimpleTask
}

type TaskRepository interface {
	SetTask(taskType string, desc string, done bool, days int, subTasks []SimpleTask) (Task, error)
	Task() ([]Task, error)
}
