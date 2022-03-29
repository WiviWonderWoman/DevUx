
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
            daysLeft = 0;
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
        public static void ShowTask(Task task, int index)
        {
            if (task.done == false)
            {
                Console.WriteLine($" - \t{index}\t{task.description}");
            }
            else
            {
                Console.WriteLine($"[{task.done}]\t{index}\t{task.description}");
            }
        }
    }
}
