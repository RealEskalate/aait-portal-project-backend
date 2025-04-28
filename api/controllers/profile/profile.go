package profile

import (
	"fmt"
	"net/http"
	"os"

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
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var userID uint
	fmt.Sscanf(id, "%d", &userID)

	profile, err := p.Usecase.GetProfileByID(userID)
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

func (p *ProfileController) UploadProfilePicture(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	file, err := ctx.FormFile("profile_picture")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "No file is received",
			"details": "Make sure to send the file with key 'profile_picture' in form-data",
		})
		return
	}

	if err := os.MkdirAll("uploads", 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create uploads directory",
			"details": err.Error(),
		})
		return
	}

	filename := fmt.Sprintf("%s_%s", id, file.Filename)
	path := "uploads/" + filename

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to save file",
			"details": err.Error(),
		})
		return
	}

	// Convert ID to uint
	var userID uint
	fmt.Sscanf(id, "%d", &userID)

	// Update profile picture in database
	if err := p.Usecase.UpdateProfilePicture(userID, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update profile picture",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Profile picture uploaded successfully",
		"url":     fmt.Sprintf("/uploads/%s", filename),
	})
}

func (p *ProfileController) SetOnlineStatus(ctx *gin.Context) {
	var input struct {
		ID       uint `json:"id"`
		IsOnline bool `json:"is_online"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.Usecase.UpdateOnlineStatus(input.ID, input.IsOnline); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update online status"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Online status updated"})
}

func (p *ProfileController) DeleteProfilePicture(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var userID uint
	fmt.Sscanf(id, "%d", &userID)
	if err := p.Usecase.DeleteProfilePicture(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete profile picture"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Profile picture deleted successfully"})
}

func (p *ProfileController) GetOnlineStatus(ctx *gin.Context) {
	var input struct {
		ID uint `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := p.Usecase.GetOnlineStatus(input.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve online status"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": input.ID, "is_online": status})
}
