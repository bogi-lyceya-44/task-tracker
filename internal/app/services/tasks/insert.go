package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *TaskService) InsertTasks(
	ctx context.Context,
	tasks []models.Task,
) ([]int64, error) {
	if err := checkForSelfDependency(
		utils.Map(
			tasks,
			func(task models.Task) idWithDependencies {
				return idWithDependencies{task.ID, task.Dependencies}
			},
		),
	); err != nil {
		return nil, errors.Wrap(err, "checking self dependency")
	}

	return s.storage.InsertTasks(ctx, tasks)
}
