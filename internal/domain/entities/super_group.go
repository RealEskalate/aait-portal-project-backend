package entities

import "time"

type SuperGroup struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Contests []Contest `gorm:"foreignKey:SuperGroupID" json:"contests"`
	Tracks   []Track   `gorm:"foreignKey:SuperGroupID" json:"tracks"`
}
