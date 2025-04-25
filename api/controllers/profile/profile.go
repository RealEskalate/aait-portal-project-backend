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

// GetProfile retrieves the user's profile.
//
// @Summary      Get Profile
// @Description  Fetches the profile of the currently authenticated user.
// @Tags         Profile
// @Produce      json
// @Success      200   {object}  entities.UserProfile
// @Failure      500   {object}  map[string]string
// @Router       /api/v1/profile [get]
func (p *ProfileController) GetProfile(ctx *gin.Context) {
	profile, err := p.Usecase.GetProfile()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch profile"})
		return
	}
	ctx.JSON(http.StatusOK, profile)
}

// CreateProfile creates a new user profile.
//
// @Summary      Create Profile
// @Description  Creates a new profile for the user with the provided details.
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        profile  body      entities.UserProfile  true  "Profile details"
// @Success      201      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router      /api/v1/profile [post]
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

// UpdateProfile updates an existing user profile.
//
// @Summary      Update Profile
// @Description  Updates the profile of the user with the provided details.
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        profile  body      entities.UserProfile  true  "Updated profile details"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /api/v1/profile [put]
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
