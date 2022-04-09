// Package tasks handels all Task types and their functions & methods
package tasks

import "fmt"

// A struct for embeding in other types of Tasks
type Task struct {
	TaskType    string        // describes task type user choose
	Description string        // what action/ToDo-task user inputs
	Done        bool          // indicates if task is done
	SubTask     []*SimpleTask // if task is of type checklist, a slice with subtasks
}

// Marks task as Done
func (t Task) MarkAsDone(task Task) Task {
	task.Done = !task.Done
	return task
}

// Interface that specifies method that diffrent Tasks types implement
type TaskRepository interface {
	MarkAsDone(t Task) Task           // Marks task as Done
	Create(description string) *Task  // Takes info from user
	ShowTask(index int)               // Displays any Task
	ShowSubTask(outer int, inner int) // Displas simple Task

}

// Displays any Task
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

/*

 */
