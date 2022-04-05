package tasks

import "fmt"

type Task struct {
	TaskType    string
	Description string
	Done        bool
	DaysLeft    int
	SubTask     []SimpleTask
}

type TaskRepository interface {
	ShowTask(task Task, index int)
	ShowSubTask(subTask []Task, outerIndex int, innerIndex int)
	MarkAsDone()
	Create() string
}

func (t Task) ShowTask(index int) {
	if !t.Done {
		fmt.Println("[ ]\t", index, "\t", t.Description)
	} else {
		fmt.Println("[X]\t", index, "\t", t.Description)
	}
}

func (t Task) ShowSubTask(subTask []SimpleTask, outerIndex int) {
	innerIndex := 0
	for _, sub := range subTask {
		innerIndex++
		if !t.Done {
			fmt.Println("[ ]\t", outerIndex, " - ", innerIndex, "\t", sub.Description)
		} else {
			fmt.Println("[X]\t", outerIndex, " - ", innerIndex, "\t", sub.Description)
		}
	}
}

func (t Task) MarkAsDone() {
	t.Done = !t.Done
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
