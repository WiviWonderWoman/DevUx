// Packafe lists handles all lists and their methods & functions
package lists

import (
	"fmt"
	"strconv"

	"github.com/WiviWonderWoman/DevUx/Go/tasks"
)

// A struct that for lists
type TaskList struct {
	ToDoList []tasks.Task // Tasks ToDo
	Archive  []tasks.Task // Archived done Tasks
}

// Public constructor function
func NewTaskList() TaskList {
	return TaskList{
		ToDoList: []tasks.Task{},
		Archive:  []tasks.Task{},
	}
}

// Adds new Task to ToDoList
func (tl *TaskList) AddToDoTask(task tasks.Task) {
	tl.ToDoList = append(tl.ToDoList, task)
}

// Adds Task to Archive
func (tl *TaskList) addTaskToArhive(task tasks.Task) {
	tl.Archive = append(tl.Archive, task)
}

// Iterates thru ToDoList to find done Task to archive
func (tl *TaskList) ArchiveTask() {
	tdl := []tasks.Task{}
	for _, task := range tl.ToDoList {
		if task.Done {
			tl.addTaskToArhive(task)
		} else {
			tdl = append(tdl, task)
		}
	}
	tl.ToDoList = append(tdl)
}

// Display ToDoList
func (tl *TaskList) ShowLeftToDo() {
	if len(tl.ToDoList) < 1 {
		fmt.Println("\t\tATT GÖRA LISTAN ÄR TOM!")
		return
	}
	fmt.Println("Status\tNr.\tUppgift")

	index := 1
	for _, task := range tl.ToDoList {
		task.ShowTask(index)
		index++
	}
}

// Find Task to mark as done / undone
func (tl *TaskList) FindTaskToMark() {
	fmt.Printf("\n\nVilken uppgift vill du markera / avmarkera? Är det en under uppgift, ange först rubrikens nummer.\n\n")
	input := readInt()

	for i := 0; i < len(tl.ToDoList); i++ {
		task := tl.ToDoList[i]

		if input == i {
			if task.TaskType == "S" {
				tl.ToDoList[i] = task.MarkAsDone(task)

			} else if task.TaskType == "C" {
				fmt.Printf("\n\nVilken underuppgift vill du markera / avmarkera?\n\n")
				subInput := readInt()

				count := len(task.SubTask)
				marked := 0

				for j := 0; j < count; j++ {
					subTask := task.SubTask[j]

					if subTask.Done {
						marked++
					}
					if subInput == j {
						*task.SubTask[j] = subTask.CreateMarkAsDone(*subTask)
						subTask.Done = !subTask.Done
						marked++
					}
				}
				if count == marked {
					tl.ToDoList[i] = task.MarkAsDone(task)
				}
			}
		}
	}
	tl.ShowLeftToDo()
}

// Parse input to an integer or displays error message
func readInt() int {

	var input string
	fmt.Scanln(&input)
	number, err := strconv.Atoi(input)

	for err != nil {
		fmt.Printf("\n\t\tDu skrev inte in en siffra. Försök igen.\n\n")
		fmt.Scanln(&input)
		number, err = strconv.Atoi(input)
	}
	return number - 1
}

// Display Archive
func (tl *TaskList) ShowArchive() {
	if len(tl.Archive) < 1 {
		fmt.Printf("\n\t\tARKIVET ÄR TOMT.\n\n")
		return
	}

	fmt.Println("UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift")
	amount := 0

	for i := 0; i < len(tl.Archive); i++ {
		task := tl.Archive[i]
		task.ShowTask(i + 1)
		amount++

		if task.TaskType == "C" {
			for j := 0; j < len(task.SubTask); j++ {
				sub := task.SubTask[j]
				sub.ShowSubTask(i+1, j+1)
				amount++
			}
		}
	}
	fmt.Printf("\n\t\tWOW! DU HAR UTFÖRT %v UPPGIFTER!\n", amount)
}
