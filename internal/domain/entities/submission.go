package entities

import (
	"time"
)

type Submission struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	ProblemID int
	UserID    int
	TimeSpent int
	Tries     int
	InContest int
	Code      string `gorm:"type:text"`
	Language  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Verified  bool

	Problem Problem `gorm:"foreignKey:ProblemID"`
	User    User    `gorm:"foreignKey:UserID"`
}
