package example

import "fmt"

func Loops() {
	count := 2
	for i := 0; i < count; i++ {
		// Take appropriate action
	}

	keepLooping := true
	for keepLooping {
		keepLooping = false
		// Take appropriate action
	}

	for !keepLooping {
		// Take appropriate action
		keepLooping = true
	}

	arr := []int{1, 2, 3}
	for _, nr := range arr {
		fmt.Println(nr)
	}
}

type example struct {
}

func (e example) ReceiveStringReturnError(parameter string) error {

	returnValue, err := returnStringAndError()
	if err != nil {
		return err
	}

	fmt.Println(parameter + returnValue)
	return nil
}

func returnStringAndError() (returnString string, err error) {
	returnString = " Retur"

	var argument string
	argument = "Argument "

	e := example{}

	if err = e.ReceiveStringReturnError(argument); err != nil {
		return " Error", err
	}
	return returnString, nil
}
