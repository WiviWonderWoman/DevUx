package menu

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/lists"
	"github.com/WiviWonderWoman/DevUx/Go/tasks"
)

type Menu struct {
}

func NewMenu() (Menu, error) {
	return Menu{}, nil
}

func intro() {
	fmt.Println("Välj i menyn nedan, ange siffran inom [] och tryck enter.")
}

func (m Menu) Main(list lists.TaskList) {
	fmt.Println("MAIN menyn", list)
	//TODO: clear console
	intro()
	fmt.Println("\nHUVUDMENY\n[1] Visa Att-göra uppgifter\n[2] Visa Arkiverade uppgifter\n[0] Avsluta")

	//TODO: read input from console
	input := 0
	switch input {
	case 1:
		list.ShowLeftToDo()
		m.todo(list)
	case 2:
		list.ShowArchive()
		m.archive(list)
	case 3:
		m.farewell()
	default:
		m.errorMsg()
		m.Main(list)
	}
}

func (m Menu) todo(list lists.TaskList) {
	intro()
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
		m.task(list)
	case 2:
		list.FindTask()
		m.todo(list)
	case 3:
		list.ArchiveTask()
		list.ShowLeftToDo()
		m.todo(list)
	case 0:
		//TODO: Console.Clear();
		m.Main(list)
	}
}

func (m Menu) task(list lists.TaskList) {
	list.ShowLeftToDo()
	intro()

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

		simple := tasks.Task{}
		simple.Create("S")
		list.AddToDoTask(simple)
		list.ShowLeftToDo()
		m.todo(list)
	case 2:
		deadline := tasks.Task{}
		deadline.Create("D")
		list.AddToDoTask(deadline)
		list.ShowLeftToDo()
		m.todo(list)
	case 3:
		checklist := tasks.Task{}
		checklist.Create("C")
		list.AddToDoTask(checklist)
		list.ShowLeftToDo()
		m.todo(list)
	case 0:
		//TODO: Console.Clear();
		m.Main(list)
	default:
		m.errorMsg()

	}
}

func (m Menu) archive(list lists.TaskList) {
	fmt.Println("\n[0] HUVUDMENY")
	//TODO: read input from console
	input := 0
	if input == 0 {
		//TODO: Console.Clear();
		m.Main(list)
	}
}

func (m Menu) farewell() {
	fmt.Println("Tack för besöket och välkommen åter!")
}

func (m Menu) errorMsg() {
	fmt.Println("\nOgiltligt val. Vargod försök igen!")
}
