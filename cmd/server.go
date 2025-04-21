package main

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/api/controllers/auth"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/application/usecase"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := infrastructure.NewDatabase()
	if err != nil {
		panic(err)
	}
	infrastructure.Migrate(db)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	authController := auth.NewAuthController(userUsecase)

	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", authController.SignUp)
	}

	r.Run(":8080")
}
