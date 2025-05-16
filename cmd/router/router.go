package router

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/auth"
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/contest"
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/problem"
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/profile"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/application/usecase"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	authController := auth.NewAuthController(userUsecase)

	profileRepo := repository.NewUserProfileRepository(db)
	profileUsecase := usecase.NewUserProfileUsecase(profileRepo)
	profileController := profile.NewProfileController(profileUsecase)

	problemRepo := repository.NewProblemRepository(db)
	problemUsecase := usecase.NewProblemUsecase(problemRepo)
	problemController := problem.NewProblemController(problemUsecase)

	contestRepo := repository.NewContestRepository(db)
	contestUsecase := usecase.NewContestUsecase(contestRepo)
	contestController := contest.NewContestController(contestUsecase)

	// Auth routes
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/signup", authController.SignUp)
		authGroup.POST("/signin", authController.SignIn)
	}

	// User routes
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/", authController.GetAllUsers)
	}

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

	problemGroup := router.Group("/api/problems")
	{
		problemGroup.POST("/", problemController.CreateProblem)
		problemGroup.GET("/", problemController.GetAllProblems)
		problemGroup.GET("/:id", problemController.GetProblem)
		problemGroup.PUT("/:id", problemController.UpdateProblem)
		problemGroup.DELETE("/:id", problemController.DeleteProblem)
		problemGroup.GET("/difficulty/:difficulty", problemController.GetProblemsByDifficulty)
		problemGroup.GET("/track/:track", problemController.GetProblemsByTrack)
	}

	contestGroup := router.Group("/api/contests")
	{
		contestGroup.GET("/", contestController.GetAllContests)
		contestGroup.POST("/", contestController.CreateContest)
	}
}
