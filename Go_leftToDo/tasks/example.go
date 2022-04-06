package tasks

import "fmt"

// lowercase initial letter = private struct not exported to other packages
type task struct {
	taskType    string        // describes task type user choose
	description string        // what action / ToDo task user inputs
	done        bool          // indicates if task is done
	subTask     []*SimpleTask // if task is of type checklist, a slice with subtasks
}

// capital initial letter = public struct exported to other packages
type TAsk struct {
	TaskType    string        // describes task type user choose
	Description string        // what action / ToDo task user inputs
	Done        bool          // indicates if task is done
	SubTask     []*SimpleTask // if task is of type checklist, a slice with subtasks
}

//[IDENTIFIER+ACCESS] [RETURNS]
func readInt() (int, error) {
	number := 1
	t := task{}
	// 			 [METHOD CALL]  [ARGUMENT] check for and handle error
	if err := t.ReceiveIntNoReturn(number); err != nil {
		return 0, err
	}
	return number, nil
}

// [METHOD]	  [IDENTIFIER+ACCESS] [PARAMETER]	[RETURN]
func (t task) ReceiveIntNoReturn(number int) error {
	// 			[FUNCTION CALL]
	integer, err := readInt()
	// check for and handle error
	if err != nil {
		return err
	}
	fmt.Println(number + integer)
	return nil
}
