package usecase

import (
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

func (u *UpdateStatusUseCase) Execute(input UpdateStatusInput) (*entity.Task, error) {
	task, err := u.repo.FindByID(u.ID)

	if err != nil {
		return nil, err
	}

	task.Status = input.Status

	return u.repo.Update(task)
}
