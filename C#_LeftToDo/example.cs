
namespace LeftToDo
{
    private class Example
    {
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

        //     [ACCESS] [MODIFIER] [RETURN] [IDENTIFIER]
        private static int ReturnInt()
        {
            try // check for errors
            {
                int number = 1;
                //      [METHOD CALL]  [ARGUMENT]
                Example.ReceiveIntNoReturn(number);
            }
            catch (System.Exception) // handle errors
            {
                throw new Error();
            }
            finally
            {
                return number;
            }
        }

        //     [ACCESS] [RETURN] [IDENTIFIER] [PARAMETER]
        public void ReceiveIntNoReturn(int number)
        {
            try // check for errors
            {
                //              [STATIC METHOD CALL]
                var integer = ReturnInt();
                Console.WriteLine(number + integer);
            }
            catch (System.Exception) // handle errors
            {
                throw new Error();
            }
        }
    }
}