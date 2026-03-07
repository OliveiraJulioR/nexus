package usecase

import (
	"context"

	"github.com/OliveiraJulioR/nexus/api/internal/repository"
)

type DeleteTaskUseCase struct {
	repo repository.TaskRepository
}

func NewDeleteTaskUseCase(repository repository.TaskRepository) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{repo: repository}
}

func (u *DeleteTaskUseCase) Execute(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
