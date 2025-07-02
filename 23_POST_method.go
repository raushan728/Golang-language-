// Filename: server.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to accept JSON input
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/user", createUserHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Print received data
	fmt.Println("New User Received:")
	fmt.Println("Name :", user.Name)
	fmt.Println("Email:", user.Email)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully!",
	})
}
