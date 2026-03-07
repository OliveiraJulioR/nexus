package usecase

import (
	"context"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	"github.com/OliveiraJulioR/nexus/api/internal/repository"
)

type FindByIDUseCase struct {
	repo repository.TaskRepository
}

func NewFindByIDUseCase(repository repository.TaskRepository) *FindByIDUseCase {
	return &FindByIDUseCase{repo: repository}
}

func (u *FindByIDUseCase) Execute(ctx context.Context, ID string) (*entity.Task, error) {
	return u.repo.FindByID(ctx, ID)
}
