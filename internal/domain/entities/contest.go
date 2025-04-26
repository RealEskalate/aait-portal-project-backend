package entities

import "time"

type Contest struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name"`
	Link         string    `json:"link"`
	ProblemCount int       `json:"problem_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Unrated      bool      `gorm:"default:false" json:"unrated"`
	SuperGroupID int       `gorm:"foreignKey:SuperGroupID" json:"super_group_id"`
	Type         string    `json:"type"`
	Link2        string    `json:"link2"`
	Link3        string    `json:"link3"`

	SuperGroup SuperGroup `gorm:"foreignKey:SuperGroupID" json:"super_group"`
}
