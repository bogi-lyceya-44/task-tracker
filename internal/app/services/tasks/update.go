package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

func (s *TaskService) UpdateTasks(
	ctx context.Context,
	tasks []models.UpdatedTask,
) error {
	return s.storage.UpdateTasks(ctx, tasks)
}
