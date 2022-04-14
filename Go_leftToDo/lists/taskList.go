package lists

import (
	"fmt"
	"strconv"

	"github.com/WiviWonderWoman/DevUx/Go/tasks"
)

type TaskList struct {
	ToDoList []tasks.Task
	Archive  []tasks.Task
}

func NewTaskList() TaskList {
	return TaskList{
		ToDoList: []tasks.Task{},
		Archive:  []tasks.Task{},
	}
}

func (tl *TaskList) AddToDoTask(task tasks.Task) {
	tl.ToDoList = append(tl.ToDoList, task)
}

func (tl *TaskList) addTaskToArhive(task tasks.Task) {
	tl.Archive = append(tl.Archive, task)
}

func (tl *TaskList) ArchiveTask() {
	tdl := []tasks.Task{}
	for i := 0; i < len(tl.ToDoList); i++ {
		task := tl.ToDoList[i]
		if task.Done == true {
			tl.addTaskToArhive(tl.ToDoList[i])
		} else if task.Done == false {
			tdl = append(tdl, tl.ToDoList[i])
		}
	}
	tl.ToDoList = append(tdl)
}

func (tl *TaskList) ShowLeftToDo() {
	if len(tl.ToDoList) < 1 {
		fmt.Println("\t\tATT GÖRA LISTAN ÄR TOM!")
		return
	}
	fmt.Println("Status\tNr.\tUppgift")

	for index, task := range tl.ToDoList {
		task.ShowTask(index)
	}
}

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

func (tl *TaskList) ShowArchive() {
	if len(tl.Archive) < 1 {
		fmt.Printf("\n\t\tARKIVET ÄR TOMT.\n\n")
		return
	}
	fmt.Println("UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift")
	amount := 0
	for i := 0; i < len(tl.Archive); i++ {
		task := tl.Archive[i]
		amount++
		task.ShowTask(i + 1)
		if task.TaskType == "C" {
			for j := 0; j < len(task.SubTask); j++ {
				sub := task.SubTask[j]
				sub.ShowSubTask(i+1, j+1)
			}
		}
	}
	fmt.Printf("\n\t\tWOW! DU HAR UTFÖRT %v UPPGIFTER!\n", amount)
}
