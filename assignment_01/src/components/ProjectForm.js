import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './ProjectForm.css';

function ProjectForm() {
  const navigate = useNavigate();
  const [project, setProject] = useState({
    name: '',
    department: '',
    status: 'Not Started',
    personInCharge: '',
    members: '',
    plannedStartTime: '',
    plannedEndTime: '',
    dependencies: '',
    milestones: '',
    remarks: '',
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setProject((prevProject) => ({
      ...prevProject,
      [name]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('New Project:', project);
    // In a real application, you would send this data to a backend API
    // For now, we'll just navigate back to the project list
    navigate('/');
  };

  return (
    <div className="project-form-container">
      <h2>Create New Project</h2>
      <form onSubmit={handleSubmit} className="project-form">
        <div className="form-group">
          <label htmlFor="name">Project Name:</label>
          <input
            type="text"
            id="name"
            name="name"
            value={project.name}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="department">Department:</label>
          <input
            type="text"
            id="department"
            name="department"
            value={project.department}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="status">Status:</label>
          <select id="status" name="status" value={project.status} onChange={handleChange}>
            <option value="Not Started">Not Started</option>
            <option value="Planning">Planning</option>
            <option value="In Progress">In Progress</option>
            <option value="On Hold">On Hold</option>
            <option value="Completed">Completed</option>
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="personInCharge">Person in Charge:</label>
          <input
            type="text"
            id="personInCharge"
            name="personInCharge"
            value={project.personInCharge}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="members">Members (comma-separated):</label>
          <input
            type="text"
            id="members"
            name="members"
            value={project.members}
            onChange={handleChange}
          />
        </div>

        <div className="form-group">
          <label htmlFor="plannedStartTime">Planned Start Time:</label>
          <input
            type="date"
            id="plannedStartTime"
            name="plannedStartTime"
            value={project.plannedStartTime}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="plannedEndTime">Planned End Time:</label>
          <input
            type="date"
            id="plannedEndTime"
            name="plannedEndTime"
            value={project.plannedEndTime}
            onChange={handleChange}
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="dependencies">Dependencies (comma-separated):</label>
          <input
            type="text"
            id="dependencies"
            name="dependencies"
            value={project.dependencies}
            onChange={handleChange}
          />
        </div>

        <div className="form-group">
          <label htmlFor="milestones">Milestones (comma-separated):</label>
          <input
            type="text"
            id="milestones"
            name="milestones"
            value={project.milestones}
            onChange={handleChange}
          />
        </div>

        <div className="form-group">
          <label htmlFor="remarks">Remarks:</label>
          <textarea
            id="remarks"
            name="remarks"
            value={project.remarks}
            onChange={handleChange}
          ></textarea>
        </div>

        <button type="submit">Create Project</button>
      </form>
    </div>
  );
}

export default ProjectForm;