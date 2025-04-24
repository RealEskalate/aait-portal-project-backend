package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/auth"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

func TestSignUpHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testUser := entities.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "test123",
		Role:     "student",
	}

	// Setup router and controller
	router := gin.Default()
	authController := auth.NewAuthController(nil) // Pass nil since we're not using mocks
	router.POST("/signup", authController.SignUp)

	// Prepare request
	body, _ := json.Marshal(testUser)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}
