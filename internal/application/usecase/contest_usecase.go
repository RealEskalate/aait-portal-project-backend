package usecase

import (
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/contracts/repository"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
)

type ContestUsecaseImpl struct {
	Repo repository.ContestRepository
}

func NewContestUsecase(repo repository.ContestRepository) *ContestUsecaseImpl {
	return &ContestUsecaseImpl{Repo: repo}
}

func (usecase *ContestUsecaseImpl) GetAllContests() ([]entities.Contest, error) {
	return usecase.Repo.GetAllContests()
}

func (usecase *ContestUsecaseImpl) CreateContest(contest *entities.Contest) error {
	return usecase.Repo.CreateContest(contest)
}
