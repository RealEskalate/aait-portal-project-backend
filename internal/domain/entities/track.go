package entities

import "time"

type Track struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Active       bool      `gorm:"default:true" json:"active"`
	SuperGroupID int       `gorm:"foreignKey:SuperGroupID" json:"super_group_id"`

	SuperGroup SuperGroup `gorm:"foreignKey:SuperGroupID" json:"super_group"`
	Problems   []Problem  `gorm:"foreignKey:TrackID" json:"problems"`
}
