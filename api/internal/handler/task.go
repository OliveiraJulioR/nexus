package handler

import (
	usecase "github.com/OliveiraJulioR/nexus/api/internal/usecase/task"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	CreateTaskUseCase *usecase.CreateTaskUseCase
}

func NewTaskHandler(createTaskUseCase *usecase.CreateTaskUseCase) *TaskHandler {
	return &TaskHandler{
		CreateTaskUseCase: createTaskUseCase,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input usecase.CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := h.CreateTaskUseCase.Execute(c, input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, result)
}
