package menu

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/lists"
)

type Menu struct {
}

func NewMenu() (Menu, error) {
	return Menu{}, nil
}

func intro() {
	fmt.Println("Välj i menyn nedan, ange siffran inom [] och tryck enter.")
}

func (m Menu) Main(list lists.TodoList) {
	fmt.Println("MAIN menyn", list)
	//TODO: clear console
	intro()
	fmt.Println("\nHUVUDMENY\n[1] Visa Att-göra uppgifter\n[2] Visa Arkiverade uppgifter\n[0] Avsluta")

	//TODO: read input from console
	input := 0
	switch input {
	case 1:
		//TODO: list.ShowLeftToDo(list.ToDo);
		m.todo(list)
	case 2:
		//TODO: list.ShowArchive(list.Archive);
		m.archive(list)
	case 3:
		m.farewell()
	default:
		m.errorMsg()
		m.Main(list)
	}
}

func (m Menu) todo(list lists.TodoList) {
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
		//TODO: list.FindTaskToMark(list);
		m.todo(list)
	case 3:
		//TODO: list.ArchiveTask();
		//TODO: list.ShowLeftToDo(list.ToDo);
		m.todo(list)
	case 0:
		//TODO: Console.Clear();
		m.Main(list)
	}
}

func (m Menu) task(list lists.TodoList) {
	//TODO: list.ShowLeftToDo(list.ToDo);
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
		//TODO: var simpelTask = new SimpleTask("");
		//TODO: simpelTask.Create();
		//TODO: list.AddToDoList(simpelTask);
		//TODO: list.ShowLeftToDo(list.ToDo);
		m.todo(list)
	case 2:
		//TODO: var deadline = new Deadline("", 0);
		//TODO: deadline.Create();
		//TODO: list.AddToDoList(deadline);
		//TODO: list.ShowLeftToDo(list.ToDo);
		m.todo(list)
	case 3:
		//TODO: var checklist = new Checklist("");
		//TODO: checklist.Create();
		//TODO: list.AddToDoList(checklist);
		//TODO: list.ShowLeftToDo(list.ToDo);
		m.todo(list)
	case 0:
		//TODO: Console.Clear();
		m.Main(list)
	default:
		m.errorMsg()

	}
}

func (m Menu) archive(list lists.TodoList) {
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
