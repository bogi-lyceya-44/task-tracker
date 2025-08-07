package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

type storage interface {
	GetTasks(ctx context.Context, ids []int64) ([]models.Task, error)
	InsertTasks(ctx context.Context, tasks []models.Task) ([]int64, error)
	UpdateTasks(ctx context.Context, tasks []models.UpdatedTask) error
	DeleteTasks(ctx context.Context, ids []int64) error
}

type TaskService struct {
	storage storage
}

func NewTaskService(storage storage) *TaskService {
	return &TaskService{
		storage: storage,
	}
}
