package usecase

import "github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"

type UserUsecase interface {
	SignUp(user *entities.User) error
	SignIn(user *entities.UserLoginPayload) (string, error)
}
