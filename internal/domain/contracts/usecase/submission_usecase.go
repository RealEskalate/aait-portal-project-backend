package usecase

type SubmissionUsecase interface {
	GetTotalSolutions(userID int) (int, error)
	GetTotalTimeSpents(userID int) (int, error)
}
