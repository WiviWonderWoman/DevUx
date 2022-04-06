using System;

namespace LeftToDo
{
    public class SimpleTask : Task // Derives from Task
    {
        public SimpleTask(string specification)
        {
            type = "S";
            description = specification;
        }
        // Displaying simple Task
        internal static void ShowTask(Task task, int index)
        {
            if (!task.done)
            {
                Console.WriteLine($"[ ]\t{index}\t{task.description}");
                return;
            }
            Console.WriteLine($"[X]\t{index}\t{task.description}");
        }
        // Overload, Displaying simple Task in Checklist
        internal static void ShowTask(Task task, int outer, int inner)
        {
            if (!task.done)
            {
                Console.WriteLine($"\t[ ]\t{outer} - {inner}\t{task.description}");
                return;
            }
            Console.WriteLine($"\t[X]\t{outer} - {inner}\t{task.description}");
        }
        // Takes info from user
        internal override void Create()
        {
            Console.WriteLine("Ange uppgift:");
            description = Console.ReadLine();
        }
    }
}
