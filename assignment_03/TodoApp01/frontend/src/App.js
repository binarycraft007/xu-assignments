import React, { useState, useEffect } from 'react';
import axios from 'axios';
import TodoForm from './components/TodoForm';
import TodoList from './components/TodoList';
import FilterButtons from './components/FilterButtons';
import ClearButtons from './components/ClearButtons';

const API_BASE_URL = 'http://127.0.0.1:8000/api/v1/todos';

function App() {
  const [todos, setTodos] = useState([]);
  const [filter, setFilter] = useState('all'); // 'all', 'unfinished', 'completed'

  useEffect(() => {
    fetchTodos();
  }, [filter]);

  const fetchTodos = async () => {
    try {
      const response = await axios.get(API_BASE_URL, {
        params: { status: filter === 'all' ? undefined : filter }
      });
      setTodos(response.data);
    } catch (error) {
      console.error("Error fetching todos:", error);
    }
  };

  const addTodo = async (title) => {
    try {
      const response = await axios.post(API_BASE_URL, { title });
      setTodos([...todos, response.data]);
      fetchTodos(); // Re-fetch to ensure correct filtering/ordering if needed
    } catch (error) {
      console.error("Error adding todo:", error);
    }
  };

  const toggleComplete = async (id, completed) => {
    try {
      const response = await axios.put(`${API_BASE_URL}/${id}`, { completed: !completed });
      setTodos(todos.map(todo => (todo.id === id ? response.data : todo)));
      fetchTodos(); // Re-fetch to ensure correct filtering/ordering if needed
    } catch (error) {
      console.error("Error toggling complete status:", error);
    }
  };

  const deleteTodo = async (id) => {
    try {
      await axios.delete(`${API_BASE_URL}/${id}`);
      setTodos(todos.filter(todo => todo.id !== id));
      fetchTodos(); // Re-fetch to ensure correct filtering/ordering if needed
    } catch (error) {
      console.error("Error deleting todo:", error);
    }
  };

  const clearCompleted = async () => {
    try {
      await axios.delete(`${API_BASE_URL}?status=completed`);
      fetchTodos();
    } catch (error) {
      console.error("Error clearing completed todos:", error);
    }
  };

  const clearAll = async () => {
    try {
      await axios.delete(`${API_BASE_URL}?status=all`);
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