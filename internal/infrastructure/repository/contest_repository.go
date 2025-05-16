package repository

import (
	"log"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"gorm.io/gorm"
)

type ContestRepositoryImpl struct {
	DB *gorm.DB
}

func NewContestRepository(db *gorm.DB) *ContestRepositoryImpl {
	return &ContestRepositoryImpl{DB: db}
}

func (repo *ContestRepositoryImpl) GetAllContests() ([]entities.Contest, error) {
	var contests []entities.Contest
	err := repo.DB.Find(&contests).Error
	if err != nil {
		log.Printf("Database error while fetching contests: %v", err)
		return nil, err
	}
	return contests, nil
}

func (repo *ContestRepositoryImpl) CreateContest(contest *entities.Contest) error {
	result := repo.DB.Create(contest)
	if result.Error != nil {
		log.Printf("Database error while creating contest: %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		log.Printf("Warning: No rows affected when creating contest with title: %s", contest.Title)
	}
	return nil
}
