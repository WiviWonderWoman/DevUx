package tasks

import "fmt"

type Task struct {
	TaskType    string        // describes task type user choose
	Description string        // what action / ToDo task user inputs
	Done        bool          // indicates if task is done
	SubTask     []*SimpleTask // if task is of type checklist, a slice with subtasks
}

type TaskRepository interface {
	ShowTask(index int)        //  Displays Task
	MarkAsDone(t Task) Task    // Marks task as Done
	Create(description string) // Add info from user as description
}

//  Displays Task
func (t Task) ShowTask(index int) {

	if !t.Done && t.TaskType == "C" {
		fmt.Println(" - \t", index, "\t", t.Description)

		for i := 0; i < len(t.SubTask); i++ {
			sub := t.SubTask[i]
			sub.ShowSubTask(index, i+1)
		}
	} else if !t.Done && t.TaskType == "S" {
		fmt.Println("[ ]\t", index, "\t", t.Description)
	} else if t.Done {
		fmt.Println("[X]\t", index, "\t", t.Description)
	}
}

// Marks task as Done
func (t Task) MarkAsDone(task Task) Task {
	task.Done = !task.Done
	return task
}
