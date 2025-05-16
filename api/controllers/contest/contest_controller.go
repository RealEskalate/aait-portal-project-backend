package contest

import (
	"net/http"
	"time"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/contracts/usecase"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

type ContestController struct {
	Usecase usecase.ContestUsecase
}

func NewContestController(u usecase.ContestUsecase) *ContestController {
	return &ContestController{Usecase: u}
}

func (c *ContestController) GetAllContests(ctx *gin.Context) {
	contests, err := c.Usecase.GetAllContests()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contests"})
		return
	}
	ctx.JSON(http.StatusOK, contests)
}

func (c *ContestController) CreateContest(ctx *gin.Context) {
	var contest entities.Contest
	if err := ctx.ShouldBindJSON(&contest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contest.PostedAt = time.Now()
	if err := c.Usecase.CreateContest(&contest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contest"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Contest created successfully"})
}
