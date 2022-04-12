using LeftToDo.Lists;
using System;
using System.Text;
using LeftToDo.Menus;
using System.Threading.Tasks;
using System.Collections.Generic;
using LeftToDo.Tasks;

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
