# To-Do List Application (Assignment 02)

This is a simple to-do list application built using native HTML, CSS, and JavaScript. It allows users to add, mark as complete, and delete tasks.

## File Structure

```
assignment_02/
├── index.html
├── style.css
├── script.js
└── main.go
```

- `index.html`: The main HTML file that provides the structure of the to-do list.
- `style.css`: Contains the CSS rules for styling the application, providing a modern and clean look.
- `script.js`: Implements the core functionality of the to-do list, including adding, completing, and deleting tasks.
- `main.go`: A simple Go server to serve the static files of the web application.

## How to Run

To run this application, you need to have Go installed on your system.

1.  Navigate to the `assignment_02` directory in your terminal:

    ```bash
    cd xu-assignments/assignment_02
    ```

2.  Run the Go server:

    ```bash
    go run main.go
    ```

3.  Open your web browser and go to `http://localhost:8080`.

Alternatively, you can open `index.html` directly in your web browser, but the Go server provides a more robust way to serve the application, especially if you were to expand it with backend functionalities.
