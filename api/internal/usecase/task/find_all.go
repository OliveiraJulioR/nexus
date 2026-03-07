package usecase

import (
	"context"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	"github.com/OliveiraJulioR/nexus/api/internal/repository"
)

type FindAllUseCase struct {
	repo repository.TaskRepository
}

func NewFindAllUseCase(repository repository.TaskRepository) *FindAllUseCase {
	return &FindAllUseCase{repo: repository}
}

func (u *FindAllUseCase) Execute(ctx context.Context) ([]entity.Task, error) {
	return u.repo.FindAll(ctx)
}
