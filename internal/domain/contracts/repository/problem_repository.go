package repository

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
)

type ProblemRepository interface {
	CreateProblem(problem *entities.Problem) error
	GetProblemByID(id uint) (*entities.Problem, error)
	GetAllProblems() ([]entities.Problem, error)
	UpdateProblem(problem *entities.Problem) error
	DeleteProblem(id uint) error
	GetProblemsByDifficulty(difficulty string) ([]entities.Problem, error)
	GetProblemsByTrack(track string) ([]entities.Problem, error)
}
