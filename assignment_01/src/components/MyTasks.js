import React from 'react';
import { tasks } from '../data/mockData';
import TaskList from './TaskList';
import './MyTasks.css';

function MyTasks() {
  // For demonstration, let's assume the current user is 'John'
  const myAssignedTasks = tasks.filter((task) => task.assignee === 'John');

  return (
    <div className="my-tasks-container">
      <h2>My Assigned Tasks</h2>
      {myAssignedTasks.length > 0 ? (
        <TaskList tasks={myAssignedTasks} />
      ) : (
        <p>No tasks assigned to you.</p>
      )}
    </div>
  );
}

export default MyTasks;
