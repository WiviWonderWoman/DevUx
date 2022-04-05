package menu

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/lists"
	"github.com/WiviWonderWoman/DevUx/Go/tasks"
)

func ShowMainMenu(list lists.TaskList) {
	//TODO: clear console
	showIntro()
	fmt.Println("\nHUVUDMENY\n[1] Visa Att-göra uppgifter\n[2] Visa Arkiverade uppgifter\n[0] Avsluta")
	//TODO: read input from console
	input := "1"
	switch input {
	case "1":
		list.ShowLeftToDo()
		showToDoMenu(list)
	case "2":
		list.ShowArchive()
		showArchiveMenu(list)
	case "3":
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
	fmt.Println(
		"\nUNDERMENY: Vad vill du göra?\n" +
			"[1] Lägg till uppgift\n" +
			"[2] Markera / avmarkera uppgift\n" +
			"[3] Arkivera utförda uppgifter\n" +
			"[0] HUVUDMENY",
	)
	//TODO: read input from console
	input := ""
	switch input {
	case "1":
		showTaskMenu(list)
	case "2":
		list.FindTaskToMark()
		showToDoMenu(list)
	case "3":
		list.ArchiveTask()
		list.ShowLeftToDo()
		showToDoMenu(list)
	case "0":
		//TODO: clear console
		ShowMainMenu(list)
	default:
		ShowErrorMsg()
	}
}

func showTaskMenu(list lists.TaskList) {
	list.ShowLeftToDo()
	showIntro()

	fmt.Println(
		"\nUPPGIFTSMENY. Välj typ av uppgift:\n" +
			"[1] Enkel uppgift\n" +
			"[2] Checklista\n" +
			"[0] HUVUDMENY",
	)
	//TODO: read input from console
	input := ""
	switch input {
	case "1":
		simple := tasks.NewSimpleTask("")
		simple.Create()
		list.AddToDoTask(simple.Task)
		list.ShowLeftToDo()
		showToDoMenu(list)
	case "2":
		checklist := tasks.NewChecklistTask("")
		checklist.Create()
		list.AddToDoTask(checklist.Task)
		list.ShowLeftToDo()
		showToDoMenu(list)
	case "0":
		//TODO: clear console
		ShowMainMenu(list)
	default:
		ShowErrorMsg()
	}
}

func showArchiveMenu(list lists.TaskList) {
	fmt.Println("\n[0] HUVUDMENY")
	//TODO: read input from console
	input := ""
	if input == "0" {
		//TODO: clear console
		ShowMainMenu(list)
	}
}

func showFarewell() {
	fmt.Println("Tack för besöket och välkommen åter!")
}

func ShowErrorMsg() {
	fmt.Println("\nOgiltligt val. Vargod försök igen!")
}
