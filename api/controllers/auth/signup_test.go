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
	"github.com/stretchr/testify/mock"
)

// MockUserUsecase mocks the UserUsecase interface
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) SignUp(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestSignUpHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Prepare mock
	mockUsecase := new(MockUserUsecase)

	testUser := entities.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "test123",
		Role:     "student",
	}

	mockUsecase.On("SignUp", mock.AnythingOfType("*entities.User")).Return(nil)

	// Setup router and controller
	router := gin.Default()
	authController := auth.NewAuthController(mockUsecase)
	router.POST("/signup", authController.SignUp)

	// Prepare request
	body, _ := json.Marshal(testUser)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User signed up successfully")
	mockUsecase.AssertExpectations(t)
}
