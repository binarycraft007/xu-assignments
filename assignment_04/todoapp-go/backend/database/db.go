package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Construct connection string from spec-req.md
	// Username: monty, Password: test001, Address: localhost:3306, Database Name: todoapp
	dsn := "monty:test001@tcp(localhost:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local"
	
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Ping the database to verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
	createTables()
}

func createTables() {
	createTodoTableSQL := `
	CREATE TABLE IF NOT EXISTS todos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		priority INT NOT NULL DEFAULT 0,
		due_date DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createTodoTableSQL)
	if err != nil {
		log.Fatalf("Error creating todos table: %v", err)
	}
	fmt.Println("Todos table created or already exists.")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database connection closed.")
	}
}
