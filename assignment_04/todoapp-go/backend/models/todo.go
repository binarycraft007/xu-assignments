package models

import (
	"database/sql"
	"time"

	"todoapp-go/backend/database"
)

// Todo represents a single todo item
type Todo struct {
	ID          int            `json:"id,omitempty"`
	Title       string         `json:"title" binding:"required"`	
	Description string         `json:"description"`
	Completed   bool           `json:"completed"`
	Priority    int            `json:"priority"`
	DueDate     *time.Time     `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"`
}

// CreateTodo inserts a new todo item into the database
func CreateTodo(todo *Todo) error {
	var nullDescription sql.NullString
	if todo.Description != "" {
		nullDescription = sql.NullString{String: todo.Description, Valid: true}
	}

	var nullDueDate sql.NullTime
	if todo.DueDate != nil {
		nullDueDate = sql.NullTime{Time: *todo.DueDate, Valid: true}
	}

	stmt, err := database.DB.Prepare(
		"INSERT INTO todos (title, description, priority, due_date) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Title, nullDescription, todo.Priority, nullDueDate)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	todo.ID = int(id)
	return nil
}

// GetTodos retrieves a list of todo items from the database with optional filtering and pagination
func GetTodos(completed *bool, limit, offset int) ([]Todo, int, error) {
	query := "SELECT id, title, description, completed, priority, due_date, created_at, updated_at FROM todos"
	countQuery := "SELECT COUNT(*) FROM todos"
	args := []interface{}{}
	countArgs := []interface{}{}

	whereClauses := []string{}
	if completed != nil {
		whereClauses = append(whereClauses, "completed = ?")
		args = append(args, *completed)
		countArgs = append(countArgs, *completed)
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + joinStrings(whereClauses, " AND ")
		countQuery += " WHERE " + joinStrings(whereClauses, " AND ")
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		var nullDescription sql.NullString
		var nullDueDate sql.NullTime
		err := rows.Scan(&todo.ID, &todo.Title, &nullDescription, &todo.Completed, &todo.Priority, &nullDueDate, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}

		if nullDescription.Valid {
			todo.Description = nullDescription.String
		}
		if nullDueDate.Valid {
			todo.DueDate = &nullDueDate.Time
		}
		todos = append(todos, todo)
	}

	var total int
	err = database.DB.QueryRow(countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}

// GetTodoByID retrieves a single todo item by its ID
func GetTodoByID(id int) (*Todo, error) {
	var todo Todo
	var nullDescription sql.NullString
	var nullDueDate sql.NullTime
	err := database.DB.QueryRow(
		"SELECT id, title, description, completed, priority, due_date, created_at, updated_at FROM todos WHERE id = ?", id).Scan(
		&todo.ID, &todo.Title, &nullDescription, &todo.Completed, &todo.Priority, &nullDueDate, &todo.CreatedAt, &todo.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil // Todo not found
	}
	if err != nil {
		return nil, err
	}

	if nullDescription.Valid {
		todo.Description = nullDescription.String
	}
	if nullDueDate.Valid {
		todo.DueDate = &nullDueDate.Time
	}

	return &todo, nil
}

// UpdateTodo updates an existing todo item in the database
func UpdateTodo(todo *Todo) error {
	var nullDescription sql.NullString
	if todo.Description != "" {
		nullDescription = sql.NullString{String: todo.Description, Valid: true}
	}

	var nullDueDate sql.NullTime
	if todo.DueDate != nil {
		nullDueDate = sql.NullTime{Time: *todo.DueDate, Valid: true}
	}

	stmt, err := database.DB.Prepare(
		"UPDATE todos SET title = ?, description = ?, completed = ?, priority = ?, due_date = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, nullDescription, todo.Completed, todo.Priority, nullDueDate, todo.ID)
	return err
}

// DeleteTodo deletes a todo item from the database by its ID
func DeleteTodo(id int) error {
	stmt, err := database.DB.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

// ToggleTodoStatus updates the completion status of a todo item
func ToggleTodoStatus(id int, completed bool) error {
	stmt, err := database.DB.Prepare("UPDATE todos SET completed = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(completed, id)
	return err
}

// DeleteCompletedTodos deletes all completed todo items
func DeleteCompletedTodos() (int64, error) {
	result, err := database.DB.Exec("DELETE FROM todos WHERE completed = TRUE")
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// ClearAllTodos deletes all todo items from the database
func ClearAllTodos() (int64, error) {
	result, err := database.DB.Exec("DELETE FROM todos")
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func joinStrings(s []string, sep string) string {
	var result string
	for i, v := range s {
		result += v
		if i < len(s)-1 {
			result += sep
		}
	}
	return result
}
