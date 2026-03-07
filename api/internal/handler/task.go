package handler

import (
	"errors"
	"fmt"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	usecase "github.com/OliveiraJulioR/nexus/api/internal/usecase/task"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type TaskHandler struct {
	CreateTaskUseCase   *usecase.CreateTaskUseCase
	UpdateStatusUseCase *usecase.UpdateStatusUseCase
	FindAllUseCase      *usecase.FindAllUseCase
	FindByIDUseCase     *usecase.FindByIDUseCase
	DeleteTaskUseCase   *usecase.DeleteTaskUseCase
	UpdateTaskUseCase   *usecase.UpdateTaskUseCase
}

func NewTaskHandler(createTaskUseCase *usecase.CreateTaskUseCase, updateStatusUseCase *usecase.UpdateStatusUseCase, findAllUseCase *usecase.FindAllUseCase, findByIDUseCase *usecase.FindByIDUseCase, deleteUsCase *usecase.DeleteTaskUseCase, updateTaskUseCase *usecase.UpdateTaskUseCase) *TaskHandler {
	return &TaskHandler{
		CreateTaskUseCase:   createTaskUseCase,
		UpdateStatusUseCase: updateStatusUseCase,
		FindAllUseCase:      findAllUseCase,
		FindByIDUseCase:     findByIDUseCase,
		DeleteTaskUseCase:   deleteUsCase,
		UpdateTaskUseCase:   updateTaskUseCase,
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

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	baseURL := fmt.Sprintf("%s://%s", scheme, c.Request.Host)

	response := BuildTaskResponse(result, baseURL)

	c.JSON(201, response)
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

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	baseURL := fmt.Sprintf("%s://%s", scheme, c.Request.Host)

	response := BuildTaskResponse(taskAtualizada, baseURL)

	c.JSON(201, response)
}

func (h *TaskHandler) FindAll(c *gin.Context) {
	tasks, err := h.FindAllUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	baseURL := "http://" + c.Request.Host // Simplificado

	var responseList []TaskResponse
	for _, t := range tasks {
		// Precisamos passar o ponteiro &t porque a variável no loop é uma cópia
		taskCopy := t
		responseList = append(responseList, BuildTaskResponse(&taskCopy, baseURL))
	}

	c.JSON(200, responseList)
}

func (h *TaskHandler) FindByID(c *gin.Context) {
	id := c.Param("id")

	task, err := h.FindByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	baseURL := fmt.Sprintf("%s://%s", scheme, c.Request.Host)

	response := BuildTaskResponse(task, baseURL)

	c.JSON(200, response)
}

func (h *TaskHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.DeleteTaskUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}

func (h *TaskHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input usecase.UpdateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := h.UpdateTaskUseCase.Execute(c, id, input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	baseURL := fmt.Sprintf("%s://%s", scheme, c.Request.Host)

	response := BuildTaskResponse(result, baseURL)

	c.JSON(200, response)
}
