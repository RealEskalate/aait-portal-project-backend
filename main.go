package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/Elizabethyonas/A2SV-Portal-Project/cmd/infrastructure/database"
)

func main() {
	_,err := database.NewDatabase()
	if err != nil{
		log.Printf("Failed to connect to database: %s", err.Error())
		return
	}
	// Basic route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "🚀 Hello from your Go backend!")
	})

	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
