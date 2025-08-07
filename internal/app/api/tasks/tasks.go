package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
)

type TaskService interface {
	GetTasks(ctx context.Context, ids []int64) ([]models.Task, error)
	InsertTasks(ctx context.Context, tasks []models.Task) ([]int64, error)
	UpdateTasks(ctx context.Context, tasks []models.UpdatedTask) error
	DeleteTasks(ctx context.Context, ids []int64) error
}

type Implementation struct {
	desc.UnimplementedTaskServiceServer

	taskService TaskService
}

func New(taskService TaskService) *Implementation {
	return &Implementation{
		taskService: taskService,
	}
}
