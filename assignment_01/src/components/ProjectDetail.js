import React from 'react';
import { useParams } from 'react-router-dom';
import { projects, tasks } from '../data/mockData';
import TaskList from './TaskList';
import './ProjectDetail.css';

function ProjectDetail() {
  const { id } = useParams();
  const project = projects.find((p) => p.id === id);
  const projectTasks = tasks.filter((task) => task.projectId === id);

  if (!project) {
    return <div className="project-detail-container">Project not found.</div>;
  }

  return (
    <div className="project-detail-container">
      <h2>{project.name} Dashboard</h2>
      <div className="project-info-grid">
        <div className="project-info-card">
          <h3>Project Details</h3>
          <p><strong>Department:</strong> {project.department}</p>
          <p><strong>Status:</strong> {project.status}</p>
          <p><strong>Person in Charge:</strong> {project.personInCharge}</p>
          <p><strong>Planned Start:</strong> {project.plannedStartTime}</p>
          <p><strong>Planned End:</strong> {project.plannedEndTime}</p>
          <p><strong>Completion Rate:</strong> {project.completionRate}%</p>
          <p><strong>Remarks:</strong> {project.remarks}</p>
        </div>
        <div className="project-info-card">
          <h3>Key Indicators</h3>
          <p><strong>Actual Completion:</strong> {project.actualCompletionTime || 'N/A'}</p>
          <p><strong>Time Consumed:</strong> {project.timeConsumed}</p>
          <p><strong>Dependencies:</strong> {project.dependencies.join(', ') || 'None'}</p>
          <p><strong>Milestones:</strong> {project.milestones.join(', ') || 'None'}</p>
        </div>
      </div>

      <h3>Progress Visualizations</h3>
      <div className="visualization-section">
        <div className="visualization-placeholder">
          <h4>Gantt Chart</h4>
          <p><em>(Placeholder for Gantt Chart visualization)</em></p>
        </div>
        <div className="visualization-placeholder">
          <h4>Kanban Board</h4>
          <p><em>(Placeholder for Kanban Board visualization)</em></p>
        </div>
        <div className="visualization-placeholder">
          <h4>Burndown/Burnup Chart</h4>
          <p><em>(Placeholder for Burndown/Burnup Chart visualization)</em></p>
        </div>
      </div>

      <h3>Tasks</h3>
      <TaskList tasks={projectTasks} />
    </div>
  );
}

export default ProjectDetail;
