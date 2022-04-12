package example

import "fmt"

func variables() {

	var declare string
	declare = "assign"

	var declareAndAssign = "Declare and " + declare

	x, y, z := "declare ", "assign ", "multiple variables"
	fmt.Println(x, y, z)

	shortDeclare := declareAndAssign + " local variable"
	fmt.Println(shortDeclare)
}

func Loops() {
	count := 2
	arr := []int{1, 2, 3}
	keepLooping := true

	for i := 0; i < count; i++ {
		// Take appropriate action
	}

	for _, nr := range arr {
		fmt.Println(nr)
	}

	for keepLooping {
		keepLooping = false
		// Take appropriate action
	}
}
