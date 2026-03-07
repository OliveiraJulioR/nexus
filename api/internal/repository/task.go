package repository

import (
	"context"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
)

type TaskRepository interface {
	Create(ctx context.Context, task *entity.Task) (*entity.Task, error)
	FindAll(ctx context.Context) ([]entity.Task, error)
	Update(ctx context.Context, task *entity.Task) (*entity.Task, error)
	FindByID(ctx context.Context, id string) (*entity.Task, error)
	Delete(ctx context.Context, id string) error
}
