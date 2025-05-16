package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Elizabethyonas/A2SV-Portal-Project/common/utilities"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure/repository"
)

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func ForgotPasswordHandler(
	userRepo *repository.UserRepository,
	tokenRepo *repository.ResetTokenRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ForgotPasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := userRepo.GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		token := utilities.GenerateSecureToken()
		expiration := time.Now().Add(15 * time.Minute)

		tokenRepo.Save(token, fmt.Sprintf("%d", user.ID), expiration)

		resetLink := "http://localhost:8080/reset-password?token=" + token

		subject := "Reset Your Password"
		body := "Click the link to reset your password: " + resetLink

		err = utilities.SendEmail(req.Email, subject, body)
		if err != nil {
			http.Error(w, "Failed to send email", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Password reset email sent."))
	}
}

func ResetPasswordHandler(
	userRepo *repository.UserRepository,
	tokenRepo *repository.ResetTokenRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ResetPasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Token == "" || req.NewPassword == "" {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		tokenData, exists := tokenRepo.Get(req.Token)
		if !exists || time.Now().After(tokenData.ExpiresAt) {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		hashedPassword, err := utilities.HashPassword(req.NewPassword)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		err = userRepo.UpdatePassword(tokenData.UserID, hashedPassword)
		if err != nil {
			http.Error(w, "Failed to update password", http.StatusInternalServerError)
			return
		}

		tokenRepo.Delete(req.Token)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Password successfully reset."))
	}
}
