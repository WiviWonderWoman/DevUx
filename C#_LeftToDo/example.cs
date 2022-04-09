using System;

namespace Example
{
    private class Example
    {
        private static void Loops()
        {
            int length = 2;
            bool keepLooping = true;
            int[] arr = { 1, 2, 3 };

            for (int i = 0; i < length; i++)
            {
                // Take appropriate action
            }

            while (keepLooping)
            {
                keepLooping = false;
            }

            do
            {
                !keepLooping = true;
            } while (true);

            foreach (var nr in arr)
            {
                Console.WriteLine(nr);
            }
        }


        public void ReceiveStringNoReturn(string parameter)
        {
            try
            {
                var returnValue = ReturnString();
                Console.WriteLine(parameter + returnValue);
            }
            catch (System.Exception)
            {
                throw new Error();
            }
        }

        private static int ReturnString()
        {
            string returnString = " Retur";

            string argument;
            argument = "Argument";

            Example e = Example;

            try
            {
                e.ReceiveStringNoReturn(argument);
            }
            catch (System.Exception)
            {
                throw new Error();
            }
            return returnString;
        }



        // private struct not exported 
        private struct task
        {
            string taskType; // describes task type user choose
            string description; // what action / ToDo task user inputs
            bool done; // indicates if task is done
            List<SimpleTask> subTask; // if task is of type checklist, a slice with subtasks 
        }

        // public struct exported 
        public struct Task
        {
            string TaskType; // describes task type user choose
            string Description; // what action / ToDo task user inputs
            bool Done; // indicates if task is done
            List<SimpleTask> SubTask; // if task is of type checklist, a slice with subtasks 
        }



        /*

        //TODO: DELETE

        */

    }
}