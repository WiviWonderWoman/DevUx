using System;
using System.Collections.Generic;

namespace LeftToDo
{
    public class TaskList /* ToDolist handles all lists and their methods */
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
            if (ToDo.Count < 1)
            {
                Console.WriteLine($"\t\tATT GÖRA LISTAN ÄR TOM!");
                return;
            }
            Console.WriteLine($"Status\tNr.\tUppgift\n");

            for (int i = 0; i < ToDo.Count; i++)
            {
                var item = ToDo[i];
                var index = i + 1;
                if (item.type == "S")
                {
                    SimpleTask.ShowTask(item, index);
                }
                else if (item.type == "C")
                {
                    Checklist.ShowTask(item, index);
                }
            }
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
                Console.WriteLine("\t\tDu skrev inte in en siffra. Försök igen.\n\n");
            }
            return number;
        }
        
        // Display archived Task
        internal void ShowArchive(List<Task> Arc)
        {
            if (Arc.Count < 1)
            {
                Console.WriteLine($"\t\tARKIVET ÄR TOMT.\n\n");
                return;
            }
            Console.WriteLine($"UTFÖRDA UPPGIFTER:\nStatus\tArkiv.\tUppgift\n");
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
            Console.WriteLine($"\t\tWOW! DU HAR UTFÖRT {amount} UPPGIFTER!\n\n");
        }
    }
}
