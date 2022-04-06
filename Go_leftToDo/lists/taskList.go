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
	//TODO: clear console
	if len(list) < 1 {
		fmt.Println("\n\nATT GÖRA LISTAN ÄR TOM!")
		return
	}
	fmt.Println("Status\tNr.\tUppgift")
	index := 0
	for _, task := range list {
		//TODO: reset console color
		index++
		if !task.Done {
			//TODO: set console color to red
		} else if task.Done {
			//TODO: set console color to yellow
		}
		task.ShowTask(index)
		if task.TaskType == "C" {
			for i := 0; i < len(task.SubTask); i++ {
				sub := task.SubTask[i]
				sub.ShowSubTask(index, i+1)
			}
		}
	}
	//TODO: reset console color
}

func FindTaskToMark(list []tasks.Task) {
	fmt.Println("\nVilken uppgift vill du markera / avmarkera? Är det en under uppgift, ange först rubrikens nummer.")
	input := readInt()
	for i := 0; i < len(list); i++ {
		task := list[i]
		if input == i {
			fmt.Println("INPUT: ", input, "INDEX: ", i)
			if task.TaskType == "S" {
				list[i] = task.MarkAsDone(task)
			} else if task.TaskType == "C" {
				fmt.Println("\nVilken underuppgift vill du markera / avmarkera?")
				subInput := readInt()
				count := len(task.SubTask)
				marked := 0
				for j := 0; j < count; j++ {
					subTask := task.SubTask[j]
					if subTask.Done {
						marked++
					}
					if subInput == j {
						fmt.Println("SUBINPUT: ", input, "SUBINDEX: ", i)
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
		fmt.Println("Du skrev inte in ett tal. Försök igen.")
		fmt.Scanln(&input)
	}
	return number - 1
}

func ShowArchive(list []tasks.Task) {
	//TODO: clear console
	if len(list) < 1 {
		fmt.Println("\n\nARKIVET ÄR TOMT.")
		return
	}
	fmt.Println("UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift")
	//TODO: set console color to green
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
	//TODO: reset console color
	fmt.Printf("\nWOW! DU HAR UTFÖRT %v UPPGIFTER!", amount)
}
