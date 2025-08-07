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
	idsWithDependencies := utils.Map(
		tasks,
		func(task models.UpdatedTask) idWithDependencies {
			return idWithDependencies{task.ID, task.Dependencies}
		},
	)

	if err := s.checkForSelfDependency(idsWithDependencies); err != nil {
		return errors.Wrap(err, "checking self dependency")
	}

	allDependencies := utils.Map(
		tasks,
		func(task models.UpdatedTask) []int64 {
			return task.Dependencies
		},
	)

	if err := s.checkForTaskDependencyExistence(ctx, allDependencies); err != nil {
		return errors.Wrap(err, "checking task existence")
	}

	return s.storage.UpdateTasks(ctx, tasks)
}
