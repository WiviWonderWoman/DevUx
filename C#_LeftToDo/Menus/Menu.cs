using System;
using LeftToDo.Lists;
using LeftToDo.Tasks;

namespace LeftToDo.Menus
{
    abstract class Menu // Menu displays and handles out- / input
    {
        // Displays main menu 
        internal static void ShowMainMenu(TaskList list)
        {
            Console.WriteLine("\tHUVUDMENY\n[1] Visa Att-göra uppgifter\n[2] Visa Arkiverade uppgifter\n[0] Avsluta");
            ShowIntro();
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

        // Displays instructions to user
        private static void ShowIntro()
        {
            Console.WriteLine("\nVälj siffra inom [] följt av enter.\n");
        }

        // Displays submenu
        private static void ShowToDoMenu(TaskList list)
        {
            Console.WriteLine(
                "\n\tUNDERMENY: Vad vill du göra?\n" +
                "[1] Lägg till uppgift\n" +
                "[2] Markera / avmarkera uppgift\n" +
                "[3] Arkivera utförda uppgifter\n" +
                "[0] Tillbaka till HUVUDMENY"
            );
            ShowIntro();
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
                    ShowMainMenu(list);
                    break;
                default:
                    ShowErrorMsg();
                    break;
            }
        }

        // Display task menu
        private static void ShowTaskMenu(TaskList list)
        {
            list.ShowLeftToDo(list.ToDoList);
            Console.WriteLine(
                "\n\tUPPGIFTSMENY: Välj typ av uppgift:\n" +
                "[1] Enkel uppgift\n" +
                "[2] Checklista\n" +
                "[0] Tillbaka till HUVUDMENY"
            );
            ShowIntro();
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
                    ShowMainMenu(list);
                    break;
                default:
                    ShowErrorMsg();
                    break;
            }
        }

        // Displays option for user to return to main menu
        private static void ShowArchiveMenu(TaskList list)
        {
            Console.WriteLine("\n[0] Tillbaka till HUVUDMENY\n");
            var input = Console.ReadLine();
            if (input == "0")
            {
                ShowMainMenu(list);
            }
        }

        // Displays farewell message
        private static void ShowFarewell()
        {
            Console.WriteLine("\t\tTack för besöket och välkommen åter!\n\n");
            return;
        }

        // Displays error message
        private static void ShowErrorMsg()
        {
            Console.WriteLine("\t\tOgiltligt val. Vargod försök igen!\n\n");
        }
    }
}
