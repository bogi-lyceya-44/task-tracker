package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

func (s *TaskService) GetTasks(
	ctx context.Context,
	ids []int64,
) ([]models.Task, error) {
	return s.storage.GetTasks(ctx, ids)
}
