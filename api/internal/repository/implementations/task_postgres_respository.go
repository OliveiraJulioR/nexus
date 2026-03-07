package implementations

import (
	"context"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	"github.com/OliveiraJulioR/nexus/api/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskPostgresRepository struct {
	postgresPool *pgxpool.Pool
}

func NewTaskPostgresRepository(postgresPool *pgxpool.Pool) repository.TaskRepository {
	return &TaskPostgresRepository{postgresPool: postgresPool}
}

func (r *TaskPostgresRepository) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	query := `
		INSERT INTO tasks (title, description, status) 
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := r.postgresPool.QueryRow(ctx, query, task.Title, task.Description, task.Status).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskPostgresRepository) FindAll(ctx context.Context) ([]entity.Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks ORDER BY created_at DESC`

	rows, err := r.postgresPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var t entity.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskPostgresRepository) FindByID(ctx context.Context, id string) (*entity.Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1`

	var t entity.Task
	err := r.postgresPool.QueryRow(ctx, query, id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &t, nil
}

func (r *TaskPostgresRepository) Update(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	query := `
		UPDATE tasks 
		SET title = $1, description = $2, status = $3, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $4
		RETURNING updated_at
	`

	err := r.postgresPool.QueryRow(ctx, query, task.Title, task.Description, task.Status, task.ID).
		Scan(&task.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return task, nil
}

func (r *TaskPostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE id = $1`

	_, err := r.postgresPool.Exec(ctx, query, id)
	return err
}
