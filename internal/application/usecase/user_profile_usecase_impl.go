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
