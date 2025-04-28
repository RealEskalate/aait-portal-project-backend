package usecase

import "github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"

type UserProfileUsecase interface {
	GetProfile() (*entities.UserProfile, error)
	GetProfileByID(id uint) (*entities.UserProfile, error)
	CreateProfile(profile *entities.UserProfile) error
	UpdateProfile(profile *entities.UserProfile) error
	UpdateProfilePicture(userID uint, path string) error
	UpdateOnlineStatus(userID uint, status bool) error
	DeleteProfilePicture(userID uint) error
	GetOnlineStatus(userID uint) (bool, error)
}
