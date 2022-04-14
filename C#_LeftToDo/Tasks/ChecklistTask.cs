
using System;
using System.Collections.Generic;

namespace LeftToDo.Tasks
{
    public class Checklist : Task // Derives from Task
    {
        // Constructor
        public Checklist(string specification)
        {
            type = "C";
            description = specification;
            subTask = new List<SimpleTask>();
        }

        // Adds input from user as description.
        internal override void Create()
        {
            Console.WriteLine("Ange Rubrik-uppgift:\n");
            description = Console.ReadLine().ToUpper();

            bool input = true;
            Console.WriteLine("Ange uppgifter, separerat med enter.[0] för att slutföra checklistan.");
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

                    subTask.Add(newTask);
                }
            } while (input);
        }
    }
}
