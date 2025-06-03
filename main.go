package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// handleFileUpload handles JPEG file uploads
func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to image/jpeg
	w.Header().Set("Content-Type", "image/jpeg")

	// Parse the multipart form (10MB max upload size)
	const maxUploadSize = 10 * 1024 * 1024 // 10MB in bytes
	if err := r.ParseMultipartForm(int64(maxUploadSize)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Failed to parse form data")
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: No file uploaded")
		return
	}
	defer file.Close()

	// Verify file type
	if handler.Header.Get("Content-Type") != "image/jpeg" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: Only JPEG files are allowed")
		return
	}

	// Copy the file contents directly to the response
	if _, err := io.Copy(w, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: Failed to process file")
		return
	}
}

func main() {
	// Set up the routes
	http.HandleFunc("/hello", helloWorldHandler)
	http.HandleFunc("/upload", handleFileUpload)

	// Start the server
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
