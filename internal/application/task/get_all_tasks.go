package task

import (
	"context"

	"github.com/williamokano/golang-web-api/internal/domain/task"
)

type GetAllTasks struct {
	taskRepository task.TaskRepository
}

func NewGetAllTasks(taskRepository task.TaskRepository) *GetAllTasks {
	return &GetAllTasks{taskRepository: taskRepository}
}

func (uc GetAllTasks) Execute(ctx context.Context) ([]*task.Task, error) {
	return uc.taskRepository.GetAll(ctx)
}
