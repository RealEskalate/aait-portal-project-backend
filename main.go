package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Basic route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "🚀 Hello from your Go backend!")
	})

	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
