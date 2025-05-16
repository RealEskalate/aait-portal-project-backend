package repository

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := ur.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) UpdatePassword(userID string, hashedPassword string) error {
	return ur.DB.Model(&entities.User{}).
		Where("id = ?", userID).
		Update("password", hashedPassword).Error
}
