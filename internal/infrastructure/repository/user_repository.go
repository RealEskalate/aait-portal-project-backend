package repository

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *entities.User) error {
	return r.DB.Create(user).Error
}
