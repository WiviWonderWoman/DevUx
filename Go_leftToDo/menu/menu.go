package menu

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/lists"
	"github.com/WiviWonderWoman/DevUx/Go/tasks"
)

func ShowMainMenu(list lists.TaskList) {
	fmt.Println("MAIN menyn", list)
	//TODO: clear console
	showIntro()
	fmt.Println("\nHUVUDMENY\n[1] Visa Att-göra uppgifter\n[2] Visa Arkiverade uppgifter\n[0] Avsluta")
	//TODO: read input from console
	input := 0
	switch input {
	case 1:
		list.ShowLeftToDo()
		showToDoMenu(list)
	case 2:
		list.ShowArchive()
		showArchiveMenu(list)
	case 3:
		showFarewell()
	default:
		ShowErrorMsg()
		ShowMainMenu(list)
	}
}

func showIntro() {
	fmt.Println("Välj i menyn nedan, ange siffran inom [] och tryck enter.")
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
	input := 0
	switch input {
	case 1:
		showTaskMenu(list)
	case 2:
		list.FindTaskToMark()
		showToDoMenu(list)
	case 3:
		list.ArchiveTask()
		list.ShowLeftToDo()
		showToDoMenu(list)
	case 0:
		//TODO: Console.Clear();
		ShowMainMenu(list)
	}
}

func showTaskMenu(list lists.TaskList) {
	list.ShowLeftToDo()
	showIntro()

	fmt.Println(
		"\nUPPGIFTSMENY. Välj typ av uppgift:\n" +
			"[1] Enkel uppgift\n" +
			"[2] Deadline\n" +
			"[3] Checklista\n" +
			"[0] HUVUDMENY",
	)
	//TODO: read input from console
	input := 0
	switch input {
	case 1:
		st := tasks.NewSimpleTask()
		simpleWrapper := tasks.NewTaskWrapper(&st)
		st.Create()
		list.AddToDoTask(*simpleWrapper)
		list.ShowLeftToDo()
		showToDoMenu(list)
	case 2:
		deadline := tasks.Task{}
		deadline.Create("D")
		list.AddToDoTask(deadline)
		list.ShowLeftToDo()
		showToDoMenu(list)
	case 3:
		checklist := tasks.Task{}
		checklist.Create("C")
		list.AddToDoTask(checklist)
		list.ShowLeftToDo()
		showToDoMenu(list)
	case 0:
		//TODO: Console.Clear();
		ShowMainMenu(list)
	default:
		ShowErrorMsg()
	}
}

func showArchiveMenu(list lists.TaskList) {
	fmt.Println("\n[0] HUVUDMENY")
	//TODO: read input from console
	input := 0
	if input == 0 {
		//TODO: Console.Clear();
		ShowMainMenu(list)
	}
}

func showFarewell() {
	fmt.Println("Tack för besöket och välkommen åter!")
}

func ShowErrorMsg() {
	fmt.Println("\nOgiltligt val. Vargod försök igen!")
}
