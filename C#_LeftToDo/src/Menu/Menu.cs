using System;

namespace LeftToDo
{
    // Menu handles out- / input 
    abstract class Menu
    {
        // Instructions to user
        private static void ShowIntro()
        {
            Console.WriteLine("Välj i menyn nedan, ange siffran inom [] och tryck enter.\n");
        }

        // Mainmenu, handle user input with switch statement
        internal static void ShowMain(TaskList list)
        {
            Console.Clear();
            ShowIntro();
            Console.WriteLine("HUVUDMENY\n[1] Visa Att-göra uppgifter\n[2] Visa Arkiverade uppgifter\n[0] Avsluta");

            var menu = Console.ReadLine();

            switch (menu)
            {
                case "1":
                    list.ShowLeftToDo(list.ToDoList);
                    ShowToDoMenu(list);
                    break;
                case "2":
                    list.ShowArchive(list.Archive);
                    ShowArcMenu(list);
                    break;
                case "0":
                    ShowFarewell();
                    break;
                default:
                    ShowError();
                    ShowMain(list);
                    break;
            }
        }

        // List menu, handle user input with switch statement
        private static void ShowToDoMenu(TaskList list)
        {
            ShowIntro();
            Console.WriteLine(
                "\nUNDERMENY: Vad vill du göra?\n" +
                "[1] Lägg till uppgift\n" +
                "[2] Markera / avmarkera uppgift\n" +
                "[3] Arkivera utförda uppgifter\n" +
                "[0] HUVUDMENY\n"
            );

            var menu = Console.ReadLine();

            switch (menu)
            {
                case "1":
                    ShowTaskMenu(list);
                    break;

                case "2":
                    list.FindTaskToMark(list);
                    ShowToDoMenu(list);
                    break;

                case "3":
                    list.ArchiveTask();
                    list.ShowLeftToDo(list.ToDoList);
                    ShowToDoMenu(list);
                    break;

                case "0":
                    Console.Clear();
                    ShowMain(list);
                    break;

                default:
                    ShowError();
                    break;
            }
        }

        // Task menu, handle user input with switch statement
        private static void ShowTaskMenu(TaskList list)
        {
            list.ShowLeftToDo(list.ToDoList);
            ShowIntro();
            Console.WriteLine(
                "\n\nUPPGIFTSMENY. Välj typ av uppgift:\n" +
                "[1] Enkel uppgift\n" +
                "[2] Deadline\n" +
                "[3] Checklista\n" +
                "[0] HUVUDMENY\n\n"
            );

            var menu = Console.ReadLine();

            switch (menu)
            {
                case "1":
                    var simpelTask = new SimpleTask("");
                    simpelTask.Create();
                    list.AddToDoList(simpelTask);
                    list.ShowLeftToDo(list.ToDoList);
                    ShowToDoMenu(list);
                    break;

                case "2":
                    var deadline = new Deadline("", 0);
                    deadline.Create();
                    list.AddToDoList(deadline);
                    list.ShowLeftToDo(list.ToDoList);
                    ShowToDoMenu(list);
                    break;

                case "3":
                    var checklist = new Checklist("");
                    checklist.Create();
                    list.AddToDoList(checklist);
                    list.ShowLeftToDo(list.ToDoList);
                    ShowToDoMenu(list);
                    break;

                case "0":
                    Console.Clear();
                    ShowMain(list);
                    break;

                default:
                    ShowError();
                    break;
            }

        }

        // Archive menu, lets user return to Mainmenu
        private static void ShowArcMenu(TaskList list)
        {
            Console.WriteLine("\n[0] HUVUDMENY\n\n");
            var input = Console.ReadLine();
            if (input == "0")
            {
                Console.Clear();
                ShowMain(list);
            }
        }

        // Farewell message
        private static void ShowFarewell()
        {
            Console.WriteLine("Tack för besöket och välkommen åter!");
            return;
        }

        // Error message
        private static void ShowError()
        {
            Console.WriteLine("\nOgiltligt val. Vargod försök igen!");
        }
    }
}
