package entities

import "time"

type Problem struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	ContestID  int       `gorm:"foreignKey:ContestID" json:"contest_id"`
	TrackID    int       `gorm:"foreignKey:TrackID" json:"track_id"`
	Name       string    `json:"name"`
	Difficulty string    `json:"difficulty"`
	Tag        string    `json:"tag"`
	Platform   string    `json:"platform"`
	Link       string    `json:"link"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Contest Contest `gorm:"foreignKey:ContestID" json:"contest"`
	Track   Track   `gorm:"foreignKey:TrackID" json:"track"`
}
