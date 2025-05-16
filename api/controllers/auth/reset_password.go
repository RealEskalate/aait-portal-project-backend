package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure/repository"
	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func ResetPasswordHandler(userRepo *repository.UserRepository, tokenRepo *repository.ResetTokenRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ResetPasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Token == "" || req.NewPassword == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		tokenData, exists := tokenRepo.Get(req.Token)
		if !exists || time.Now().After(tokenData.ExpiresAt) {
			http.Error(w, "Token invalid or expired", http.StatusUnauthorized)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		err = userRepo.UpdatePassword(tokenData.UserID, string(hashedPassword))
		if err != nil {
			http.Error(w, "Failed to update password", http.StatusInternalServerError)
			return
		}

		tokenRepo.Delete(req.Token)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Password reset successful"))
	}
}
