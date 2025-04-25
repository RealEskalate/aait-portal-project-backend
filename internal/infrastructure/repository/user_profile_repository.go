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

func (repo *UserProfileRepository) GetProfile() (*entities.UserProfile, error) {
	var profile entities.UserProfile
	err := repo.DB.First(&profile).Error
	return &profile, err
}

func (repo *UserProfileRepository) CreateProfile(profile *entities.UserProfile) error {
	return repo.DB.Create(profile).Error
}

func (repo *UserProfileRepository) UpdateProfile(profile *entities.UserProfile) error {
	return repo.DB.Save(profile).Error
}

func (repo *UserProfileRepository) UpdateProfilePicture(userID uint, path string) error {
	return repo.DB.Model(&entities.UserProfile{}).Where("id = ?", userID).Update("profile_picture_url", path).Error
}

func (repo *UserProfileRepository) UpdateOnlineStatus(userID uint, status bool) error {
	return repo.DB.Model(&entities.UserProfile{}).Where("id = ?", userID).Update("is_online", status).Error
}

func (repo *UserProfileRepository) DeleteProfilePicture(userID uint) error {
	return repo.DB.Model(&entities.UserProfile{}).Where("id = ?", userID).Update("profile_picture", "").Error
}

func (repo *UserProfileRepository) GetOnlineStatus(userID uint) (bool, error) {
	var user entities.UserProfile
	err := repo.DB.Select("is_online").Where("id = ?", userID).First(&user).Error
	return user.IsOnline, err
}

func (repo *UserProfileRepository) GetProfileByID(id uint) (*entities.UserProfile, error) {
	var profile entities.UserProfile
	err := repo.DB.Where("id = ?", id).First(&profile).Error
	return &profile, err
}
