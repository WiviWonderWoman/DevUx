using System.Collections.Generic;

namespace LeftToDo
{
    // Abstract super class, handles all types of Task
    public abstract class Task
    {
        public string type
        {
            get;
            set;
        }
        public string description
        {
            get;
            set;
        }
        public string done
        {
            get;
            set;
        }
        public int daysLeft
        {
            get;
            set;
        }

        public List<SimpleTask> subTask;

        // Marks task as Done
        public void MarkAsDone()
        {
            if (done == null)
            {
                done = "X";
            }
            else if (done == "X")
            {
                done = null;
            }

        }

        // Takes info from user, overrides in each subclass
        internal abstract void Create();
    }
}
