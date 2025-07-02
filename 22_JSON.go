package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	// 🔹 Step 1: JSON string → Struct
	jsonData := `{"name": "Raushan", "email": "raushan@example.com", "age": 22}`

	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error while parsing JSON:", err)
		return
	}

	fmt.Println("JSON → Struct:")
	fmt.Println("Name :", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Age  :", user.Age)

	// 🔹 Step 2: Struct → JSON
	jsonOutput, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error while encoding to JSON:", err)
		return
	}

	fmt.Println("\nStruct → JSON:")
	fmt.Println(string(jsonOutput))
}
