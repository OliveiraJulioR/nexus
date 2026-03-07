package handler

import (
	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	usecase "github.com/OliveiraJulioR/nexus/api/internal/usecase/task"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	CreateTaskUseCase   *usecase.CreateTaskUseCase
	UpdateStatusUseCase *usecase.UpdateStatusUseCase
}

func NewTaskHandler(createTaskUseCase *usecase.CreateTaskUseCase, updateStatusUseCase *usecase.UpdateStatusUseCase) *TaskHandler {
	return &TaskHandler{
		CreateTaskUseCase:   createTaskUseCase,
		UpdateStatusUseCase: updateStatusUseCase,
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

func (h *TaskHandler) UpdateStatus(c *gin.Context) {
	// 1. Pega o ID direto da URL (ex: /tasks/123e4567-e89b-12d3-a456-426614174000/status)
	id := c.Param("id")

	// 2. Lê o JSON do corpo da requisição
	var input usecase.UpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "O campo status é obrigatório"})
		return
	}

	// 3. Lembra do nosso Enum TaskStatus? Já podemos validar aqui se o cara mandou um status que existe!
	// Substitua 'entity' pelo pacote real da sua entidade
	statusEnum := entity.TaskStatus(input.Status)
	if !statusEnum.IsValid() {
		c.JSON(400, gin.H{"error": "Status inválido. Use TODO, IN_PROGRESS ou DONE"})
		return
	}

	// 4. Chama o UseCase para atualizar no banco (você vai criar esse UseCase depois)
	taskAtualizada, err := h.UpdateStatusUseCase.Execute(c.Request.Context(), id, statusEnum)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 5. Retorna sucesso
	c.JSON(200, gin.H{
		"message": "Status atualizado com sucesso!",
		"id":      id,
		"task":    taskAtualizada,
		"status":  input.Status,
	})
}
