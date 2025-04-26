package submission

import (
	"strconv"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/contracts/usecase"
	"github.com/gin-gonic/gin"
)

type SubmissionController struct {
	SubmissionUsecase usecase.SubmissionUsecase
}

func NewSubmissionController(submissionUsecase usecase.SubmissionUsecase) *SubmissionController {
	return &SubmissionController{
		SubmissionUsecase: submissionUsecase,
	}
}

func (s *SubmissionController) GetTotalSolutions(ctx *gin.Context) {
	userIDStr, exists := ctx.Params.Get("user_id")
	if !exists {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "User ID must be an integer"})
		return
	}

	totalSolutions, err := s.SubmissionUsecase.GetTotalSolutions(userID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get total solutions"})
		return
	}

	ctx.JSON(200, gin.H{"total_solutions": totalSolutions})
}
func (s *SubmissionController) GetTotalTimeSpents(ctx *gin.Context) {
	userIDStr, exists := ctx.Params.Get("user_id")
	if !exists {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "User ID must be an integer"})
		return
	}

	totalTimeSpents, err := s.SubmissionUsecase.GetTotalTimeSpents(userID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get total time spent"})
		return
	}

	ctx.JSON(200, gin.H{"total_time_spent": totalTimeSpents})
}
