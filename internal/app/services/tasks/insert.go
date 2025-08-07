package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

func (s *TaskService) InsertTasks(
	ctx context.Context,
	tasks []models.Task,
) ([]int64, error) {
	return s.storage.InsertTasks(ctx, tasks)
}
