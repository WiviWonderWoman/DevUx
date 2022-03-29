package lists

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/tasks"
)

type TaskList struct {
	ToDoList []tasks.Task
	Archive  []tasks.Task
}

func NewTaskList() (TaskList, error) {
	return TaskList{
		ToDoList: []tasks.Task{},
		Archive:  []tasks.Task{},
	}, nil
}

func (tdl TaskList) AddToDoTask(task tasks.Task) {
	tdl.ToDoList = append(tdl.ToDoList, task)
}

func (tdl TaskList) addToArhive(task tasks.Task) {
	tdl.Archive = append(tdl.Archive, task)
}

func (tdl TaskList) ArchiveTask() {
	for _, task := range tdl.ToDoList {
		if task.Done {
			tdl.addToArhive(task)
		} else {
			tdl.ToDoList = append(tdl.ToDoList, task)
		}
	}
}

func (tdl TaskList) ShowLeftToDo() {
	//ToDo: Console.Clear();
	if len(tdl.ToDoList) < 1 {
		fmt.Println("\n\nATT GÖRA LISTAN ÄR TOM!")
	} else {
		fmt.Println("Status\tNr.\tUppgift")
		index := 0
		for _, task := range tdl.ToDoList {
			//ToDo: Console.ResetColor();
			index++

			if task.Done == false {
				//ToDo: Console.ForegroundColor = ConsoleColor.Red;
			} else if task.Done {
				//ToDo: Console.ForegroundColor = ConsoleColor.Yellow;
			}

			if task.TaskType == "S" {
				task.ShowTask(task, index)
			} else if task.TaskType == "C" {
				task.ShowTask(task, index)
				subTask(task.SubTask, index)
			} else if task.TaskType == "D" {
				task.ShowTask(task, index)
			}
		}
	}
	//ToDo: Console.ResetColor();
}

func subTask(sub []tasks.Task, taskIndex int) {
	index := 0
	for _, subTask := range sub {
		//ToDo: Console.ResetColor();
		index++

		if subTask.Done == false {
			//ToDo: Console.ForegroundColor = ConsoleColor.Red;
		} else if subTask.Done {
			//ToDo: Console.ForegroundColor = ConsoleColor.Yellow;
		}
		subTask.ShowSubTask(subTask, index, taskIndex)
	}
	//TODO: Console.ResetColor();
}

func (tdl TaskList) FindTask() {
	fmt.Println("\nVilken uppgift vill du markera / avmarkera? Är det en under uppgift, ange först rubrikens nummer.")
	//TODO: read input from console
	//TODO: int index = ReadInt() - 1;
	input := 1 - 1

	index := 0
	for _, task := range tdl.ToDoList {
		if input == index {
			if task.TaskType == "C" {
				fmt.Println("\nVilken underuppgift vill du markera / avmarkera?")
				//TODO: read input from console
				//TODO: int subIndex = ReadInt() - 1;
				subInput := 1 - 1
				subIndex := 0
				count := len(task.SubTask)
				marked := 0
				for _, subTask := range task.SubTask {
					if subTask.Done == true {
						marked++
					}

					if subInput == subIndex {
						task.MarkAsDone()
						marked++
					}
					subIndex++
				}

				if count == marked {
					task.MarkAsDone()
				}

			} else {
				task.MarkAsDone()
			}
		}

		index++
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
	} else {
		fmt.Println("UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift")
		//TODO: Console.ForegroundColor = ConsoleColor.Green;
		amount := 0
		index := 0
		subIndex := 0
		for _, task := range tdl.Archive {
			if task.TaskType == "C" {
				index++
				task.ShowTask(task, index)
				// fmt.Println(task.Done, "\t\t", task.Description)
				amount++
				for _, subTask := range task.SubTask {
					subIndex++
					subTask.ShowSubTask(task, index, subIndex)
					// fmt.Println(subTask.Done, "\t\t", subTask.Description)
					amount++
				}
			} else {
				task.ShowTask(task, index)
				// fmt.Println(task.Done, "\t\t", task.Description)
				amount++
			}
		}
		//TODO: Console.ResetColor();
		fmt.Println("\nWOW! DU HAR UTFÖRT {amount} UPPGIFTER!")
	}
}
