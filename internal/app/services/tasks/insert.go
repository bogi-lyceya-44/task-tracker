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

	allDependencies := utils.Map(
		tasks,
		func(task models.Task) []int64 {
			return task.Dependencies
		},
	)

	if err := s.checkForTaskDependencyExistence(ctx, allDependencies); err != nil {
		return nil, errors.Wrap(err, "checking task existence")
	}

	for _, pair := range idsWithDependencies {
		if err := s.checkForCyclicDependency(ctx, pair); err != nil {
			return nil, errors.Wrap(err, "checking cyclic dependency")
		}
	}

	return s.storage.InsertTasks(ctx, tasks)
}
