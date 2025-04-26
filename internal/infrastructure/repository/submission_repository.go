package repository

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"gorm.io/gorm"
)

type SubmissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) *SubmissionRepository {
	return &SubmissionRepository{db: db}
}
func (r *SubmissionRepository) GetTotalSolutions(userID int) (int, error) {
	var count int64
	err := r.db.Model(&entities.Submission{}).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
func (r *SubmissionRepository) GetTotalTimeSpents(userID int) (int, error) {
	var totalTimeSpent int64
	err := r.db.Model(&entities.Submission{}).Where("user_id = ?", userID).Select("COALESCE(SUM(time_spent), 0)").Scan(&totalTimeSpent).Error
	if err != nil {
		return 0, err
	}
	return int(totalTimeSpent), nil
}
