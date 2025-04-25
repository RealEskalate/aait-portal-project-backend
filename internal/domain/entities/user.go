package entities

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name" gorm:"column:name"`
	Email    string    `json:"email" gorm:"column:email;unique"`
	Password string    `json:"password" gorm:"column:password"`
	Role     string    `json:"role" gorm:"column:role"`
}

type UserLoginPayload struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}
