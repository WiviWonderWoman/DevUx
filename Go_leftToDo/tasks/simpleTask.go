package tasks

import "fmt"

type SimpleTask struct {
	TaskType    string
	Description string
	Done        bool
}

func NewSimpleTask() SimpleTask {
	return SimpleTask{
		TaskType: "S",
		Done:     false,
	}
}

func NewSimpleTaskWrapper() *TaskWrapper {
	return &TaskWrapper{
		SimpleTask: NewSimpleTask(),
	}
}

// func (s *SimpleTask) SetTask(desc string, days int, subTasks []SimpleTask) (TaskWrapper, error) {
// 	tw := NewSimpleTaskWrapper()
// 	tw.SimpleTask.Description = desc
// 	return *tw, nil
// }

func (s *SimpleTask) Create() (TaskWrapper, error) {
	tw := NewSimpleTaskWrapper()
	fmt.Println("Ange uppgift:")
	input := "" //TO-DO: read input
	tw.SimpleTask.Description = input
	return *tw, nil
}

// func (s SimpleTask) CreateSinmpleTask() (SimpleTask, error) {
// 	fmt.Println("Ange uppgift:")
// 	input := "" //TO-DO: read input
// 	s.Description = input
// 	return s, nil
// }

func (s *SimpleTask) ShowTask(task TaskWrapper, index int) {
	panic("not implemented") // TODO: Implement
}

// func (s *SimpleTask) ShowTask(task TaskWrapper, index int) {
// 	if !task.Done {
// 		fmt.Println("[ ]\t", index, "\t", task.Description)
// 		return
// 	}
// 	fmt.Println("[X]\t", index, "\t", task.Description)
// }

func (s *SimpleTask) ShowSubTask(task TaskWrapper, outer int, inner int) {
	panic("not implemented") // TODO: Implement
}

func (s *SimpleTask) MarkAsDone() {
	panic("not implemented") // TODO: Implement
}

// func (s SimpleTask) MarkAsDone() {
// 	s.Done = !s.Done
// }

/*_______________________________________________
__________________________________________________*/
// func (tw *TaskWrapper) SetTask(desc string, days int, subTasks []Task) (SimpleTask, error) {
// 	return tw.simpleTask.SetTask(desc, days, subTasks)
// }

// func (tw *TaskWrapper) Create() (SimpleTask, error) {
// 	return tw.simpleTask.CreateSinmpleTask()
// }

// func (tw *TaskWrapper) Create() (SimpleTask, error) {
// 	return tw.simpleTask.Create()
// }

// func (tw *TaskWrapper) ShowTask(task TaskWrapper, index int) {
// 	if !task.Done {
// 		fmt.Println("[ ]\t", index, "\t", task.Description)
// 		return
// 	}
// 	fmt.Println("[X]\t", index, "\t", task.Description)
// }

// func (tw *TaskWrapper) MarkAsDone() {
// 	tw.simpleTask.MarkAsDone()
// }
