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

func (r *UserProfileRepository) UpdateProfilePicture(userID uint, path string) error {
	return r.DB.Model(&entities.UserProfile{}).Where("id = ?", userID).Update("profile_picture_url", path).Error
}

func (r *UserProfileRepository) UpdateOnlineStatus(userID uint, status bool) error {
	return r.DB.Model(&entities.UserProfile{}).Where("id = ?", userID).Update("is_online", status).Error
}

func (r *UserProfileRepository) DeleteProfilePicture(userID uint) error {
	return r.DB.Model(&entities.UserProfile{}).Where("id = ?", userID).Update("profile_picture", "").Error
}

func (r *UserProfileRepository) GetOnlineStatus(userID uint) (bool, error) {
	var user entities.UserProfile
	err := r.DB.Select("is_online").Where("id = ?", userID).First(&user).Error
	return user.IsOnline, err
}

func (r *UserProfileRepository) GetProfileByID(id uint) (*entities.UserProfile, error) {
	var profile entities.UserProfile
	err := r.DB.Where("id = ?", id).First(&profile).Error
	return &profile, err
}
