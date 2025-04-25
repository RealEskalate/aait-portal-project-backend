package entities

// UserProfile represents a user's profile information
// @Description User profile information
type UserProfile struct {
	// Standard gorm model fields
	ID        uint   `json:"id" gorm:"primarykey"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty" gorm:"index"`

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
