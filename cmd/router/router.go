package router

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/auth"
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/profile"
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/submission"
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
	apiGroup := router.Group("/api/profile")
	{
		apiGroup.GET("/", profileController.GetProfile)
		apiGroup.POST("/", profileController.CreateProfile)
		apiGroup.PUT("/", profileController.UpdateProfile)
	}

	// submission setup
	submissionRepo := repository.NewSubmissionRepository(db)
	submissionUsecase := usecase.NewSubmissionUsecase(submissionRepo)
	submissionController := submission.NewSubmissionController(submissionUsecase)

	// Submission routes
	submissionGroup := router.Group("/api/submissions")
	{
		submissionGroup.GET("/total-solutions/:user_id", submissionController.GetTotalSolutions)
		submissionGroup.GET("/total-time-spent/:user_id", submissionController.GetTotalTimeSpents)
	}
}
