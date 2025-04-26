package entities

import (
	"time"

	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Title      string    `json:"title"`
	Tags       []string  `json:"tags" gorm:"type:text[]"`
	Difficulty string    `json:"difficulty"` // Easy, Medium, Hard
	Track      string    `json:"track"`      // e.g., "Data Structures", "Algorithms"
	Link       string    `json:"link"`
	PostedAt   time.Time `json:"posted_at"`
}
