package usecase

import "github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"

type UserProfileUsecase interface {
	GetProfile() (*entities.UserProfile, error)
	CreateProfile(profile *entities.UserProfile) error
	UpdateProfile(profile *entities.UserProfile) error
}
