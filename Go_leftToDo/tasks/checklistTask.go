package tasks

type ChecklistTask struct {
	TaskType    string
	Description string
	Done        bool
	SubTask     []SimpleTask
}

// func NewChecklistTask(desc string, done bool, days int, subTasks []SimpleTask) (ChecklistTask, error) {
// 	return ChecklistTask{
// 		TaskType:    "D",
// 		Description: desc,
// 		Done:        done,
// 		SubTask:     subTasks,
// 	}, nil
// }

// func (c ChecklistTask) SetTask(desc string, days int, subTasks []Task) (Task, error) {

// 	ct := Task{
// 		TaskType:    "C",
// 		Description: desc,
// 		Done:        false,
// 		SubTask:     subTasks,
// 	}
// 	return ct, nil
// }
