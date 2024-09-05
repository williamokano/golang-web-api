package persistence

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/williamokano/golang-web-api/internal/domain/task"
)

type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{db: db}
}

func (t TaskRepositoryImpl) GetByID(ctx context.Context, id int) (*task.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepositoryImpl) GetAll(ctx context.Context) ([]*task.Task, error) {
	rows, err := t.db.QueryContext(ctx, GET_ALL_QUERY)
	if err != nil {
		return nil, fmt.Errorf("failed to query get all tasks: %w", err)
	}
	defer rows.Close()

	var tasks = make([]*task.Task, 0)
	for rows.Next() {
		var t task.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Done, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan rows: %w", err)
	}

	return tasks, nil
}

func (t TaskRepositoryImpl) Save(ctx context.Context, task *task.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepositoryImpl) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

var _ task.TaskRepository = &TaskRepositoryImpl{}

var GET_ALL_QUERY = `SELECT id, title, description, due_date, done, created_at, updated_at, deleted_at FROM tasks;`
