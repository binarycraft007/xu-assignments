import React from 'react';
import { Link } from 'react-router-dom';
import './Header.css';

function Header() {
  return (
    <header className="App-header">
      <nav>
        <ul>
          <li>
            <Link to="/">Project Portfolio</Link>
          </li>
          <li>
            <Link to="/projects/new">Create Project</Link>
          </li>
          <li>
            <Link to="/tasks">My Tasks</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
}

export default Header;