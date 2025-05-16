package entities

import (
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	Title       string    `json:"title"`
	PostedAt    time.Time `json:"posted_at"`
	NumProblems int       `json:"num_problems"`
}
