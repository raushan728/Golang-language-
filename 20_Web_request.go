package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Struct to hold GitHub user info
type GitHubUser struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	HTMLURL     string `json:"html_url"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
}

func main() {
	// Ask user for GitHub username
	fmt.Print("Enter GitHub username: ")
	var username string
	fmt.Scanln(&username)

	// Build the URL
	url := "https://api.github.com/users/" + username

	// Make GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("User-Agent", "Raushan-Go-Client")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	// Handle non-200 status
	if resp.StatusCode != http.StatusOK {
		fmt.Println("User not found or API limit reached!")
		return
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse JSON into struct
	var user GitHubUser
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print result
	fmt.Println("\n GitHub User Info:")
	fmt.Println("Username   :", user.Login)
	fmt.Println("ID         :", user.ID)
	fmt.Println("Profile URL:", user.HTMLURL)
	fmt.Println("Repos      :", user.PublicRepos)
	fmt.Println("Followers  :", user.Followers)
}
