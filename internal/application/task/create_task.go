package task

import (
	"context"
	"time"

	"github.com/williamokano/golang-web-api/internal/domain/task"
)

type CreateTask struct {
	taskRepository task.TaskRepository
}

func NewCreateTask(taskRepository task.TaskRepository) *CreateTask {
	return &CreateTask{
		taskRepository: taskRepository,
	}
}

func (uc *CreateTask) Execute(ctx context.Context, title, description string, dueDate *time.Time) (*task.Task, error) {
	newTask := &task.Task{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := uc.taskRepository.Save(ctx, newTask); err != nil {
		return nil, err
	}

	return newTask, nil
}
