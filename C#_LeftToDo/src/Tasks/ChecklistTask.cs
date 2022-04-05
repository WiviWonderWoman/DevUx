
using System;
using System.Collections.Generic;

namespace LeftToDo
{
    // Derives from Task
    public class Checklist : Task
    {
        public Checklist(string specification)
        {
            type = "C";
            description = specification;
            subTask = new List<SimpleTask>();
        }

        // Takes info from user
        internal override void Create()
        {

            Console.WriteLine("\nAnge Rubrik-uppgift:\n");
            description = Console.ReadLine().ToUpper();

            bool input = true;
            Console.WriteLine("\nAnge uppgifter, separerat med enter.\n\n[0] för att slutföra checklistan.\n");
            do
            {
                var specification = Console.ReadLine();

                if (specification == "0")
                {
                    input = false;
                }
                else
                {
                    var newTask = new SimpleTask(specification);
                    newTask.type = "sub";
                    AddSubTask(newTask);
                }

            } while (input);
        }

        // Adding objects to subTask
        public void AddSubTask(SimpleTask task)
        {
            subTask.Add(task);
        }

        //  Displaying checklist "Header" task
        public static void ShowTask(Task task, int outer)
        {
            if (!task.done)
            {
                Console.WriteLine($" - \t{outer}\t{task.description}");
                return;
            }
            Console.WriteLine($"[{task.done}]\t{outer}\t{task.description}");

            for (int i = 0; i < task.subTask.Count; i++)
            {
                var sub = task.subTask[i];
                var inner = i + 1;
                SimpleTask.ShowTask(sub, outer, inner);
            }
        }
    }
}
