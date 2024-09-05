package task

import (
	"context"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type TaskRepository interface {
	GetByID(ctx context.Context, id int) (*Task, error)
	GetAll(ctx context.Context) ([]*Task, error)
	Save(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id int) error
}
