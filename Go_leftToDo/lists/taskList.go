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

func AddToDoTask(todoList []tasks.Task, task tasks.Task) []tasks.Task {
	todoList = append(todoList, task)
	return todoList
}

func addTaskToArhive(archive []tasks.Task, task tasks.Task) []tasks.Task {
	archive = append(archive, task)
	return archive
}

func ArchiveTask(list TaskList) TaskList {
	tdl := []tasks.Task{}
	for i := 0; i < len(list.ToDoList); i++ {
		task := list.ToDoList[i]
		if task.Done == true {
			list.Archive = addTaskToArhive(list.Archive, list.ToDoList[i])
		} else if task.Done == false {
			tdl = append(tdl, list.ToDoList[i])
		}
	}
	list.ToDoList = append(tdl)
	return list
}

func ShowLeftToDo(list []tasks.Task) {
	if len(list) < 1 {
		fmt.Println("\t\tATT GÖRA LISTAN ÄR TOM!")
		return
	}
	fmt.Println("Status\t\tNr.\t\tUppgift")
	index := 0
	for _, task := range list {
		index++
		task.ShowTask(index)
		if task.TaskType == "C" {
			for i := 0; i < len(task.SubTask); i++ {
				sub := task.SubTask[i]
				sub.ShowSubTask(index, i+1)
			}
		}
	}
}

func FindTaskToMark(list []tasks.Task) {
	fmt.Printf("\n\nVilken uppgift vill du markera / avmarkera? Är det en under uppgift, ange först rubrikens nummer.\n\n")
	input := readInt()
	for i := 0; i < len(list); i++ {
		task := list[i]
		if input == i {
			if task.TaskType == "S" {
				list[i] = task.MarkAsDone(task)
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
					list[i] = task.MarkAsDone(task)
				}
			}
		}
	}
	ShowLeftToDo(list)
}

func readInt() int {
	var input string
	fmt.Scanln(&input)
	var number int
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("\n\t\tDu skrev inte in en siffra. Försök igen.\n\n")
		fmt.Scanln(&input)
	}
	return number - 1
}

func ShowArchive(list []tasks.Task) {
	if len(list) < 1 {
		fmt.Printf("\n\t\tARKIVET ÄR TOMT.\n\n")
		return
	}
	fmt.Println("UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift")
	amount := 0
	for i := 0; i < len(list); i++ {
		task := list[i]
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
