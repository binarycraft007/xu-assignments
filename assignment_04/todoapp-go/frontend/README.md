# To-Do List Frontend (React)

This directory contains the frontend for the To-Do List application, built with React. It interacts with the FastAPI backend to manage to-do items.

## Project Structure

```
frontend/
├── public/
│   └── ... (standard create-react-app public files)
├── src/
│   ├── App.js
│   ├── index.js
│   ├── index.css
│   └── components/
│       ├── ClearButtons.js
│       ├── FilterButtons.js
│       ├── TodoForm.js
│       ├── TodoItem.js
│       └── TodoList.js
├── package.json
├── package-lock.json
└── README.md
```

- `App.js`: The main React component that manages the application state and orchestrates other components.
- `index.js`: The entry point for the React application.
- `index.css`: Global styles for the application.
- `components/`: Contains individual React components for different parts of the UI.

## Setup and Installation

1.  **Navigate to the frontend directory:**

    ```bash
    cd assignment_03/TodoApp01/frontend
    ```

2.  **Install dependencies:**

    ```bash
    npm install
    ```

## Running the Application

To start the React development server, run the following command from the `frontend` directory:

```bash
npm start
```

The application will typically open in your browser at `http://localhost:3000`.

**Note:** Ensure the backend server is running (usually on `http://127.0.0.1:8000`) for the frontend to be able to fetch and manage to-do items.