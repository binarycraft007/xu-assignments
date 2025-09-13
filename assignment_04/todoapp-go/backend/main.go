package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"todoapp-go/backend/database"
	"todoapp-go/backend/handlers"

	// For Swagger documentation
	_ "todoapp-go/backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todo Application API
// @version 1.0
// @description A RESTful to-do item management API.
// @host localhost:8000
// @BasePath /api/v1
// @schemes http
func main() {
	// Initialize database connection
	database.InitDB()
	defer database.CloseDB()

	// Set Gin to production mode in a real application
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Allow your frontend origin
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", handlers.GetTodos)
		v1.POST("/todos", handlers.CreateTodo)
		v1.PUT("/todos/:todo_id", handlers.UpdateTodo)
		v1.DELETE("/todos/:todo_id", handlers.DeleteTodo)
		v1.PATCH("/todos/:todo_id/toggle", handlers.ToggleTodoStatus)
		v1.DELETE("/todos/completed", handlers.DeleteCompletedTodos)
		v1.DELETE("/todos/all", handlers.ClearAllTodos)
	}

	log.Println("Server starting on :8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
