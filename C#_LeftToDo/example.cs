using System;
using System.Collections.Generic;

namespace LeftToDo
{
    private class Example
    {
        //[ACCESS][MODIFIER][RETURN][IDENTIFIER]
        private static int ReturnInt()
        {
            int returnNr = 1;
            var e = Example;
            // check for errors
            try
            {  //[INSTANCE METHOD CALL] [ARGUMENT]
                e.ReceiveIntNoReturn(returnNr);
            }
            catch (System.Exception) // handle errors
            {
                throw new Error();
            }
            return returnNr;
        }

        //[ACCESS] [RETURN] [IDENTIFIER] [PARAMETER]
        public void ReceiveIntNoReturn(int number)
        {
            try // check for errors
            {   //[STATIC METHOD CALL]
                var integer = ReturnInt();
                Console.WriteLine(number + integer);
            }
            catch (System.Exception) // handle errors
            {
                throw new Error();
            }
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