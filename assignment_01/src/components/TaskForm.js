import React, { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import './TaskForm.css';

function TaskForm() {
  const { projectId } = useParams();
  const navigate = useNavigate();
  const [task, setTask] = useState({
    name: '',
    description: '',
    assignee: '',
    dueDate: '',
    status: 'To Do',
    progress: 0,
    projectId: projectId || '', // Pre-fill projectId if available from URL
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setTask((prevTask) => ({
      ...prevTask,
      [name]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('New Task:', task);
    // In a real application, you would send this data to a backend API
    // For now, we'll just navigate back to the project detail or task list
    if (projectId) {
      navigate(`/projects/${projectId}`);
    } else {
      navigate('/tasks');
    }
  };

  return (
    <div className="task-form-container">
      <h2>Create New Task {projectId && `for Project ${projectId}`}</h2>
      <form onSubmit={handleSubmit} className="task-form">
        <div className="form-group">
          <label htmlFor="name">Task Name:</label>
          <input
            type="text"
            id="name"
            name="name"
            value={task.name}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="description">Description:</label>
          <textarea
            id="description"
            name="description"
            value={task.description}
            onChange={handleChange}
          ></textarea>
        </div>

        <div className="form-group">
          <label htmlFor="assignee">Assignee:</label>
          <input
            type="text"
            id="assignee"
            name="assignee"
            value={task.assignee}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="dueDate">Due Date:</label>
          <input
            type="date"
            id="dueDate"
            name="dueDate"
            value={task.dueDate}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="status">Status:</label>
          <select id="status" name="status" value={task.status} onChange={handleChange}>
            <option value="To Do">To Do</option>
            <option value="In Progress">In Progress</option>
            <option value="In Review">In Review</option>
            <option value="Done">Done</option>
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="progress">Progress (%):</label>
          <input
            type="number"
            id="progress"
            name="progress"
            value={task.progress}
            onChange={handleChange}
            min="0"
            max="100"
          />
        </div>

        {/* Hidden projectId field if it's pre-filled from URL */}
        {task.projectId && (
          <input type="hidden" name="projectId" value={task.projectId} />
        )}

        <button type="submit">Create Task</button>
      </form>
    </div>
  );
}

export default TaskForm;
