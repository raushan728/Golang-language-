package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// User struct represents a user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-memory database
var users = []User{
	{ID: 1, Name: "Raushan", Email: "raushan@example.com"},
	{ID: 2, Name: "Aryan", Email: "aryan@example.com"},
}

func main() {
	http.HandleFunc("/user/", updateUserHandler) // expects /user/{id}

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	//Only allow PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Extract ID from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find user by ID
	index := -1
	for i, u := range users {
		if u.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode request body
	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update user fields (except ID)
	users[index].Name = updatedUser.Name
	users[index].Email = updatedUser.Email

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users[index])
}
