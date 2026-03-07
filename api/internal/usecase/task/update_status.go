package usecase

import (
	"context"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	"github.com/OliveiraJulioR/nexus/api/internal/repository"
)

type UpdateStatusInput struct {
	Status string `json:"status" binding:"required"`
}

type UpdateStatusUseCase struct {
	repo repository.TaskRepository
}

func NewUpdateStatusUseCase(repository repository.TaskRepository) *UpdateStatusUseCase {
	return &UpdateStatusUseCase{repo: repository}
}

func (u *UpdateStatusUseCase) Execute(ctx context.Context, ID string, status entity.TaskStatus) (*entity.Task, error) {
	task, err := u.repo.FindByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, nil
	}

	task.Status = status

	return u.repo.Update(ctx, task)
}
