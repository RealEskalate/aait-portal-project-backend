package profile

import (
	"net/http"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/contracts/usecase"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	Usecase usecase.UserProfileUsecase
}

func NewProfileController(u usecase.UserProfileUsecase) *ProfileController {
	return &ProfileController{Usecase: u}
}

func (p *ProfileController) GetProfile(ctx *gin.Context) {
	profile, err := p.Usecase.GetProfile()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch profile"})
		return
	}
	ctx.JSON(http.StatusOK, profile)
}

func (p *ProfileController) CreateProfile(ctx *gin.Context) {
	var profile entities.UserProfile
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.Usecase.CreateProfile(&profile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create profile"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Profile created successfully"})
}

func (p *ProfileController) UpdateProfile(ctx *gin.Context) {
	var profile entities.UserProfile
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.Usecase.UpdateProfile(&profile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update profile"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
