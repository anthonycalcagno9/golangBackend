package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func otherEndpointHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "other endpoint")
}

func main() {
	// Set up the routes
	http.HandleFunc("/hello", helloWorldHandler)
	http.HandleFunc("/other", otherEndpointHandler)

	// Start the server
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
