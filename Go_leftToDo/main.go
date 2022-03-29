package main

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/lists"
	"github.com/WiviWonderWoman/DevUx/Go/menu"
)

func main() {
	// % warning:   fmt.Println arg list ends with reduntant newline (\n)
	fmt.Println("Välkommen till Din digitala Att Göra Lista!")
	todo, _ := lists.NewTaskList()
	menu, _ := menu.NewMenu()
	menu.Main(todo)
}
