package main

import (
	"fmt"
	"net/http"
)

// HelloHandler handles all incoming HTTP requests
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with "Hello, all!" to any HTTP request
	fmt.Fprintln(w, "Hello, all!")
}

// Server
func main() {
	// Register the HelloHandler for all HTTP requests
	http.HandleFunc("/", HelloHandler)

	// Start the server on port 8080
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
