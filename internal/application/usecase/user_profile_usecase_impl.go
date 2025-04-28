package usecase

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure/repository"
)

type UserProfileUsecaseImpl struct {
	Repo *repository.UserProfileRepository
}

func NewUserProfileUsecase(repo *repository.UserProfileRepository) *UserProfileUsecaseImpl {
	return &UserProfileUsecaseImpl{Repo: repo}
}

func (u *UserProfileUsecaseImpl) GetProfile() (*entities.UserProfile, error) {
	return u.Repo.GetProfile()
}

func (u *UserProfileUsecaseImpl) CreateProfile(profile *entities.UserProfile) error {
	return u.Repo.CreateProfile(profile)
}

func (u *UserProfileUsecaseImpl) UpdateProfile(profile *entities.UserProfile) error {
	return u.Repo.UpdateProfile(profile)
}

func (u *UserProfileUsecaseImpl) UpdateProfilePicture(userID uint, path string) error {
	return u.Repo.UpdateProfilePicture(userID, path)
}

func (u *UserProfileUsecaseImpl) UpdateOnlineStatus(userID uint, status bool) error {
	return u.Repo.UpdateOnlineStatus(userID, status)
}

func (u *UserProfileUsecaseImpl) DeleteProfilePicture(userID uint) error {
	return u.Repo.DeleteProfilePicture(userID)
}

func (u *UserProfileUsecaseImpl) GetOnlineStatus(userID uint) (bool, error) {
	return u.Repo.GetOnlineStatus(userID)
}

func (u *UserProfileUsecaseImpl) GetProfileByID(id uint) (*entities.UserProfile, error) {
	return u.Repo.GetProfileByID(id)
}
