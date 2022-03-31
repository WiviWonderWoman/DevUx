package lists

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/tasks"
)

type TaskList struct {
	ToDoList []tasks.TaskWrapper
	Archive  []tasks.TaskWrapper
}

func NewTaskList() TaskList {
	return TaskList{
		ToDoList: []tasks.TaskWrapper{},
		Archive:  []tasks.TaskWrapper{},
	}
}

func (tdl TaskList) AddToDoTask(task tasks.TaskWrapper) {
	tdl.ToDoList = append(tdl.ToDoList, task)
}

func (tdl TaskList) addTaskToArhive(task tasks.TaskWrapper) {
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
	//ToDo: Console.Clear();
	if len(tdl.ToDoList) < 1 {
		fmt.Println("\n\nATT GÖRA LISTAN ÄR TOM!")
		return
	}

	fmt.Println("Status\tNr.\tUppgift")
	index := 0
	for _, task := range tdl.ToDoList {
		//ToDo: Console.ResetColor();
		index++

		if !task.Done {
			//ToDo: Console.ForegroundColor = ConsoleColor.Red;
		} else if task.Done {
			//ToDo: Console.ForegroundColor = ConsoleColor.Yellow;
		}

		if task.TaskType == "S" {
			task.ShowTask(task, index)
		} else if task.TaskType == "C" {
			task.ShowTask(task, index)
			showSubTask(task.SubTask, index)
		} else if task.TaskType == "D" {
			task.ShowTask(task, index)
		}
	}

	//ToDo: Console.ResetColor();
}

func showSubTask(sub []tasks.Task, outer int) {
	inner := 0
	for _, subTask := range sub {
		//ToDo: Console.ResetColor();
		inner++

		if !subTask.Done {
			//ToDo: Console.ForegroundColor = ConsoleColor.Red;
		} else if subTask.Done {
			//ToDo: Console.ForegroundColor = ConsoleColor.Yellow;
		}
		subTask.ShowSubTask(subTask, outer, inner)
	}
	//TODO: Console.ResetColor();
}

func (tdl TaskList) FindTaskToMark() {
	fmt.Println("\nVilken uppgift vill du markera / avmarkera? Är det en under uppgift, ange först rubrikens nummer.")
	input := 1 - 1 //TODO: read input from console int index = ReadInt() - 1;

	for i := 0; i < len(tdl.ToDoList); i++ {
		task := tdl.ToDoList[i]

		if input == i {
			if task.TaskType == "C" {
				fmt.Println("\nVilken underuppgift vill du markera / avmarkera?")
				subInput := 1 - 1 //TODO: read input from console int subIndex = ReadInt() - 1;

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
	//TODO: while (int.TryParse(Console.ReadLine(), out number) == false) { Console.WriteLine("Du skrev inte in ett tal. Försök igen.");}
	return number
}

func (tdl TaskList) ShowArchive() {
	//TODO: Console.Clear();
	if len(tdl.Archive) < 1 {
		fmt.Println("\n\nARKIVET ÄR TOMT.")
		return
	}
	fmt.Println("UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift")
	//TODO: Console.ForegroundColor = ConsoleColor.Green;
	amount := 0
	for i := 0; i < len(tdl.Archive); i++ {
		task := tdl.Archive[i]
		amount++

		task.ShowTask(task, i+1)

		if task.TaskType == "C" {
			showSubTask(task.SubTask, i+1)
		}
	}
	//TODO: Console.ResetColor();
	fmt.Printf("\nWOW! DU HAR UTFÖRT %v UPPGIFTER!", amount)
}
