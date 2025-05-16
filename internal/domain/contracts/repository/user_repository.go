package repository

import "github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
}
