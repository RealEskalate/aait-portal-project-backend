package repository

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *entities.User) error {
	user.ID = uuid.New()
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
