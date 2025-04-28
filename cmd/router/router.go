package router

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/auth"
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/profile"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/application/usecase"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Setup Auth Controller
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	authController := auth.NewAuthController(userUsecase)

	// Setup Profile Controller
	profileRepo := repository.NewUserProfileRepository(db)
	profileUsecase := usecase.NewUserProfileUsecase(profileRepo)
	profileController := profile.NewProfileController(profileUsecase)

	// Auth routes
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/signup", authController.SignUp)
		authGroup.POST("/signin", authController.SignIn)
	}

	// Profile routes
	profileGroup := router.Group("/api/profile")
	{
		profileGroup.GET("/", profileController.GetProfile)
		profileGroup.POST("/", profileController.CreateProfile)
		profileGroup.PUT("/", profileController.UpdateProfile)
		profileGroup.POST("/upload-picture", profileController.UploadProfilePicture)
		profileGroup.DELETE("/delete-picture", profileController.DeleteProfilePicture)
		profileGroup.POST("/set-status", profileController.SetOnlineStatus)
		profileGroup.GET("/get-status", profileController.GetOnlineStatus)
	}
}
