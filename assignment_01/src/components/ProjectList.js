import React from 'react';
import { Link } from 'react-router-dom';
import { projects } from '../data/mockData';
import './ProjectList.css';

function ProjectList() {
  return (
    <div className="project-list-container">
      <h2>Project Portfolio</h2>
      <div className="project-cards">
        {projects.map((project) => (
          <div key={project.id} className="project-card">
            <h3><Link to={`/projects/${project.id}`}>{project.name}</Link></h3>
            <p><strong>Department:</strong> {project.department}</p>
            <p><strong>Status:</strong> {project.status}</p>
            <p><strong>Person in Charge:</strong> {project.personInCharge}</p>
            <p><strong>Completion Rate:</strong> {project.completionRate}%</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default ProjectList;