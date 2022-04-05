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

func AddToDoTask(list []tasks.Task, task tasks.Task) []tasks.Task {
	list = append(list, task)
	return list
}

func addTaskToArhive(list []tasks.Task, task tasks.Task) {
	list = append(list, task)
}

func ArchiveTask(list TaskList) {
	for i := 0; i < len(list.ToDoList); i++ {
		task := list.ToDoList[i]

		if task.Done {
			addTaskToArhive(list.Archive, task)
			list.ToDoList = append(list.ToDoList[:i], list.ToDoList[i+1:]...)
		}
	}
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
	input := readInt() //TODO: read input from console and parse to int readInt()

	for i := 0; i < len(list); i++ {
		task := list[i]

		if input == i {
			fmt.Println("INPUT: ", input, "INDEX: ", i)
			if task.TaskType == "S" {
				task.MarkAsDone()
				// task.Done = !task.Done
			} else if task.TaskType == "C" {
				fmt.Println("\nVilken underuppgift vill du markera / avmarkera?")
				subInput := readInt() //TODO: read input from console and parse to int readInt()

				count := len(task.SubTask)
				marked := 0

				for j := 0; j < count; j++ {
					subTask := *task.SubTask[j]

					if subTask.Done {
						marked++
					}

					if subInput == j {
						fmt.Println("SUBINPUT: ", input, "SUBINDEX: ", i)
						subTask.MarkAsDone()
						subTask.Done = !subTask.Done
						marked++
					}
				}

				if count == marked {
					task.MarkAsDone()
					// task.Done = !task.Done
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
