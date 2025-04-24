package entities

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model

	Location    string `json:"location"`
	Email       string `json:"email"`
	TechStack   string `json:"tech_stack"`
	Group       string `json:"group"`
	Institution string `json:"institution"`

	LinkedIn   string `json:"linkedin"`
	Telegram   string `json:"telegram"`
	LeetCode   string `json:"leetcode"`
	Codeforces string `json:"codeforces"`
	HackerRank string `json:"hackerrank"`
	GitHub     string `json:"github"`
}
