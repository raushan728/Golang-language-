package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Root URL "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Home Page!")
	})

	// /about route
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the About Page.")
	})

	// /hello?name=Raushan route with query param
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		fmt.Fprintf(w, "Hello, %s! Welcome to Go web server.\n", name)
	})

	// Start the server on port 8080
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
