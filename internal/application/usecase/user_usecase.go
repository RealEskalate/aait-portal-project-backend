package usecase

import (
	jwt "github.com/Elizabethyonas/A2SV-Portal-Project/common/utilities"
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

func (uc *UserUsecaseImpl) SignIn(user *entities.UserLoginPayload) (string, error) {
	existingUser, err := uc.UserRepo.FindByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token, err := jwt.CreateToken(existingUser.ID, existingUser.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
