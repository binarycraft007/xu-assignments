import React from 'react';
import './TaskList.css';

function TaskList({ tasks }) {
  if (tasks.length === 0) {
    return <p>No tasks for this project.</p>;
  }

  return (
    <div className="task-list">
      {tasks.map((task) => (
        <div key={task.id} className="task-item">
          <h4>{task.name}</h4>
          <p><strong>Assignee:</strong> {task.assignee}</p>
          <p><strong>Due Date:</strong> {task.dueDate}</p>
          <p><strong>Status:</strong> {task.status}</p>
          <p><strong>Progress:</strong> {task.progress}%</p>
        </div>
      ))}
    </div>
  );
}

export default TaskList;
