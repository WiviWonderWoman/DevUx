using System;

namespace LeftToDo
{
    // Menu handles out- / input 
    abstract class Menu
    {
        // Mainmenu, handle user input with switch statement
        internal static void ShowMainMenu(TaskList list)
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
                    ShowArchiveMenu(list);
                    break;
                case "0":
                    ShowFarewell();
                    break;
                default:
                    ShowErrorMsg();
                    ShowMainMenu(list);
                    break;
            }
        }

        // Instructions to user
        private static void ShowIntro()
        {
            Console.WriteLine("Välj i menyn nedan, ange siffran inom [] och tryck enter.\n");
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
                    ShowMainMenu(list);
                    break;

                default:
                    ShowErrorMsg();
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
                "[2] Checklista\n" +
                "[0] HUVUDMENY\n\n"
            );

            var menu = Console.ReadLine();

            switch (menu)
            {
                case "1":
                    var simpelTask = new SimpleTask("");
                    simpelTask.Create();
                    list.AddToDoTask(simpelTask);
                    list.ShowLeftToDo(list.ToDoList);
                    ShowToDoMenu(list);
                    break;

                case "2":
                    var checklist = new Checklist("");
                    checklist.Create();
                    list.AddToDoTask(checklist);
                    list.ShowLeftToDo(list.ToDoList);
                    ShowToDoMenu(list);
                    break;

                case "0":
                    Console.Clear();
                    ShowMainMenu(list);
                    break;

                default:
                    ShowErrorMsg();
                    break;
            }

        }

        // Archive menu, lets user return to Mainmenu
        private static void ShowArchiveMenu(TaskList list)
        {
            Console.WriteLine("\n[0] HUVUDMENY\n\n");
            var input = Console.ReadLine();
            if (input == "0")
            {
                Console.Clear();
                ShowMainMenu(list);
            }
        }

        // Farewell message
        private static void ShowFarewell()
        {
            Console.WriteLine("Tack för besöket och välkommen åter!");
            return;
        }

        // Error message
        private static void ShowErrorMsg()
        {
            Console.WriteLine("\nOgiltligt val. Vargod försök igen!");
        }
    }
}
