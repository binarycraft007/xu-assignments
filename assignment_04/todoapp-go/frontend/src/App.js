import React, { useState, useEffect } from 'react';
import axios from 'axios';
import TodoForm from './components/TodoForm';
import TodoList from './components/TodoList';
import FilterButtons from './components/FilterButtons';
import ClearButtons from './components/ClearButtons';

const API_BASE_URL = 'http://localhost:8000/api/v1';

function App() {
  const [todos, setTodos] = useState([]);
  const [filter, setFilter] = useState('all'); // 'all', 'unfinished', 'completed'

  useEffect(() => {
    fetchTodos();
  }, [filter]);

  const fetchTodos = async () => {
    try {
      let params = {};
      if (filter === 'unfinished') {
        params.completed = false;
      } else if (filter === 'completed') {
        params.completed = true;
      }

      const response = await axios.get(`${API_BASE_URL}/todos`, {
        params: params
      });
      setTodos(response.data.data);
    } catch (error) {
      console.error("Error fetching todos:", error);
    }
  };

  const addTodo = async (title) => {
    try {
      const response = await axios.post(`${API_BASE_URL}/todos`, { 
        title,
        description: "", // Default empty description
        priority: 0,     // Default priority
        due_date: null   // Default null due_date
      });
      setTodos([...todos, response.data.data]);
      fetchTodos(); // Re-fetch to ensure correct filtering/ordering if needed
    } catch (error) {
      console.error("Error adding todo:", error);
    }
  };

  const toggleComplete = async (id, completed) => {
    try {
      const response = await axios.patch(`${API_BASE_URL}/todos/${id}/toggle`);
      setTodos(todos.map(todo => (todo.id === id ? response.data.data : todo)));
      fetchTodos(); // Re-fetch to ensure correct filtering/ordering if needed
    } catch (error) {
      console.error("Error toggling complete status:", error);
    }
  };

  const deleteTodo = async (id) => {
    try {
      await axios.delete(`${API_BASE_URL}/todos/${id}`);
      setTodos(todos.filter(todo => todo.id !== id));
      fetchTodos(); // Re-fetch to ensure correct filtering/ordering if needed
    } catch (error) {
      console.error("Error deleting todo:", error);
    }
  };

  const clearCompleted = async () => {
    try {
      await axios.delete(`${API_BASE_URL}/todos/completed`);
      fetchTodos();
    } catch (error) {
      console.error("Error clearing completed todos:", error);
    }
  };

  const clearAll = async () => {
    try {
      await axios.delete(`${API_BASE_URL}/todos/all`);
      fetchTodos();
    } catch (error) {
      console.error("Error clearing all todos:", error);
    }
  };

  return (
    <div className="container">
      <h1>My To-Do List</h1>
      <TodoForm addTodo={addTodo} />
      <FilterButtons currentFilter={filter} setFilter={setFilter} />
      <TodoList todos={todos} toggleComplete={toggleComplete} deleteTodo={deleteTodo} />
      <ClearButtons clearCompleted={clearCompleted} clearAll={clearAll} />
    </div>
  );
}

export default App;