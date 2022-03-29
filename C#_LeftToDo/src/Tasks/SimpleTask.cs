using System;

namespace LeftToDo
{
    // Derives from Task
    public class SimpleTask : Task
    {
        public SimpleTask(string specification)
        {
            type = "S";
            description = specification;
            daysLeft = 0;
        }

        // Displaying simple Task
        internal static void ShowTask(Task task, int index)
        {
            if (task.done == false)
            {
                Console.WriteLine($"[ ]\t{index}\t{task.description}");
            }
            else
            {
                Console.WriteLine($"[X]\t{index}\t{task.description}");
            }
        }

        // Overload, Displaying simple Task in Checklist
        internal static void ShowTask(Task task, int outer, int inner)
        {
            if (task.done == false)
            {
                Console.WriteLine($"\t[ ]\t{outer} - {inner}\t{task.description}");

            }
            else
            {
                Console.WriteLine($"\t[X]\t{outer} - {inner}\t{task.description}");
            }

        }

        // Takes info from user
        internal override void Create()
        {
            Console.WriteLine("Ange uppgift:");
            description = Console.ReadLine();
        }

    }
}
