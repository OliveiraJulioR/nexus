package implementations

import (
	"context"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskPostgresRepository struct {
	postgresPool *pgxpool.Pool
}

func NewTaskPostgresRepository(postgresPool *pgxpool.Pool) *TaskPostgresRepository {
	return &TaskPostgresRepository{postgresPool: postgresPool}
}

func (r *TaskPostgresRepository) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)`
	_, err := r.postgresPool.Exec(ctx, query, task.Title, task.Description, task.Status)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskPostgresRepository) FindAll(ctx context.Context) ([]entity.Task, error) {
	return nil, nil
}

func (r *TaskPostgresRepository) Update(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	return nil, nil
}

func (r *TaskPostgresRepository) FindByID(id string) (*entity.Task, error) {
	return nil, nil
}

func (r *TaskPostgresRepository) Delete(ctx context.Context, id string) error {
	return nil
}
