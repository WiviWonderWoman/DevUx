using System;
using System.Collections.Generic;

namespace LeftToDo
{

    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Välkommen till Din digitala Att Göra Lista!\n");

            TaskList list = new TaskList();

            Menu.ShowMainMenu(list);
        }
    }
}
