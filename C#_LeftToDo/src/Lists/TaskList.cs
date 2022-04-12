using System;
using System.Collections.Generic;
using LeftToDo.Tasks;

namespace LeftToDo.Lists
{
    /* ToDolist a class that handles all lists and their methods */
    public class TaskList
    {
        public List<Task> ToDoList // Tasks ToDo
        {
            get;
            private set;
        }
        public List<Task> Archive // Archived done Tasks
        {
            get;
            private set;
        }

        // Constructor
        public TaskList()
        {
            ToDoList = new List<Task>();
            Archive = new List<Task>();
        }

        // Adds new Task to ToDoList
        public void AddToDoTask(Task task)
        {
            ToDoList.Add(task);
        }

        // Adds Task to Archive
        private void AddTaskToArhive(Task task)
        {
            Archive.Add(task);
        }

        // Iterates thru ToDoList to find done Task to archive
        public void ArchiveTask()
        {
            for (int i = 0; i < ToDoList.Count; i++)
            {
                var task = ToDoList[i];
                if (task.done)
                {
                    AddTaskToArhive(task);
                    ToDoList.Remove(task);
                    i--;
                }
            }
        }

        /// <summary>Display ToDoList</summary>  
        internal void ShowLeftToDo(List<Task> ToDoList)
        {
            if (ToDoList.Count < 1)
            {
                Console.WriteLine($"\t\tATT GÖRA LISTAN ÄR TOM!");
                return;
            }
            Console.WriteLine($"Status\tNr.\tUppgift\n");

            int index = 1;
            foreach (var task in ToDoList)
            {
                Task.ShowTask(task, index);
                index++;
            }
        }

        // Find Task to mark as done / undone
        internal void FindTaskToMark(TaskList list)
        {
            Console.WriteLine("\nVilken uppgift vill du markera / avmarkera? Är det en under uppgift, ange först rubrikens nummer.\n");
            int index = ReadInt() - 1;

            for (int i = 0; i < list.ToDoList.Count; i++)
            {
                var task = list.ToDoList[i];
                if (i == index)
                {
                    if (task.type == "C")
                    {
                        Console.WriteLine("\nVilken underuppgift vill du markera / avmarkera?\n");
                        int subIndex = ReadInt() - 1;

                        var count = task.subTask.Count;
                        var marked = 0;

                        for (int j = 0; j < count; j++)
                        {
                            var subTask = task.subTask[j];
                            if (subTask.done)
                            {
                                marked++;
                            }
                            if (j == subIndex)
                            {
                                subTask.MarkAsDone();
                                marked++;
                            }
                        }
                        if (count == marked)
                        {
                            task.MarkAsDone();
                        }
                    }
                    else
                    {
                        task.MarkAsDone();
                    }
                }
            }
            list.ShowLeftToDo(list.ToDoList);
        }

        // Parse input to an integer or displays error message
        private static int ReadInt()
        {
            int number;
            while (int.TryParse(Console.ReadLine(), out number) == false)
            {
                Console.WriteLine("\t\tDu skrev inte in en siffra. Försök igen.\n\n");
            }
            return number;
        }

        // Display Archive
        internal void ShowArchive(List<Task> Archive)
        {
            if (Archive.Count < 1)
            {
                Console.WriteLine($"\t\tARKIVET ÄR TOMT.\n\n");
                return;
            }

            Console.WriteLine($"UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift\n");
            int amount = 0;

            for (int i = 0; i < Archive.Count; i++)
            {
                var task = Archive[i];
                amount++;
                if (task.type == "C")
                {
                    Checklist.ShowTask(task, i + 1);
                    for (int j = 0; j < task.subTask.Count; j++)
                    {
                        var sub = task.subTask[j];
                        SimpleTask.ShowTask(sub, i + 1, j + 1);
                        amount++;
                    }
                }
                else if (task.type == "S")
                {
                    SimpleTask.ShowTask(task, i + 1);
                }
            }
            Console.WriteLine($"\t\tWOW! DU HAR UTFÖRT {amount} UPPGIFTER!\n\n");
        }
    }
}
