package tasks

import "fmt"

//[IDENTIFIER+ACCESS] [RETURNS]
func readInt() (returnNr int, err error) {
	returnNr = 1
	e := example{}
	// [INSTANCE][METHOD CALL]  [ARGUMENT]   // check for and handle error
	if err = e.ReceiveIntReturnError(returnNr); err != nil {
		return 0, err
	}
	return returnNr, nil
}

//[INSTANCE] [IDENTIFIER+ACCESS]   [PARAMETER] [RETURN]
func (t example) ReceiveIntReturnError(number int) error {
	// 			[FUNCTION CALL]
	integer, err := readInt()
	// check for and handle error
	if err != nil {
		return err
	}
	fmt.Println(number + integer)
	return nil
}

/*





 */
// lowercase initial letter = private struct not exported to other packages
type example struct {
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

/*

//TODO: DELETE

*/
