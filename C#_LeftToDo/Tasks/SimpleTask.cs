using System;

namespace LeftToDo.Tasks
{
    public class SimpleTask : Task // Derives from Task
    {
        // Constructor
        public SimpleTask(string specification)
        {
            type = "S";
            description = specification;
        }

        // Adds input from user as description.
        internal override void Create()
        {
            Console.WriteLine("Ange uppgift:");
            description = Console.ReadLine();
        }

        // Overload of ShowTask, displays SimpleTask in ChecklistTasks subTask
        internal static void ShowTask(Task task, int outer, int inner)
        {
            if (!task.done)
            {
                Console.WriteLine($"[ ]\t{outer} - {inner}\t{task.description}");
                return;
            }
            Console.WriteLine($"[X]\t{outer} - {inner}\t{task.description}");
        }


    }
}
