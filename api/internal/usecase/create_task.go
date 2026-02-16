package usecase

import (
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

func (u *CreateTaskUseCase) Execute(input CreateTaskInput) (*entity.Task, error) {
	task := &entity.Task{
		Title:       input.Title,
		Description: input.Description,
	}

	return u.repo.Create(task)
}
