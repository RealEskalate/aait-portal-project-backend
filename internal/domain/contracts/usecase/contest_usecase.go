package usecase

import "github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"

type ContestUsecase interface {
	GetAllContests() ([]entities.Contest, error)
	CreateContest(contest *entities.Contest) error
}
