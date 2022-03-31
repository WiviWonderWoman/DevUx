package tasks

// type Task struct {
// 	TaskType    string
// 	Description string
// 	Done        bool
// 	DaysLeft    int
// 	SubTask     []Task
// }

type TaskWrapper struct {
	SimpleTask    SimpleTask
	DeadlineTask  DeadlineTask
	ChecklistTask ChecklistTask
}

type TaskRepository interface {
	// SetTask(desc string, days int, subTasks []SimpleTask) (TaskWrapper, error)
	Create() (TaskWrapper, error)
	ShowTask(task TaskWrapper, index int)
	ShowSubTask(task TaskWrapper, outer int, inner int)
	MarkAsDone()
}

/*






 */
// func (tw *TaskWrapper) SetSinmpleTask(desc string, days int, subTasks []SimpleTask) (SimpleTask, error) {
// 	return tw.SimpleTask.SetSinmpleTask(desc, days, subTasks)
// }
// func (s SimpleTask) SetSinmpleTask(desc string, days int, subTasks []SimpleTask) (SimpleTask, error) {

// 	s.TaskType = "S"
// 	s.Description = desc
// 	s.Done = false

// 	return s, nil
// }

//____________________________________________________
// func (t Task) Create() (Task, error) {
// 	fmt.Println("Ange uppgift:")
// 	input := "" // TO-DO: read input
// 	t.Description = input
// 	inputInt := 0 // TO-DO: read input
// 	fmt.Println("Ange dagar till deadline:")
// 	t.DaysLeft = inputInt
// 	return t, nil
// }

// func (t Task) ShowTask(task Task, index int) {
// 	if task.Done == false {
// 		fmt.Println("[ ]\t", index, "\t", task.Description)
// 	} else {
// 		fmt.Println("[X]\t", index, "\t", task.Description)
// 	}
// }

// func (t Task) ShowSubTask(task Task, outer int, inner int) {
// 	if task.Done == false {
// 		fmt.Println("[ ]\t", outer, " - ", inner, "\t", task.Description)
// 	} else {
// 		fmt.Println("[X]\t", outer, " - ", inner, "\t", task.Description)
// 	}
// }

// func (t Task) MarkAsDone() {
// 	t.Done = !t.Done
// }
