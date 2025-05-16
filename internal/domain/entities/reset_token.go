package entities

import "time"

type ResetToken struct {
	Token     string
	UserID    string
	ExpiresAt time.Time
}
