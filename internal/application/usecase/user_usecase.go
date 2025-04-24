package usecase

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	UserRepo *repository.UserRepository
}

func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{UserRepo: userRepo}
}

func (uc *UserUsecaseImpl) SignUp(user *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return uc.UserRepo.Create(user)
}
