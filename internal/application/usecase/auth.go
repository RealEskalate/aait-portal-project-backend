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

func ForgotPasswordHandler(
	userRepo *repository.UserRepository,
	tokenRepo *repository.ResetTokenRepository,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ForgotPasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Email == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
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

		resetLink := "https://your-frontend/reset-password?token=" + token
		body := "Click the link to reset your password: " + resetLink
		if err := utilities.SendEmail(req.Email, "Password Reset", body); err != nil {
			http.Error(w, "Failed to send email", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reset email sent"))
	}
}
