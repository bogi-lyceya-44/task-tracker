package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *TaskService) UpdateTasks(
	ctx context.Context,
	tasks []models.UpdatedTask,
) error {
	if err := checkForSelfDependency(
		utils.Map(
			tasks,
			func(task models.UpdatedTask) idWithDependencies {
				return idWithDependencies{task.ID, task.Dependencies}
			},
		),
	); err != nil {
		return errors.Wrap(err, "checking self dependency")
	}

	return s.storage.UpdateTasks(ctx, tasks)
}
