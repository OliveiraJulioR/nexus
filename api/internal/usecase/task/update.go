package usecase

import (
	"context"

	"github.com/OliveiraJulioR/nexus/api/internal/entity"
	"github.com/OliveiraJulioR/nexus/api/internal/repository"
)

type UpdateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskUseCase struct {
	repo repository.TaskRepository
}

func NewUpdateTask(repository repository.TaskRepository) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{repo: repository}
}

func (u *UpdateTaskUseCase) Execute(ctx context.Context, ID string, taskInput UpdateTaskInput) (*entity.Task, error) {
	task, err := u.repo.FindByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, nil
	}

	if taskInput.Title != "" {
		task.Title = taskInput.Title
	}

	if taskInput.Description != "" {
		task.Description = taskInput.Description
	}

	return u.repo.Update(ctx, task)
}
