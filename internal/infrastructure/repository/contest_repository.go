package repository

import (
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
	return contests, err
}

func (repo *ContestRepositoryImpl) CreateContest(contest *entities.Contest) error {
	return repo.DB.Create(contest).Error
}
