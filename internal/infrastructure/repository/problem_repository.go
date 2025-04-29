package repository

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"gorm.io/gorm"
)

type ProblemRepository struct {
	DB *gorm.DB
}

func NewProblemRepository(db *gorm.DB) *ProblemRepository {
	return &ProblemRepository{DB: db}
}

func (repo *ProblemRepository) CreateProblem(newProblem *entities.Problem) error {
	return repo.DB.Create(newProblem).Error
}

func (repo *ProblemRepository) GetProblemByID(problemID uint) (*entities.Problem, error) {
	var problem entities.Problem
	err := repo.DB.First(&problem, problemID).Error
	return &problem, err
}

func (repo *ProblemRepository) GetAllProblems() ([]entities.Problem, error) {
	var problems []entities.Problem
	err := repo.DB.Find(&problems).Error
	return problems, err
}

func (repo *ProblemRepository) UpdateProblem(updatedProblem *entities.Problem) error {
	return repo.DB.Save(updatedProblem).Error
}

func (repo *ProblemRepository) DeleteProblem(problemID uint) error {
	return repo.DB.Delete(&entities.Problem{}, problemID).Error
}

func (repo *ProblemRepository) GetProblemsByDifficulty(difficultyLevel string) ([]entities.Problem, error) {
	var problems []entities.Problem
	err := repo.DB.Where("difficulty = ?", difficultyLevel).Find(&problems).Error
	return problems, err
}

func (repo *ProblemRepository) GetProblemsByTrack(trackName string) ([]entities.Problem, error) {
	var problems []entities.Problem
	err := repo.DB.Where("track = ?", trackName).Find(&problems).Error
	return problems, err
}
