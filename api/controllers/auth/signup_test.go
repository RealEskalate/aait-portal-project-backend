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
	"github.com/stretchr/testify/assert"
)

func TestSignUpHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testUser := entities.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "test123",
		Role:     "student",
	}

	router := gin.Default()
	authController := auth.NewAuthController(nil)
	router.POST("/signup", authController.SignUp)

	body, _ := json.Marshal(testUser)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User signed up successfully")
}
