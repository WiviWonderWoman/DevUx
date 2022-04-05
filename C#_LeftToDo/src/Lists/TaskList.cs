using System;
using System.Collections.Generic;

namespace LeftToDo
{
    /* ToDolist handles all lists and their methods */
    public class TaskList
    {
        public List<Task> ToDoList
        {
            get;
            private set;
        }
        public List<Task> Archive
        {
            get;
            private set;
        }


        public TaskList()
        {
            ToDoList = new List<Task>();
            Archive = new List<Task>();
        }

        // Adds new Task to ToDo
        public void AddToDoTask(Task task)
        {
            ToDoList.Add(task);
        }

        // Adds Task that should be archived to Arc
        private void AddTaskToArhive(Task task)
        {
            Archive.Add(task);
        }

        // Iterates thru ToDo to find done Task to archive
        public void ArchiveTask()
        {
            foreach (var task in ToDoList)
            {
                if (task.done)
                {
                    AddTaskToArhive(task);
                    ToDoList.Remove(task);
                }
            }
        }

        // Display List of Task 
        internal void ShowLeftToDo(List<Task> ToDo)
        {
            // FIXME: Console.Clear();
            if (ToDo.Count < 1)
            {
                Console.WriteLine($"\n\nATT GÖRA LISTAN ÄR TOM!\n\n");
                return;
            }
            Console.WriteLine($"Status\tNr.\tUppgift\n");

            for (int i = 0; i < ToDo.Count; i++)
            {
                // FIXME: Console.ResetColor();
                var item = ToDo[i];
                var index = i + 1;

                if (!item.done)
                {
                    // FIXME: Console.ForegroundColor = ConsoleColor.Red;
                }
                else if (item.done)
                {
                    // FIXME: Console.ForegroundColor = ConsoleColor.Yellow;
                }

                if (item.type == "S")
                {
                    SimpleTask.ShowTask(item, index);
                }
                else if (item.type == "C")
                {
                    Checklist.ShowTask(item, index);
                    // ShowSubTask(item.subTask, index);
                }
            }
            // FIXME: Console.ResetColor();
        }

        // Find task to mark as done / undone
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
                            if (subTask.done == true)
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
        private static int ReadInt()
        {
            int number;
            while (int.TryParse(Console.ReadLine(), out number) == false)
            {
                Console.WriteLine("Du skrev inte in ett tal. Försök igen.");
            }
            return number;
        }

        // Display archived Task
        internal void ShowArchive(List<Task> Arc)
        {
            // FIXME: Console.Clear();
            if (Arc.Count < 1)
            {
                Console.WriteLine($"\n\nARKIVET ÄR TOMT.\n\n");
                return;
            }
            Console.WriteLine($"UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift\n");
            // FIXME: Console.ForegroundColor = ConsoleColor.Green;

            int amount = 0;
            for (int i = 0; i < Arc.Count; i++)
            {
                var task = Arc[i];
                amount++;

                if (task.type == "C")
                {
                    Checklist.ShowTask(task, i + 1);
                }
                else if (task.type == "S")
                {
                    SimpleTask.ShowTask(task, i + 1);
                }
            }
            // FIXME: Console.ResetColor();
            Console.WriteLine($"\nWOW! DU HAR UTFÖRT {amount} UPPGIFTER!");
        }
    }
}
