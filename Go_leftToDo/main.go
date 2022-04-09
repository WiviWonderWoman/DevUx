// Package main is the starting point of the program
package main

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/lists"
	"github.com/WiviWonderWoman/DevUx/Go/menu"
)

func main() {
	fmt.Printf("Välkommen till Din digitala Att Göra Lista!\n")
	list := lists.NewTaskList()
	menu.ShowMainMenu(list)
}
