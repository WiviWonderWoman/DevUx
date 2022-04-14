// Package menu displays and handles out- / input
package menu

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/lists"
	"github.com/WiviWonderWoman/DevUx/Go/tasks"
)

// Displays main menu
func ShowMainMenu(list lists.TaskList) {
	fmt.Printf("\n\tHUVUDMENY\n[1] Visa Att-göra uppgifter\n[2] Visa Arkiverade uppgifter\n[0] Avsluta\n")
	showIntro()
	var input string
	fmt.Scanln(&input)
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
		showErrorMsg()
		ShowMainMenu(list)
	}
}

// Displays instructions to user
func showIntro() {
	fmt.Printf("\nVälj siffra inom [] följt av enter.\n")
}

// Displays submenu
func showToDoMenu(list lists.TaskList) {
	fmt.Printf(
		"\n\tUNDERMENY: Vad vill du göra?\n" +
			"[1] Lägg till uppgift\n" +
			"[2] Markera / avmarkera uppgift\n" +
			"[3] Arkivera utförda uppgifter\n" +
			"[0] Tillbaka till HUVUDMENY\n",
	)
	showIntro()
	var input string
	fmt.Scanln(&input)
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
		ShowMainMenu(list)
	default:
		showErrorMsg()
		showToDoMenu(list)
	}
}

// Display task menu
func showTaskMenu(list lists.TaskList) {
	list.ShowLeftToDo()
	fmt.Printf(
		"\n\tUPPGIFTSMENY. Välj typ av uppgift:\n" +
			"[1] Enkel uppgift\n" +
			"[2] Checklista\n" +
			"[0] Tillbaka till HUVUDMENY\n",
	)
	showIntro()
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		simple := tasks.SimpleTask{}.Create()
		list.AddToDoTask(simple.Task)
		list.ShowLeftToDo()
		showToDoMenu(list)
	case "2":
		checklist := tasks.ChecklistTask{}.Create()
		list.AddToDoTask(checklist.Task)
		list.ShowLeftToDo()
		showToDoMenu(list)
	case "0":
		ShowMainMenu(list)
	default:
		showErrorMsg()
		showTaskMenu(list)
	}
}

// Displays option for user to return to main menu
func showArchiveMenu(list lists.TaskList) {
	fmt.Printf("\n[0] Tillbaka till HUVUDMENY\n")
	var input string
	fmt.Scanln(&input)
	if input == "0" {
		ShowMainMenu(list)
	} else {
		showErrorMsg()
		showArchiveMenu(list)
	}
}

// Displays farewell message
func showFarewell() {
	fmt.Printf("\n\t\tTack för besöket och välkommen åter!\n\n")
}

// Displays error message
func showErrorMsg() {
	fmt.Printf("\n\t\tOgiltligt val. Vargod försök igen!\n\n")
}
