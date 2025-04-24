package repository

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"gorm.io/gorm"
)

type UserProfileRepository struct {
	DB *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) *UserProfileRepository {
	return &UserProfileRepository{DB: db}
}

func (r *UserProfileRepository) GetProfile() (*entities.UserProfile, error) {
	var profile entities.UserProfile
	err := r.DB.First(&profile).Error
	return &profile, err
}

func (r *UserProfileRepository) CreateProfile(profile *entities.UserProfile) error {
	return r.DB.Create(profile).Error
}

func (r *UserProfileRepository) UpdateProfile(profile *entities.UserProfile) error {
	return r.DB.Save(profile).Error
}
