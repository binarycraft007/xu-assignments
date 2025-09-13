# Todo Application Backend API

A RESTful to-do item management API based on gin framework, supporting full CRUD operations, status toggling, bulk deletion, pagination, filtering, health checks, Swagger documentation, and more.

## Technology Stack
- gin
- golang
- mariadb 12.0.2
- Swagger(for golang)

## Database Design

-   **Address**: localhost:3306
-   **Username**: monty
-   **Password**: test001
-   **Database Name**: todoapp

## API Specification

### Basic Information
- **Base URL**: `http://localhost:8000`
- **API Prefix**: `/api/v1`
- **Data Format**: JSON
- **HTTP Status Codes**: Follows RESTful conventions

### Endpoint List

#### 1. Get All Todos
```http
GET /api/v1/todos
```

**Query Parameters**:
- `completed` (optional): boolean - Filter by completion status
- `limit` (optional): integer - Limit the number of returned items, default is 100
- `offset` (optional): integer - Offset for pagination, default is 0

**Response Example**:
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "title": "Learn React",
      "description": "Complete the basic React tutorial",
      "completed": false,
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z",
      "priority": 1,
      "due_date": null
    }
  ],
  "total": 1
}
```

#### 2. Create a Todo
```http
POST /api/v1/todos
```

**Request Body**:
```json
{
  "title": "New Todo Item",
  "description": "Description (optional)",
  "priority": 0,
  "due_date": null
}
```

**Response Example**:
```json
{
  "code": 201,
  "message": "Todo created successfully",
  "data": {
    "id": 2,
    "title": "New Todo Item",
    "description": "Description (optional)",
    "completed": false,
    "created_at": "2024-01-01T11:00:00Z",
    "updated_at": "2024-01-01T11:00:00Z",
    "priority": 0,
    "due_date": null
  }
}
```

#### 3. Update a Todo
```http
PUT /api/v1/todos/{todo_id}
```

**Path Parameters**:
- `todo_id`: integer - The ID of the todo item

**Request Body**:
```json
{
  "title": "Updated Title",
  "description": "Updated Description",
  "completed": true,
  "priority": 2,
  "due_date": "2024-12-31T23:59:59Z"
}
```

**Response Example**:
```json
{
  "code": 200,
  "message": "Todo updated successfully",
  "data": {
    "id": 1,
    "title": "Updated Title",
    "description": "Updated Description",
    "completed": true,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z",
    "priority": 2,
    "due_date": "2024-12-31T23:59:59Z"
  }
}
```

#### 4. Delete a Todo
```http
DELETE /api/v1/todos/{todo_id}
```

**Path Parameters**:
- `todo_id`: integer - The ID of the todo item

**Response Example**:
```json
{
  "code": 200,
  "message": "Todo deleted successfully"
}
```

#### 5. Mark as Complete / Incomplete (Toggle)
```http
PATCH /api/v1/todos/{todo_id}/toggle
```

**Path Parameters**:
- `todo_id`: integer - The ID of the todo item

**Response Example**:
```json
{
  "code": 200,
  "message": "Todo status toggled successfully",
  "data": {
    "id": 1,
    "completed": true,
    "updated_at": "2024-01-01T13:00:00Z"
  }
}
```

#### 6. Bulk Delete Completed Items
```http
DELETE /api/v1/todos/completed
```

**Response Example**:
```json
{
  "code": 200,
  "message": "Completed todos deleted successfully",
  "deleted_count": 5
}
```

#### 7. Clear All Todos
```http
DELETE /api/v1/todos/all
```

**Response Example**:
```json
{
  "code": 200,
  "message": "All todos deleted successfully",
  "deleted_count": 10
}
```

### Error Response Format

```json
{
  "code": 404,
  "message": "Todo not found",
  "detail": "Todo with id 999 does not exist"
}
```

## Testing Strategy

1.  **Unit Testing**
2.  **Integration Testing**: API endpoint tests
3.  **E2E Testing**
4.  **Performance Testing**: Testing with large data scenarios