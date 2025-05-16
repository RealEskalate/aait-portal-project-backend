package usecase

import (
	"errors"
	"log"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/contracts/repository"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
)

var (
	ErrInvalidContestTitle = errors.New("contest title cannot be empty")
	ErrInvalidNumProblems  = errors.New("number of problems must be greater than 0")
)

type ContestUsecaseImpl struct {
	Repo repository.ContestRepository
}

func NewContestUsecase(repo repository.ContestRepository) *ContestUsecaseImpl {
	return &ContestUsecaseImpl{Repo: repo}
}

func (usecase *ContestUsecaseImpl) GetAllContests() ([]entities.Contest, error) {
	contests, err := usecase.Repo.GetAllContests()
	if err != nil {
		log.Printf("Error fetching contests: %v", err)
		return nil, err
	}
	return contests, nil
}

func (usecase *ContestUsecaseImpl) CreateContest(contest *entities.Contest) error {
	if contest.Title == "" {
		log.Printf("Validation error: empty contest title")
		return ErrInvalidContestTitle
	}
	if contest.NumProblems <= 0 {
		log.Printf("Validation error: invalid number of problems: %d", contest.NumProblems)
		return ErrInvalidNumProblems
	}

	err := usecase.Repo.CreateContest(contest)
	if err != nil {
		log.Printf("Error creating contest: %v", err)
		return err
	}

	return nil
}
