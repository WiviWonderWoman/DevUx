package example

import "fmt"

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
