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

// In-memory "database"
var users = []User{
	{ID: 1, Name: "Raushan", Email: "raushan@example.com"},
	{ID: 2, Name: "Aryan", Email: "aryan@example.com"},
}

func main() {
	http.HandleFunc("/user/", deleteUserHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	//Find index of user
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

	// Delete from slice
	users = append(users[:index], users[index+1:]...)

	// Return confirmation
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("User with ID %d deleted successfully", id),
	})
}
