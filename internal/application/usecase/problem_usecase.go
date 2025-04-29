package usecase

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/contracts/repository"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
)

type ProblemUsecase struct {
	ProblemRepo repository.ProblemRepository
}

func NewProblemUsecase(problemRepo repository.ProblemRepository) *ProblemUsecase {
	return &ProblemUsecase{
		ProblemRepo: problemRepo,
	}
}

func (usecase *ProblemUsecase) CreateProblem(newProblem *entities.Problem) error {
	return usecase.ProblemRepo.CreateProblem(newProblem)
}

func (usecase *ProblemUsecase) GetProblemByID(problemID uint) (*entities.Problem, error) {
	return usecase.ProblemRepo.GetProblemByID(problemID)
}

func (usecase *ProblemUsecase) GetAllProblems() ([]entities.Problem, error) {
	return usecase.ProblemRepo.GetAllProblems()
}

func (usecase *ProblemUsecase) UpdateProblem(updatedProblem *entities.Problem) error {
	return usecase.ProblemRepo.UpdateProblem(updatedProblem)
}

func (usecase *ProblemUsecase) DeleteProblem(problemID uint) error {
	return usecase.ProblemRepo.DeleteProblem(problemID)
}

func (usecase *ProblemUsecase) GetProblemsByDifficulty(difficultyLevel string) ([]entities.Problem, error) {
	return usecase.ProblemRepo.GetProblemsByDifficulty(difficultyLevel)
}

func (usecase *ProblemUsecase) GetProblemsByTrack(trackName string) ([]entities.Problem, error) {
	return usecase.ProblemRepo.GetProblemsByTrack(trackName)
}
