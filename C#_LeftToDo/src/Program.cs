using System;
using LeftToDo.Lists;
using LeftToDo.Menus;

namespace LeftToDo
{
    class Program // Class Program is the starting point of the program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Välkommen till Din digitala Att Göra Lista!\n");
            TaskList list = new TaskList();
            Menu.ShowMainMenu(list);
        }
    }
}
