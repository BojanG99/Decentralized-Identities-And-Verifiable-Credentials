package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function that writes "Hello, World!" to the response.
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	// Register the handler function for the "/" route.
	http.HandleFunc("/", helloHandler)

	// Start the server on port 7000.
	fmt.Println("Server is listening on :7000...")
	err := http.ListenAndServe("0.0.0.0:7000", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
