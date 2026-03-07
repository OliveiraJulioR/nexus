package usecase

import (
	"context"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	"github.com/OliveiraJulioR/nexus/api/internal/repository"
)

type UpdateStatusInput struct {
	Status string `json:title`
}

type UpdateStatusUseCase struct {
	ID   string
	repo repository.TaskRepository
}

func NewUpdateStatusUseCase(ID string, repository repository.TaskRepository) *UpdateStatusUseCase {
	return &UpdateStatusUseCase{ID: ID, repo: repository}
}

func (u *UpdateStatusUseCase) Execute(ctx context.Context, input UpdateStatusInput) (*entity.Task, error) {
	task, err := u.repo.FindByID(ctx, u.ID)

	if err != nil {
		return nil, err
	}

	task.Status = input.Status

	return u.repo.Update(ctx, task)
}
