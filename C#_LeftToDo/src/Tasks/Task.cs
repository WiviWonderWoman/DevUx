using System.Collections.Generic;

namespace LeftToDo
{
    public abstract class Task // Abstract super class, handles all types of Task
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
        public bool done
        {
            get;
            set;
        }
        public List<SimpleTask> subTask;
        // Marks task as Done
        public void MarkAsDone()
        {
            done = !done;
        }
        // Takes info from user, overrides in each subclass
        internal abstract void Create();
    }
}
