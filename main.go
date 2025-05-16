package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/auth"
	database "github.com/Elizabethyonas/A2SV-Portal-Project/cmd/infrastructure/database"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure/repository"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Printf("Failed to connect to database: %s", err.Error())
		return
	}

	err = database.Migrate(db)
	if err != nil {
		log.Printf("Failed to migrate database: %s", err.Error())
		return
	}

	userRepo := repository.NewUserRepository(db)
	tokenRepo := repository.NewResetTokenRepository()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from your Go backend!")
	})

	http.HandleFunc("/forgot-password", auth.ForgotPasswordHandler(userRepo, tokenRepo))
	http.HandleFunc("/reset-password", auth.ResetPasswordHandler(userRepo, tokenRepo))

	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
