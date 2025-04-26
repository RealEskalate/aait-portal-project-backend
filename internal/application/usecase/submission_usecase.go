package usecase

import "github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure/repository"

type SubmissionUsecaseImpl struct {
	Repo *repository.SubmissionRepository
}

func NewSubmissionUsecase(repo *repository.SubmissionRepository) *SubmissionUsecaseImpl {
	return &SubmissionUsecaseImpl{Repo: repo}
}
func (u *SubmissionUsecaseImpl) GetTotalSolutions(userID int) (int, error) {
	return u.Repo.GetTotalSolutions(userID)
}
func (u *SubmissionUsecaseImpl) GetTotalTimeSpents(userID int) (int, error) {
	return u.Repo.GetTotalTimeSpents(userID)
}
