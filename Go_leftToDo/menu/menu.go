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
	//TODO: read input from console
	var input string
	fmt.Scanln(&input)
	// fmt.Printf("Du tryckte: %s i HUVUDMENYN", input)
	switch input {
	case "1":
		list.ShowLeftToDo()
		showToDoMenu(list)
	case "2":
		list.ShowArchive()
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
	// fmt.Printf("Du tryckte: %s i UNDERMENYN", input)
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
	var input string
	fmt.Scanln(&input)
	// fmt.Printf("Du tryckte: %s i UPPGIFTSMENYN", input)
	switch input {
	case "1":
		s := tasks.SimpleTask{}
		simple := s.Create()
		list.AddToDoTask(simple.Task)
		fmt.Println("LISTAN: ", list)
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
	fmt.Printf("\n[0] HUVUDMENY\n")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("Du tryckte: %s i ARKIVET", input)
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
