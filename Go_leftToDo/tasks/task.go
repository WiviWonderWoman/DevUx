package tasks

type task struct {
	TaskType    string
	Description string
	Done        string
	DaysLeft    int
	SubTask     []SimpleTask
}

type TaskRepository interface {
	SetTask(taskType string, desc string, done string, days int, subTasks []SimpleTask) (task, error)
	Task() (task, error)
}

type TaskService struct {
}

func NewTaskService() TaskService {
	return TaskService{}
}

func (ts TaskService) SetTask(taskType string, desc string, done string, days int, subTasks []SimpleTask) (task, error) {

	t := task{
		TaskType:    taskType,
		Description: desc,
		Done:        done,
		DaysLeft:    days,
		SubTask:     subTasks,
	}
	return t, nil
}

func (ts TaskService) Task() task {
	t := task{}
	return t
}
