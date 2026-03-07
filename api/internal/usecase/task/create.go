// Executa a criação da task e persiste no banco
package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	"github.com/OliveiraJulioR/nexus/api/internal/repository"
)

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTaskUseCase struct {
	repo repository.TaskRepository
}

func NewCreateTaskUseCase(repository repository.TaskRepository) *CreateTaskUseCase {
	return &CreateTaskUseCase{repo: repository}
}

func (u *CreateTaskUseCase) Execute(ctx context.Context, input CreateTaskInput) (*entity.Task, error) {

	ID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	task := &entity.Task{
		ID:          ID.String(),
		Title:       input.Title,
		Description: input.Description,
		Status:      entity.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return u.repo.Create(ctx, task)
}
