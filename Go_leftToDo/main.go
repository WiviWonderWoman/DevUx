// Package main is the startingpoint of the program
package main

import (
	"fmt"

	"github.com/WiviWonderWoman/DevUx/Go/lists"
	"github.com/WiviWonderWoman/DevUx/Go/menu"
)

func main() {
	fmt.Printf("Välkommen till Din digitala Att Göra Lista!\n")
	todo := lists.NewTaskList()
	menu.ShowMainMenu(todo)
}
