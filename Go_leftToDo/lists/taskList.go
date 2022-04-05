package lists

import (
	"fmt"

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

func (tdl TaskList) AddToDoTask(task tasks.Task) {
	tdl.ToDoList = append(tdl.ToDoList, task)
}

func (tdl TaskList) addTaskToArhive(task tasks.Task) {
	tdl.Archive = append(tdl.Archive, task)
}

func (tdl TaskList) ArchiveTask() {
	for i := 0; i < len(tdl.ToDoList); i++ {
		task := tdl.ToDoList[i]

		if task.Done {
			tdl.addTaskToArhive(task)
			tdl.ToDoList = append(tdl.ToDoList[:i], tdl.ToDoList[i+1:]...)
		}
	}
}

func (tdl TaskList) ShowLeftToDo() {
	//TODO: clear console
	if len(tdl.ToDoList) < 1 {
		fmt.Println("\n\nATT GÖRA LISTAN ÄR TOM!")
		return
	}

	fmt.Println("Status\tNr.\tUppgift")
	index := 0
	for _, task := range tdl.ToDoList {
		//TODO: reset console color
		index++

		if !task.Done {
			//TODO: set console color to red
		} else if task.Done {
			//TODO: set console color to yellow
		}

		task.ShowTask(index)

		if task.TaskType == "C" {
			task.ShowSubTask(index)
		}
	}
	//TODO: reset console color
}

// func showSubTask(sub []tasks.Task, outer int) {
// 	inner := 0
// 	for _, subTask := range sub {
// 		//To-Do: Console.ResetColor();
// 		inner++

// 		if !subTask.Done {
// 			//To-Do: Console.ForegroundColor = ConsoleColor.Red;
// 		} else if subTask.Done {
// 			//To-Do: Console.ForegroundColor = ConsoleColor.Yellow;
// 		}
// 		subTask.ShowSubTask(subTask, outer, inner)
// 	}
// 	//TO-DO: Console.ResetColor();
// }

func (tdl TaskList) FindTaskToMark() {
	fmt.Println("\nVilken uppgift vill du markera / avmarkera? Är det en under uppgift, ange först rubrikens nummer.")
	input := 1 - 1 //TO-DO: read input from console int index = ReadInt() - 1;

	for i := 0; i < len(tdl.ToDoList); i++ {
		task := tdl.ToDoList[i]

		if input == i {
			if task.TaskType == "C" {
				fmt.Println("\nVilken underuppgift vill du markera / avmarkera?")
				subInput := 1 - 1 //TO-DO: read input from console int subIndex = ReadInt() - 1;

				count := len(task.SubTask)
				marked := 0

				for j := 0; j < count; j++ {
					subTask := task.SubTask[j]

					if subTask.Done {
						marked++
					}

					if subInput == j {
						task.MarkAsDone()
						marked++
					}
				}

				if count == marked {
					task.MarkAsDone()
				}

			} else {
				task.MarkAsDone()
			}
		}
	}
	tdl.ShowLeftToDo()
}

func readInt() int {
	number := 0
	//TO-DO: while (int.TryParse(Console.ReadLine(), out number) == false) { Console.WriteLine("Du skrev inte in ett tal. Försök igen.");}
	return number
}

func (tdl TaskList) ShowArchive() {
	//TO-DO: Console.Clear();
	if len(tdl.Archive) < 1 {
		fmt.Println("\n\nARKIVET ÄR TOMT.")
		return
	}
	fmt.Println("UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift")
	//TO-DO: Console.ForegroundColor = ConsoleColor.Green;
	amount := 0
	for i := 0; i < len(tdl.Archive); i++ {
		task := tdl.Archive[i]
		amount++

		task.ShowTask(i + 1)

		if task.TaskType == "C" {
			task.ShowSubTask(i + 1)
		}
	}
	//TO-DO: Console.ResetColor();
	fmt.Printf("\nWOW! DU HAR UTFÖRT %v UPPGIFTER!", amount)
}
