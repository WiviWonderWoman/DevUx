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
                if (task.done == true)
                {
                    AddTaskToArhive(task);
                    ToDoList.Remove(task);
                }
            }
        }

        // Display List of Task 
        internal void ShowLeftToDo(List<Task> ToDo)
        {
            Console.Clear();
            if (ToDo.Count < 1)
            {
                Console.WriteLine($"\n\nATT GÖRA LISTAN ÄR TOM!\n\n");
                return;
            }
            Console.WriteLine($"Status\tNr.\tUppgift\n");

            for (int i = 0; i < ToDo.Count; i++)
            {
                Console.ResetColor();
                var item = ToDo[i];
                var index = i + 1;

                if (!item.done)
                {
                    Console.ForegroundColor = ConsoleColor.Red;
                }
                else if (item.done)
                {
                    Console.ForegroundColor = ConsoleColor.Yellow;
                }

                if (item.type == "S")
                {
                    SimpleTask.ShowTask(item, index);
                }
                else if (item.type == "C")
                {
                    Checklist.ShowTask(item, index);
                    ShowSubTask(item.subTask, index);
                }
                else if (item.type == "D")
                {
                    Deadline.ShowTask(item, index);
                }
            }
            Console.ResetColor();
        }

        // Display Checklist Task
        private void ShowSubTask(List<SimpleTask> subList, int outer)
        {
            for (int i = 0; i < subList.Count; i++)
            {
                Console.ResetColor();
                var inner = i + 1;

                var item = subList[i];

                if (!item.done)
                {
                    Console.ForegroundColor = ConsoleColor.Red;
                }
                else if (item.done)
                {
                    Console.ForegroundColor = ConsoleColor.Yellow;
                }

                SimpleTask.ShowTask(item, outer, inner);
            }
            Console.ResetColor();
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
            Console.Clear();
            if (Arc.Count < 1)
            {
                Console.WriteLine($"\n\nARKIVET ÄR TOMT.\n\n");
            }
            else
            {
                Console.WriteLine($"UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift\n");
                Console.ForegroundColor = ConsoleColor.Green;

                int amount = 0;
                for (int i = 0; i < Arc.Count; i++)
                {
                    var task = Arc[i];

                    if (task.type == "C")
                    {
                        Checklist.ShowTask(task, i + 1);
                        // Console.WriteLine($"[{task.done}]\t\t{task.description}\n");
                        amount++;

                        for (int j = 0; j < task.subTask.Count; j++)
                        {

                            var subTask = task.subTask[j];
                            SimpleTask.ShowTask(subTask, j + 1);
                            // Console.WriteLine($"\t\t[{subTask.done}]\t\t{subTask.description}\n");
                            amount++;
                        }
                    }
                    else
                    {
                        Checklist.ShowTask(task, i + 1);
                        // Console.WriteLine($"[{task.done}]\t\t{task.description}\n");
                        amount++;
                    }
                }
                Console.ResetColor();
                Console.WriteLine($"\nWOW! DU HAR UTFÖRT {amount} UPPGIFTER!");

            }
        }
    }
}
