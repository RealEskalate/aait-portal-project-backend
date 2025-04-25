package router

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/auth"
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/profile"
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/swagger"
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
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/signup", authController.SignUp)
	}

	// Profile routes
	apiGroup := router.Group("/api/v1/profile")
	{
		apiGroup.GET("/", profileController.GetProfile)
		apiGroup.POST("/", profileController.CreateProfile)
		apiGroup.PUT("/", profileController.UpdateProfile)
	}

	// Swagger routes
	swagger.SetupSwaggerRoutes(router)
}
