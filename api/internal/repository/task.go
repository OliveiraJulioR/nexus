package repository

import "github.com/OliveiraJulioR/nexus/api/internal/entity"

type TaskRepository interface {
	Create(task *entity.Task) (*entity.Task, error)
	FindAll() ([]entity.Task, error)
}
