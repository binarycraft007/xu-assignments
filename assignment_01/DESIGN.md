## Design Document: Project Task Management Software Prototype

### 1. Feature List

Based on the current implementation, the prototype includes the following features:

*   **Project Portfolio View:**
    *   Displays a list of all projects with key information (Name, Department, Status, Person in Charge, Completion Rate).
    *   Each project card is clickable, leading to a detailed project dashboard.
    *   Responsive layout for different screen sizes.
*   **Project Creation:**
    *   A dedicated form to create new projects with fields for: Project Name, Department, Status, Person in Charge, Members, Planned Start Time, Planned End Time, Dependencies, Milestones, and Remarks.
    *   Basic form validation (required fields).
    *   Navigates back to the project list upon submission (data is currently logged to console).
    *   Responsive form layout.
*   **Project Dashboard (Detail View):**
    *   Displays comprehensive details of a selected project, including all attributes from project creation.
    *   Separated into "Project Details" and "Key Indicators" sections.
    *   Includes placeholder sections for "Progress Visualizations" (Gantt Chart, Kanban Board, Burndown/Burnup Chart).
    *   Lists all tasks associated with the project.
    *   Responsive layout for project information and visualizations.
*   **Task Listing (within Project Dashboard):**
    *   Displays a list of tasks for the specific project, showing Task Name, Assignee, Due Date, Status, and Progress.
    *   Responsive layout for task items.
*   **My Tasks View:**
    *   Displays a list of tasks assigned to a specific user (currently hardcoded to "John" for demonstration).
    *   Shows Task Name, Assignee, Due Date, Status, and Progress.
    *   Responsive layout.
*   **Task Creation:**
    *   A dedicated form to create new tasks with fields for: Task Name, Description, Assignee, Due Date, Status, and Progress (%).
    *   Can be accessed from the navigation bar or potentially from a project detail page (though not explicitly linked yet).
    *   Pre-fills `projectId` if accessed via a project-specific route.
    *   Navigates back to the project detail or task list upon submission (data is currently logged to console).
    *   Responsive form layout.
*   **Navigation:**
    *   A persistent header with links to "Project Portfolio," "Create Project," and "My Tasks."
    *   Uses `react-router-dom` for client-side routing.
    *   Responsive navigation menu for smaller screens.
*   **Styling:**
    *   Modern, "technological" aesthetic with a dark theme, accent colors, and subtle shadows.
    *   Adaptive design supporting both PC and mobile terminals through CSS media queries and flexible layouts (Flexbox/Grid).
*   **Mock Data:**
    *   Uses static `mockData.js` for projects and tasks, simulating a backend data source.

### 2. Technical Architecture

**2.1. Frontend Framework:**
*   **React.js:** The application is built using React, a JavaScript library for building user interfaces. It leverages a component-based architecture for modularity and reusability.
*   **Create React App (CRA):** The project was bootstrapped with CRA, providing a pre-configured development environment with Webpack, Babel, and other tools.

**2.2. Routing:**
*   **React Router DOM (v6):** Used for declarative routing within the single-page application. It manages navigation between different views (Project List, Project Detail, Project Form, My Tasks, Task Form) without full page reloads.

**2.3. State Management:**
*   **React Hooks (`useState`, `useParams`, `useNavigate`):** Local component state is managed using `useState`. `useParams` is used to extract parameters from the URL (e.g., `projectId`), and `useNavigate` is used for programmatic navigation.
*   **Mock Data:** Data is currently stored in `src/data/mockData.js` as JavaScript arrays. This serves as a placeholder for a future backend API.

**2.4. Styling:**
*   **CSS Modules/Global CSS:** Standard CSS files are used for styling, with a mix of global styles (`index.css`, `App.css`) and component-specific styles (`Header.css`, `ProjectList.css`, etc.).
*   **Responsive Design:** CSS media queries are extensively used to adapt the layout and styling for various screen sizes, ensuring a consistent user experience on both desktop and mobile devices. Flexbox and CSS Grid are employed for flexible layouts.

**2.5. Project Structure:**

```
assignment_01/
├── public/
├── src/
│   ├── components/
│   │   ├── Header.css
│   │   ├── Header.js
│   │   ├── MyTasks.css
│   │   ├── MyTasks.js
│   │   ├── ProjectDetail.css
│   │   ├── ProjectDetail.js
│   │   ├── ProjectForm.css
│   │   ├── ProjectForm.js
│   │   ├── ProjectList.css
│   │   ├── ProjectList.js
│   │   ├── TaskForm.css
│   │   ├── TaskForm.js
│   │   ├── TaskList.css
│   │   └── TaskList.js
│   ├── data/
│   │   └── mockData.js
│   ├── App.css
│   ├── App.js
│   ├── App.test.js
│   ├── index.css
│   └── index.js
├── package.json
├── README.md
└── ... (other CRA generated files)
```

### 3. Deployment Instructions

This prototype is a standard Create React App (CRA) project, and its deployment follows typical React application deployment procedures.

**3.1. Prerequisites:**
*   Node.js (LTS version recommended)
*   npm (Node Package Manager) or Yarn

**3.2. Local Development:**
1.  **Navigate to the project directory:**
    ```bash
    cd /home/elliot/Desktop/Projects/xu-assignments/assignment_01
    ```
2.  **Install dependencies:**
    ```bash
    npm install
    ```
3.  **Start the development server:**
    ```bash
    npm start
    ```
    This will open the application in your default web browser at `http://localhost:3000` (or another available port). The development server provides hot-reloading for a smooth development experience.

**3.3. Building for Production:**
1.  **Navigate to the project directory:**
    ```bash
    cd /home/elliot/Desktop/Projects/xu-assignments/assignment_01
    ```
2.  **Build the application:**
    ```bash
    npm run build
    ```
    This command compiles the React application into static files (HTML, CSS, JavaScript) and places them in the `build/` directory. These files are optimized for production, including minification and bundling.

**3.4. Deployment to a Static Hosting Service:**
The `build` directory can be deployed to any static file hosting service. Popular options include:

*   **Netlify:**
    1.  Connect your Git repository (GitHub, GitLab, Bitbucket).
    2.  Configure build settings: `Build command: npm run build`, `Publish directory: build`.
    3.  Netlify will automatically deploy your site on every push to the configured branch.
*   **Vercel:**
    1.  Connect your Git repository.
    2.  Vercel automatically detects Create React App and configures build settings.
    3.  Deploy on every push.
*   **GitHub Pages:**
    1.  Install `gh-pages` package: `npm install --save-dev gh-pages`.
    2.  Add `homepage` field to `package.json`: `"homepage": "http://<YOUR_GITHUB_USERNAME>.github.io/<YOUR_REPO_NAME>"`.
    3.  Add `deploy` script to `package.json` scripts: `"predeploy": "npm run build"`, `"deploy": "gh-pages -d build"`.
    4.  Run `npm run deploy`.
*   **Amazon S3 + CloudFront:**
    1.  Upload the contents of the `build` folder to an S3 bucket configured for static website hosting.
    2.  Optionally, use Amazon CloudFront for CDN and SSL.

**3.5. Deployment to a Node.js Server (e.g., Express):**
If you have a Node.js backend, you can serve the static `build` files:
1.  Install `serve-static` or `express.static` in your Node.js application.
2.  Configure your server to serve the `build` directory:
    ```javascript
    const express = require('express');
    const path = require('path');
    const app = express();

    app.use(express.static(path.join(__dirname, 'client/build'))); // Assuming 'client' is your React app folder

    app.get('*', (req, res) => {
      res.sendFile(path.join(__dirname, 'client/build', 'index.html'));
    });

    const PORT = process.env.PORT || 5000;
    app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
    ```
    (Adjust `client/build` path as per your project structure).

---

## Competitive Product Analysis Report

### 1. Introduction

This report provides a general analysis of features and design logic commonly found in leading project task management software. Given the broad landscape of such tools (e.g., Jira, Asana, Trello, Monday.com, ClickUp), this analysis focuses on common patterns and underlying principles rather than specific product-by-product comparisons. The goal is to understand the established best practices and user expectations in this domain.

### 2. Feature List of Competing Products (General)

Most mature project management tools offer a comprehensive set of features, often categorized as follows:

*   **Project & Portfolio Management:**
    *   **Project Creation & Configuration:** Detailed project setup (name, description, dates, owner, team, status, budget, custom fields).
    *   **Project Hierarchy:** Ability to nest projects, sub-projects, or link related projects.
    *   **Portfolio View:** High-level overview of multiple projects, often with health indicators, progress summaries, and filtering capabilities.
    *   **Templates:** Pre-defined project templates for common workflows.
*   **Task Management:**
    *   **Task Creation & Details:** Comprehensive task fields (name, description, assignee, due date, priority, status, subtasks, attachments, comments).
    *   **Task Assignment & Collaboration:** Assigning tasks to individuals or teams, real-time comments, mentions, and file sharing.
    *   **Status & Progress Tracking:** Customizable task statuses (e.g., To Do, In Progress, In Review, Done), percentage completion, time tracking.
    *   **Dependencies:** Linking tasks with "blocks," "is blocked by," "relates to" relationships.
    *   **Recurring Tasks:** Ability to set tasks that repeat on a schedule.
*   **Progress Visualization & Reporting:**
    *   **Gantt Charts:** Visual timelines showing task durations, dependencies, and critical paths.
    *   **Kanban Boards:** Visual workflow management, allowing tasks to be moved through stages (columns).
    *   **List Views:** Simple, tabular display of tasks with sorting and filtering.
    *   **Calendar Views:** Tasks displayed on a calendar.
    *   **Table/Spreadsheet Views:** Detailed, customizable tables for data entry and analysis.
    *   **Dashboards:** Customizable widgets displaying key metrics (e.g., task completion, overdue tasks, team workload).
    *   **Burndown/Burnup Charts:** Tracking work remaining/completed against time.
    *   **Custom Reports:** Generating reports based on various project data.
*   **Communication & Collaboration:**
    *   **Comments & Discussions:** In-task or project-level comment threads.
    *   **Notifications:** In-app, email, and sometimes mobile push notifications for updates, mentions, and deadlines.
    *   **Integrations:** Connections with communication tools (Slack, Microsoft Teams), file storage (Google Drive, Dropbox), version control (GitHub, GitLab), and CRM systems.
*   **Resource Management:**
    *   **Workload Management:** Visualizing team member workload and capacity.
    *   **Time Tracking:** Logging time spent on tasks.
*   **Automation:**
    *   **Rule-Based Automation:** Setting up rules to automate actions (e.g., "when task status changes to 'Done', notify project leader").
*   **Access Control & Permissions:**
    *   Role-based access control (RBAC) to manage who can view, edit, or create different types of information.
*   **Search & Filtering:**
    *   Powerful search capabilities and advanced filtering options to quickly find projects, tasks, or specific data points.

### 3. Design Logic Behind Competing Products

The design logic of leading project management tools is driven by several core principles aimed at enhancing productivity, collaboration, and visibility:

**3.1. User-Centric Design & Intuitive Interfaces:**
*   **Logic:** Complex project data is presented in easily digestible formats. The UI prioritizes clarity, minimizing cognitive load. Common actions are readily accessible. This caters to users with varying technical proficiencies.
*   **Implementation:** Clean layouts, consistent iconography, drag-and-drop functionality (e.g., Kanban boards), and clear navigation paths. Onboarding flows are often guided.

**3.2. Flexibility & Customization:**
*   **Logic:** No two teams or projects are identical. Tools provide flexibility to adapt to diverse workflows, terminologies, and reporting needs.
*   **Implementation:** Customizable fields, statuses, workflows, and views (Kanban, Gantt, List, Calendar). Users can often create their own templates or dashboards. This empowers teams to tailor the tool to their specific context.

**3.3. Visual Communication & Data Visualization:**
*   **Logic:** Visual representations are more effective than raw data for quickly understanding project health, progress, and bottlenecks. "A picture is worth a thousand words" applies strongly here.
*   **Implementation:** Extensive use of Gantt charts for timelines, Kanban boards for workflow, burndown/burnup charts for progress, and customizable dashboards with various widgets (e.g., pie charts for status distribution, bar charts for workload). Color-coding is often used to highlight status or priority.

**3.4. Collaboration & Communication Hub:**
*   **Logic:** Projects are team efforts. The tool should facilitate seamless communication and information sharing, reducing the need to switch between multiple applications.
*   **Implementation:** In-task comments, @mentions, file attachments, activity feeds, and integrations with popular communication platforms (Slack, Teams). Centralizing discussions around tasks and projects prevents information silos.

**3.5. Proactive Insights & Automation:**
*   **Logic:** Moving beyond reactive management to proactive problem-solving. Identifying potential issues before they escalate saves time and resources. Automating repetitive tasks frees up human effort for more strategic work.
*   **Implementation:** Rule-based automation (e.g., "if task is overdue, change status to 'Blocked' and notify assignee"), AI-driven predictions for delays (as outlined in the PRD), and smart notifications. This shifts the focus from manual oversight to system-assisted management.

**3.6. Scalability & Performance:**
*   **Logic:** Tools must perform well for small teams and large enterprises with thousands of projects and users. Data integrity and quick response times are critical.
*   **Implementation:** Robust backend architectures (often microservices), efficient database designs, caching mechanisms, and optimized frontend rendering. Non-functional requirements like performance, reliability, and security are paramount.

**3.7. Accessibility & Cross-Platform Support:**
*   **Logic:** Users need to access project information from anywhere, on any device.
*   **Implementation:** Web-based applications with responsive design (like the prototype), dedicated mobile applications (iOS/Android), and sometimes desktop clients. APIs for third-party integrations are also common.

### 4. Conclusion

The current React prototype lays a foundational groundwork for a project task management system, aligning with many core features found in established products, particularly in project and task creation/viewing. The "technological" aesthetic and responsive design are good starting points for usability.

To evolve into a competitive product, the next steps would involve:
*   Implementing the visualization features (Gantt, Kanban, Burndown charts) with actual data.
*   Developing a robust backend for persistent data storage, user authentication, and API endpoints.
*   Adding advanced collaboration features (comments, notifications).
*   Integrating the AI requirements for predictive warnings and task prioritization.
*   Expanding user roles and permissions.

The design logic of existing products emphasizes user-centricity, flexibility, visual communication, and proactive insights. Future development of this prototype should continue to prioritize these principles to create a truly effective and competitive tool.