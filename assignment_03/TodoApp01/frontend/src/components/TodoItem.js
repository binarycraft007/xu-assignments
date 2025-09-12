import React from 'react';

function TodoItem({ todo, toggleComplete, deleteTodo }) {
  return (
    <li className={todo.completed ? 'completed' : ''}>
      <span>{todo.title}</span>
      <div className="task-actions">
        <button className="complete-btn" onClick={() => toggleComplete(todo.id, todo.completed)}>
          {todo.completed ? 'Uncomplete' : 'Complete'}
        </button>
        <button className="delete-btn" onClick={() => deleteTodo(todo.id)}>
          Delete
        </button>
      </div>
    </li>
  );
}

export default TodoItem;
