# Technical Architecture Document: To-Do List Application

This document outlines the technical architecture for the To-Do List application, covering technical requirements, database schema, and API interface specifications.

## 1. Technical Requirements

*   **Frontend:** React
*   **Backend:** FastAPI (Python)
*   **Database:** SQLite

## 2. Project Structure

The project will be organized into two main directories within `TodoApp01`:

```
assignment_03/TodoApp01/
├── backend/
└── frontend/
```

## 3. Database Table Design (SQLite)

The application will use a single table, `todos`, to store to-do items. Each item will have an ID, a title (task description), and a completion status.

**Table Name:** `todos`

| Column Name | Data Type | Constraints           |
| :---------- | :-------- | :-------------------- |
| `id`        | INTEGER   | PRIMARY KEY AUTOINCREMENT |
| `title`     | TEXT      | NOT NULL              |
| `completed` | INTEGER   | DEFAULT 0 (0 for false, 1 for true) |

**SQL Statement to Create Table:**

```sql
CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    completed INTEGER DEFAULT 0
);
```

## 4. API Interface Specifications (FastAPI)

The backend will expose a RESTful API to manage to-do items. The base URL for the API will be `/api/v1/todos`.

### 4.1. To-Do Item Model

**Request Body for creating/updating a task:**

```json
{
    "title": "string",
    "completed": "boolean" (optional for update)
}
```

**Response Body for a task:**

```json
{
    "id": "integer",
    "title": "string",
    "completed": "boolean"
}
```

### 4.2. Endpoints

#### **GET /api/v1/todos**
*   **Description:** Retrieve all to-do items, with optional filtering by status.
*   **Query Parameters:**
    *   `status`: (Optional) `all`, `unfinished`, `completed`. Defaults to `all`.
*   **Responses:**
    *   `200 OK`: Returns a list of to-do items.
        ```json
        [
            {"id": 1, "title": "Buy groceries", "completed": false},
            {"id": 2, "title": "Walk the dog", "completed": true}
        ]
        ```

#### **GET /api/v1/todos/{id}**
*   **Description:** Retrieve a single to-do item by its ID.
*   **Path Parameters:**
    *   `id`: Integer, the ID of the to-do item.
*   **Responses:**
    *   `200 OK`: Returns the specified to-do item.
        ```json
        {"id": 1, "title": "Buy groceries", "completed": false}
        ```
    *   `404 Not Found`: If the to-do item with the given ID does not exist.

#### **POST /api/v1/todos**
*   **Description:** Create a new to-do item.
*   **Request Body:** `application/json`
    ```json
    {"title": "New task description"}
    ```
*   **Responses:**
    *   `201 Created`: Returns the newly created to-do item.
        ```json
        {"id": 3, "title": "New task description", "completed": false}
        ```
    *   `422 Unprocessable Entity`: If the `title` is missing or empty.

#### **PUT /api/v1/todos/{id}**
*   **Description:** Update an existing to-do item.
*   **Path Parameters:**
    *   `id`: Integer, the ID of the to-do item.
*   **Request Body:** `application/json`
    ```json
    {"title": "Updated task description", "completed": true}
    ```
    (Both `title` and `completed` are optional, but at least one must be provided for an update.)
*   **Responses:**
    *   `200 OK`: Returns the updated to-do item.
        ```json
        {"id": 1, "title": "Updated task description", "completed": true}
        ```
    *   `404 Not Found`: If the to-do item with the given ID does not exist.
    *   `422 Unprocessable Entity`: If the request body is empty or invalid.

#### **DELETE /api/v1/todos/{id}**
*   **Description:** Delete a to-do item by its ID.
*   **Path Parameters:**
    *   `id`: Integer, the ID of the to-do item.
*   **Responses:**
    *   `204 No Content`: If the to-do item was successfully deleted.
    *   `404 Not Found`: If the to-do item with the given ID does not exist.

#### **DELETE /api/v1/todos?status={status}**
*   **Description:** Clear to-do items based on their status.
*   **Query Parameters:**
    *   `status`: `completed` (to clear all completed tasks) or `all` (to clear all tasks).
*   **Responses:**
    *   `204 No Content`: If tasks were successfully cleared.
    *   `422 Unprocessable Entity`: If the `status` query parameter is missing or invalid.
