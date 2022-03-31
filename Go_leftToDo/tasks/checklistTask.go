package tasks

import "fmt"

type ChecklistTask struct {
	TaskType    string
	Description string
	Done        bool
	SubTask     []SimpleTask
}

func NewChecklistTask() ChecklistTask {
	return ChecklistTask{
		TaskType: "D",
		Done:     false,
	}
}

func NewChecklistTaskWrapper() *TaskWrapper {
	return &TaskWrapper{
		ChecklistTask: NewChecklistTask(),
	}
}

// func (c *ChecklistTask) SetTask(desc string, days int, subTasks []SimpleTask) (TaskWrapper, error) {
// 	tw := NewChecklistTaskWrapper()
// 	tw.ChecklistTask.Description = desc
// 	tw.ChecklistTask.SubTask = subTasks
// 	return *tw, nil
// }

func (c *ChecklistTask) Create() (TaskWrapper, error) {
	tw := NewSimpleTaskWrapper()
	fmt.Println("Ange Rubrik-uppgift:")
	input := "" //TO-DO: read input
	tw.SimpleTask.Description = input
	fmt.Println("Ange uppgifter, separerat med enter.\n\n[0] för att slutföra checklistan.")
	// bool done = true;
	// do
	// {
	// 		var specification = Console.ReadLine();
	//      if (specification == "0")
	//      {
	//       	input = false;
	//      }
	//      else
	//      {
	//      	var newTask = new SimpleTask(specification);
	//          newTask.type = "sub";
	//          AddSubTask(newTask);
	//      }
	// } while (input);
	return *tw, nil
}

func (c *ChecklistTask) ShowTask(task TaskWrapper, index int) {
	panic("not implemented") // TODO: Implement
}

func (c *ChecklistTask) ShowSubTask(task TaskWrapper, outer int, inner int) {
	panic("not implemented") // TODO: Implement
}

func (c *ChecklistTask) MarkAsDone() {
	panic("not implemented") // TODO: Implement
}
