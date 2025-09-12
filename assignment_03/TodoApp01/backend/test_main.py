from fastapi.testclient import TestClient
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from main import app, get_db
from database import Base

# Setup for testing database
SQLALCHEMY_DATABASE_URL = "sqlite:///./test.db"

engine = create_engine(
    SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False}
)
TestingSessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)


Base.metadata.create_all(bind=engine)

def override_get_db():
    try:
        db = TestingSessionLocal()
        yield db
    finally:
        db.close()


app.dependency_overrides[get_db] = override_get_db

client = TestClient(app)

def test_create_todo():
    response = client.post(
        "/api/v1/todos",
        json={"title": "Test Todo", "completed": False},
    )
    assert response.status_code == 201
    assert response.json()["title"] == "Test Todo"
    assert response.json()["completed"] == False
    assert "id" in response.json()

def test_read_todos():
    response = client.get("/api/v1/todos")
    assert response.status_code == 200
    assert isinstance(response.json(), list)

def test_read_single_todo():
    # First create a todo
    response = client.post(
        "/api/v1/todos",
        json={"title": "Single Todo", "completed": False},
    )
    todo_id = response.json()["id"]

    response = client.get(f"/api/v1/todos/{todo_id}")
    assert response.status_code == 200
    assert response.json()["title"] == "Single Todo"

def test_read_nonexistent_todo():
    response = client.get("/api/v1/todos/999")
    assert response.status_code == 404

def test_update_todo():
    # First create a todo
    response = client.post(
        "/api/v1/todos",
        json={"title": "Update Me", "completed": False},
    )
    todo_id = response.json()["id"]

    response = client.put(
        f"/api/v1/todos/{todo_id}",
        json={"title": "Updated Todo", "completed": True},
    )
    assert response.status_code == 200
    assert response.json()["title"] == "Updated Todo"
    assert response.json()["completed"] == True

def test_delete_todo():
    # First create a todo
    response = client.post(
        "/api/v1/todos",
        json={"title": "Delete Me", "completed": False},
    )
    todo_id = response.json()["id"]

    response = client.delete(f"/api/v1/todos/{todo_id}")
    assert response.status_code == 204

    response = client.get(f"/api/v1/todos/{todo_id}")
    assert response.status_code == 404

def test_clear_completed_todos():
    # Create some todos
    client.post("/api/v1/todos", json={"title": "Task 1", "completed": False})
    client.post("/api/v1/todos", json={"title": "Task 2", "completed": True})
    client.post("/api/v1/todos", json={"title": "Task 3", "completed": True})

    response = client.delete("/api/v1/todos?status=completed")
    assert response.status_code == 204

    response = client.get("/api/v1/todos?status=completed")
    assert len(response.json()) == 0

def test_clear_all_todos():
    # Create some todos
    client.post("/api/v1/todos", json={"title": "Task A", "completed": False})
    client.post("/api/v1/todos", json={"title": "Task B", "completed": True})

    response = client.delete("/api/v1/todos?status=all")
    assert response.status_code == 204

    response = client.get("/api/v1/todos")
    assert len(response.json()) == 0