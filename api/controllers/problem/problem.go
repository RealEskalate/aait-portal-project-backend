package problem

import (
	"net/http"
	"strconv"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/contracts/usecase"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

type ProblemController struct {
	ProblemUsecase usecase.ProblemUsecase
}

func NewProblemController(problemUsecase usecase.ProblemUsecase) *ProblemController {
	return &ProblemController{
		ProblemUsecase: problemUsecase,
	}
}

func (controller *ProblemController) CreateProblem(ctx *gin.Context) {
	var newProblem entities.Problem
	if err := ctx.ShouldBindJSON(&newProblem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.ProblemUsecase.CreateProblem(&newProblem); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create problem"})
		return
	}

	ctx.JSON(http.StatusCreated, newProblem)
}

func (controller *ProblemController) GetProblem(ctx *gin.Context) {
	problemIDStr := ctx.Param("id")
	problemID, err := strconv.ParseUint(problemIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid problem ID"})
		return
	}

	problem, err := controller.ProblemUsecase.GetProblemByID(uint(problemID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
		return
	}

	ctx.JSON(http.StatusOK, problem)
}

func (controller *ProblemController) GetAllProblems(ctx *gin.Context) {
	problems, err := controller.ProblemUsecase.GetAllProblems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch problems"})
		return
	}

	ctx.JSON(http.StatusOK, problems)
}

func (controller *ProblemController) UpdateProblem(ctx *gin.Context) {
	problemIDStr := ctx.Param("id")
	problemID, err := strconv.ParseUint(problemIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid problem ID"})
		return
	}

	var updatedProblem entities.Problem
	if err := ctx.ShouldBindJSON(&updatedProblem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProblem.ID = uint(problemID)
	if err := controller.ProblemUsecase.UpdateProblem(&updatedProblem); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update problem"})
		return
	}

	ctx.JSON(http.StatusOK, updatedProblem)
}

func (controller *ProblemController) DeleteProblem(ctx *gin.Context) {
	problemIDStr := ctx.Param("id")
	problemID, err := strconv.ParseUint(problemIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid problem ID"})
		return
	}

	if err := controller.ProblemUsecase.DeleteProblem(uint(problemID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete problem"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Problem deleted successfully"})
}

func (controller *ProblemController) GetProblemsByDifficulty(ctx *gin.Context) {
	difficultyLevel := ctx.Param("difficulty")
	problems, err := controller.ProblemUsecase.GetProblemsByDifficulty(difficultyLevel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch problems"})
		return
	}

	ctx.JSON(http.StatusOK, problems)
}

func (controller *ProblemController) GetProblemsByTrack(ctx *gin.Context) {
	trackName := ctx.Param("track")
	problems, err := controller.ProblemUsecase.GetProblemsByTrack(trackName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch problems"})
		return
	}

	ctx.JSON(http.StatusOK, problems)
}
