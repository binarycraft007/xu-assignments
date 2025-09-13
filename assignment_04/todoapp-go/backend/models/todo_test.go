package models

import (
	"database/sql"
	"testing"
	"time"
	"regexp"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"todoapp-go/backend/database"
)

func TestCreateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	todo := &Todo{
		Title:       "Test Todo",
		Description: sql.NullString{String: "Test Description", Valid: true},
		Priority:    1,
		DueDate:     sql.NullTime{Time: time.Now(), Valid: true},
	}

	mock.ExpectPrepare("INSERT INTO todos").ExpectExec().WithArgs(
		todo.Title, todo.Description, todo.Priority, todo.DueDate).WillReturnResult(sqlmock.NewResult(1, 1))

	err = CreateTodo(todo)
	assert.NoError(t, err)
	assert.Equal(t, 1, todo.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTodos(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	rows := sqlmock.NewRows([]string{"id", "title", "description", "completed", "priority", "due_date", "created_at", "updated_at"}).
		AddRow(1, "Test Todo 1", "Desc 1", false, 0, time.Now(), time.Now(), time.Now()).
		AddRow(2, "Test Todo 2", "Desc 2", true, 1, time.Now(), time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, title, description, completed, priority, due_date, created_at, updated_at FROM todos WHERE completed = ? LIMIT ? OFFSET ?")).
		WithArgs(false, 100, 0).WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM todos WHERE completed = ?")).WithArgs(false).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	completed := false
	todos, total, err := GetTodos(&completed, 100, 0)
	assert.NoError(t, err)
	assert.Len(t, todos, 2)
	assert.Equal(t, 1, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTodoByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	rows := sqlmock.NewRows([]string{"id", "title", "description", "completed", "priority", "due_date", "created_at", "updated_at"}).
		AddRow(1, "Test Todo", "Desc", false, 0, time.Now(), time.Now(), time.Now())

	mock.ExpectQuery("SELECT id, title, description, completed, priority, due_date, created_at, updated_at FROM todos WHERE id = ?").
		WithArgs(1).WillReturnRows(rows)

	todo, err := GetTodoByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, 1, todo.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	todo := &Todo{
		ID:          1,
		Title:       "Updated Todo",
		Description: sql.NullString{String: "Updated Desc", Valid: true},
		Completed:   true,
		Priority:    2,
		DueDate:     sql.NullTime{Time: time.Now(), Valid: true},
	}

	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE todos SET title = ?, description = ?, completed = ?, priority = ?, due_date = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?")).
		ExpectExec().WithArgs(todo.Title, todo.Description, todo.Completed, todo.Priority, todo.DueDate, todo.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err = UpdateTodo(todo)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	mock.ExpectPrepare("DELETE FROM todos WHERE id = ?").ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	err = DeleteTodo(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestToggleTodoStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE todos SET completed = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?")).ExpectExec().WithArgs(true, 1).WillReturnResult(sqlmock.NewResult(0, 1))

	err = ToggleTodoStatus(1, true)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteCompletedTodos(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	mock.ExpectExec("DELETE FROM todos WHERE completed = TRUE").WillReturnResult(sqlmock.NewResult(0, 5))

	rowsAffected, err := DeleteCompletedTodos()
	assert.NoError(t, err)
	assert.Equal(t, int64(5), rowsAffected)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestClearAllTodos(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	mock.ExpectExec("DELETE FROM todos").WillReturnResult(sqlmock.NewResult(0, 10))

	rowsAffected, err := ClearAllTodos()
	assert.NoError(t, err)
	assert.Equal(t, int64(10), rowsAffected)
	assert.NoError(t, mock.ExpectationsWereMet())
}
