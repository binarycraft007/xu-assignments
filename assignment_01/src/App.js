import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/Header';
import ProjectList from './components/ProjectList';
import ProjectDetail from './components/ProjectDetail';
import ProjectForm from './components/ProjectForm';
import MyTasks from './components/MyTasks';
import TaskForm from './components/TaskForm';
import './App.css';

function App() {
  return (
    <Router>
      <div className="App">
        <Header />
        <main>
          <Routes>
            <Route path="/" element={<ProjectList />} />
            <Route path="/projects/new" element={<ProjectForm />} />
            <Route path="/projects/:id" element={<ProjectDetail />} />
            <Route path="/tasks" element={<MyTasks />} />
            <Route path="/projects/:projectId/tasks/new" element={<TaskForm />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;