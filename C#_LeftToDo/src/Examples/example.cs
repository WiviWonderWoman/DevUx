using System;
using System.Collections.Generic;
using LeftToDo.Tasks;

namespace LeftToDo.Example
{
    class Example
    {
        private static void Variables()
        {
            string declare;
            declare = "assign";

            var declareAndAssign = "Declare and " + declare;

            string x, y, z;
            x = "declare ";
            y = "assign ";
            z = "multiple variables";
            Console.WriteLine(x, y, z);
        }

        private static void Loops()
        {

            int length = 2;
            int[] arr = { 1, 2, 3 };
            bool keepLooping = true;

            for (int i = 0; i < length; i++)
            {
                // Take appropriate action
            }

            foreach (var nr in arr)
            {
                Console.WriteLine(nr);
            }

            while (keepLooping)
            {
                keepLooping = false;
            }

            do
            {
                keepLooping = true;
            } while (true);
        }
    }
}