package entities

import "time"

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string

	ResetPasswordToken  string    `gorm:"size:255" json:"reset_password_token"`
	ResetPasswordExpiry time.Time `json:"reset_password_expiry"`
}
