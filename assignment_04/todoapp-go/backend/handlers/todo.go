package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"todoapp-go/backend/models"
)

// APIResponse defines the structure for API responses
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
	DeletedCount int      `json:"deleted_count,omitempty"`
	Detail       string      `json:"detail,omitempty"`
}

// CreateTodo handles the creation of a new todo item
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	if err := models.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create todo",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, APIResponse{
		Code:    http.StatusCreated,
		Message: "Todo created successfully",
		Data:    todo,
	})
}

// GetTodos handles retrieving a list of todo items
func GetTodos(c *gin.Context) {
	completedStr := c.Query("completed")
	limitStr := c.DefaultQuery("limit", "100")
	offsetStr := c.DefaultQuery("offset", "0")

	var completed *bool
	if completedStr != "" {
		parsedCompleted, err := strconv.ParseBool(completedStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, APIResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid 'completed' parameter",
				Data:    err.Error(),
			})
			return
		}
		completed = &parsedCompleted
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid 'limit' parameter",
			Data:    err.Error(),
		})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid 'offset' parameter",
			Data:    err.Error(),
		})
		return
	}

	todos, total, err := models.GetTodos(completed, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todos",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    todos,
		Total:   total,
	})
}

// GetTodoByID handles retrieving a single todo item by ID
func GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid todo ID",
			Data:    err.Error(),
		})
		return
	}

	todo, err := models.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todo",
			Data:    err.Error(),
		})
		return
	}

	if todo == nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Code:    http.StatusNotFound,
			Message: "Todo not found",
			Detail:  "Todo with id " + c.Param("todo_id") + " does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    todo,
	})
}

// UpdateTodo handles updating an existing todo item
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid todo ID",
			Data:    err.Error(),
		})
		return
	}

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}
	todo.ID = id

	// Check if the todo exists before attempting to update
	existingTodo, err := models.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todo for update",
			Data:    err.Error(),
		})
		return
	}
	if existingTodo == nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Code:    http.StatusNotFound,
			Message: "Todo not found",
			Detail:  "Todo with id " + c.Param("todo_id") + " does not exist",
		})
		return
	}

	if err := models.UpdateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update todo",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: "Todo updated successfully",
		Data:    todo,
	})
}

// DeleteTodo handles deleting a todo item
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid todo ID",
			Data:    err.Error(),
		})
		return
	}

	// Check if the todo exists before attempting to delete
	existingTodo, err := models.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todo for deletion",
			Data:    err.Error(),
		})
		return
	}
	if existingTodo == nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Code:    http.StatusNotFound,
			Message: "Todo not found",
			Detail:  "Todo with id " + c.Param("todo_id") + " does not exist",
		})
		return
	}

	if err := models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete todo",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: "Todo deleted successfully",
	})
}

// ToggleTodoStatus handles toggling the completion status of a todo item
func ToggleTodoStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid todo ID",
			Data:    err.Error(),
		})
		return
	}

	// Check if the todo exists before attempting to toggle
	existingTodo, err := models.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todo for status toggle",
			Data:    err.Error(),
		})
		return
	}
	if existingTodo == nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Code:    http.StatusNotFound,
			Message: "Todo not found",
			Detail:  "Todo with id " + c.Param("todo_id") + " does not exist",
		})
		return
	}

	newStatus := !existingTodo.Completed
	if err := models.ToggleTodoStatus(id, newStatus); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to toggle todo status",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: "Todo status toggled successfully",
		Data:    gin.H{"id": id, "completed": newStatus},
	})
}

// DeleteCompletedTodos handles deleting all completed todo items
func DeleteCompletedTodos(c *gin.Context) {
	rowsAffected, err := models.DeleteCompletedTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete completed todos",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: "Completed todos deleted successfully",
		DeletedCount: int(rowsAffected),
	})
}

// ClearAllTodos handles deleting all todo items
func ClearAllTodos(c *gin.Context) {
	rowsAffected, err := models.ClearAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete all todos",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: "All todos deleted successfully",
		DeletedCount: int(rowsAffected),
	})
}
