package auth

import (
	"net/http"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/contracts/usecase"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserUsecase usecase.UserUsecase
}

func NewAuthController(userUsecase usecase.UserUsecase) *AuthController {
	return &AuthController{UserUsecase: userUsecase}
}

func (ctrl *AuthController) SignUp(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.UserUsecase.SignUp(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})
}
