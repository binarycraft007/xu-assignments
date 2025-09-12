package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files from the current directory
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}