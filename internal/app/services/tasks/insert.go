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
	idsWithDependencies := utils.Map(
		tasks,
		func(task models.Task) idWithDependencies {
			return idWithDependencies{task.ID, task.Dependencies}
		},
	)

	if err := s.checkForSelfDependency(idsWithDependencies); err != nil {
		return nil, errors.Wrap(err, "checking self dependency")
	}

	if err := s.checkForTaskDependencyExistence(ctx, idsWithDependencies); err != nil {
		return nil, errors.Wrap(err, "checking task existence")
	}

	return s.storage.InsertTasks(ctx, tasks)
}
