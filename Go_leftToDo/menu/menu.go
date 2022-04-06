package menu

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/lists"
	"github.com/WiviWonderWoman/DevUx/Go/tasks"
)

func ShowMainMenu(list lists.TaskList) {
	//TODO: clear console
	showIntro()
	fmt.Printf("\nHUVUDMENY\n[1] Visa Att-göra uppgifter\n[2] Visa Arkiverade uppgifter\n[0] Avsluta\n")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		lists.ShowLeftToDo(list.ToDoList)
		showToDoMenu(list)
	case "2":
		lists.ShowArchive(list.Archive)
		showArchiveMenu(list)
	case "0":
		showFarewell()
	default:
		ShowErrorMsg()
		ShowMainMenu(list)
	}
}

func showIntro() {
	fmt.Println("\nVälj i menyn nedan, ange siffran inom [] och tryck enter.")
}

func showToDoMenu(list lists.TaskList) {
	showIntro()
	fmt.Printf(
		"\nUNDERMENY: Vad vill du göra?\n" +
			"[1] Lägg till uppgift\n" +
			"[2] Markera / avmarkera uppgift\n" +
			"[3] Arkivera utförda uppgifter\n" +
			"[0] HUVUDMENY\n",
	)
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		showTaskMenu(list)
	case "2":
		lists.FindTaskToMark(list.ToDoList)
		showToDoMenu(list)
	case "3":
		list = lists.ArchiveTask(list)
		lists.ShowLeftToDo(list.ToDoList)
		showToDoMenu(list)
	case "0":
		//TODO: clear console
		ShowMainMenu(list)
	default:
		ShowErrorMsg()
	}
}

func showTaskMenu(list lists.TaskList) {
	lists.ShowLeftToDo(list.ToDoList)
	showIntro()
	fmt.Println(
		"\nUPPGIFTSMENY. Välj typ av uppgift:\n" +
			"[1] Enkel uppgift\n" +
			"[2] Checklista\n" +
			"[0] HUVUDMENY",
	)
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		simple := tasks.SimpleTask{}.Create()
		list.ToDoList = lists.AddToDoTask(list.ToDoList, simple.Task)
		lists.ShowLeftToDo(list.ToDoList)
		showToDoMenu(list)
	case "2":
		checklist := tasks.ChecklistTask{}.Create()
		list.ToDoList = lists.AddToDoTask(list.ToDoList, checklist.Task)
		lists.ShowLeftToDo(list.ToDoList)
		showToDoMenu(list)
	case "0":
		//TODO: clear console
		ShowMainMenu(list)
	default:
		ShowErrorMsg()
	}
}

func showArchiveMenu(list lists.TaskList) {
	fmt.Printf("\n[0] HUVUDMENY\n")
	var input string
	fmt.Scanln(&input)
	if input == "0" {
		//TODO: clear console
		ShowMainMenu(list)
	}
}

func showFarewell() {
	fmt.Printf("\nTack för besöket och välkommen åter!\n")
}

func ShowErrorMsg() {
	fmt.Printf("\nOgiltligt val. Vargod försök igen!\n")
}
