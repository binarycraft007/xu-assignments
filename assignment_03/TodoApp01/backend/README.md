# To-Do List Backend (FastAPI)

This directory contains the backend for the To-Do List application, built with FastAPI and using SQLite as the database.

## Project Structure

```
backend/
├── main.py
├── database.py
├── models.py
├── schemas.py
├── __init__.py
└── test_main.py
```

- `main.py`: The main FastAPI application, defining all API endpoints.
- `database.py`: Handles the SQLite database connection and session management.
- `models.py`: Defines the SQLAlchemy ORM model for the `Todo` item.
- `schemas.py`: Defines Pydantic models for request and response data validation and serialization.
- `requirements.txt`: Lists all Python dependencies required for the backend.
- `test_main.py`: Contains unit and integration tests for the API endpoints.

## Setup and Installation

1.  **Navigate to the TodoApp01 directory:**

    ```bash
    cd assignment_03/TodoApp01
    ```

2.  **Install dependencies:**

    It is highly recommended to use a virtual environment.

    ```bash
    # Create a virtual environment and install dependencies using uv
    uv sync
    ```

## Running the Application

To start the FastAPI server, navigate to the `TodoApp01` directory (the parent of `backend`) and run:

```bash
uv run uvicorn backend.main:app --reload
```

The API will be accessible at `http://127.0.0.1:8000`.

## Running Tests

To run the tests, navigate to the `TodoApp01` directory (the parent of `backend`) and run:

```bash
uv run pytest backend/test_main.py
```

This will run all tests defined in `test_main.py` against a separate test database (`test.db`).

## API Documentation

Once the server is running, you can access the interactive API documentation (Swagger UI) at:

`http://127.0.0.1:8000/docs`

And the alternative ReDoc documentation at:

`http://127.0.0.1:8000/redoc`
