package tasks

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/bogi-lyceya-44/common/pkg/set"
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

type idWithDependencies utils.Pair[int64, []int64]

func (s *TaskService) checkForSelfDependency(pairs []idWithDependencies) error {
	selfDependent := utils.Filter(
		pairs,
		func(pair idWithDependencies) bool {
			return slices.Contains(pair.Second, pair.First)
		},
	)

	if len(selfDependent) > 0 {
		ids := utils.Map(
			pairs,
			func(pair idWithDependencies) string {
				return strconv.Itoa(int(pair.First))
			},
		)

		return errors.Wrap(
			ErrSelfDependent,
			fmt.Sprintf(
				"found self dependent tasks: %v",
				strings.Join(ids, ","),
			),
		)
	}

	return nil
}

func (s *TaskService) checkForTaskDependencyExistence(
	ctx context.Context,
	pairs []idWithDependencies,
) error {
	idsMap := set.NewEmptyWithCapacity[int64](len(pairs))

	for _, pair := range pairs {
		idsMap.Add(pair.Second...)
	}

	uniqueIds := idsMap.Slice()

	found, err := s.storage.GetTasks(ctx, uniqueIds)
	if err != nil {
		return errors.Wrap(err, "checking for task existence")
	}

	foundIds := set.New(
		utils.Map(
			found,
			(models.Task).GetID,
		)...,
	)

	nonExistentTasks := make([]int, 0, len(uniqueIds)-len(foundIds))

	for _, pair := range pairs {
		for _, dependency := range pair.Second {
			if !foundIds.Contains(dependency) {
				nonExistentTasks = append(nonExistentTasks, int(dependency))
			}
		}
	}

	if len(nonExistentTasks) > 0 {
		stringified := utils.Map(
			nonExistentTasks,
			strconv.Itoa,
		)

		return errors.Wrap(
			ErrTaskDoesNotExist,
			strings.Join(stringified, ","),
		)
	}

	return nil
}
