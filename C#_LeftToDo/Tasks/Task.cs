using System;
using System.Collections.Generic;

namespace LeftToDo.Tasks
{
    // Abstract super class, all types of Task derives from Task
    public abstract class Task
    {
        // describes task type user choose
        public string type { get; set; }

        // what action/ToDo-task user inputs
        public string description { get; set; }
        // indicates if task is done
        public bool done { get; set; }

        // if task is of type checklist, a slice with subtasks
        public List<SimpleTask> subTask;

        // Marks task as Done
        public void MarkAsDone()
        {
            done = !done;
        }

        // Takes info from user as description, overrides in each subclass
        internal abstract void Create();

        //  Displays any Task
        public static void ShowTask(Task task, int index)
        {
            if ((!task.done) && (task.type == "C"))
            {
                Console.WriteLine($" - \t{index}\t{task.description}");

                for (int i = 0; i < task.subTask.Count; i++)
                {
                    var sub = task.subTask[i];
                    var inner = i + 1;
                    SimpleTask.ShowTask(sub, index, inner);
                }
                return;
            }
            else if ((!task.done) && (task.type == "S"))
            {
                Console.WriteLine($"[ ]\t{index}\t{task.description}");
                return;
            }
            Console.WriteLine($"[X]\t{index}\t{task.description}");
        }
    }
}