package tasks

import (
	"context"
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

		return errors.Wrapf(
			ErrSelfDependent,
			"tasks: %v",
			strings.Join(ids, ","),
		)
	}

	return nil
}

func (s *TaskService) checkForTaskDependencyExistence(
	ctx context.Context,
	dependencies [][]int64,
) error {
	totalSize := utils.Reduce(
		dependencies,
		func(accumulator int, value []int64) int {
			return accumulator + len(value)
		},
		0,
	)

	idsMap := set.NewEmptyWithCapacity[int64](totalSize)

	for _, deps := range dependencies {
		idsMap.Add(deps...)
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

	if len(uniqueIds) > len(foundIds) {
		nonExistentTasks := make([]int, 0, len(uniqueIds)-len(foundIds))

		for _, deps := range dependencies {
			for _, dependency := range deps {
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

			return errors.Wrapf(
				ErrTaskDoesNotExist,
				"tasks: %v",
				strings.Join(stringified, ","),
			)
		}
	}

	return nil
}

func (s *TaskService) checkForCyclicDependency(
	ctx context.Context,
	pair idWithDependencies,
) error {
	var deps []int64

	for _, dependency := range pair.Second {
		fetched, err := s.storage.GetDependencies(ctx, dependency)
		if err != nil {
			return errors.Wrap(err, "getting dependencies")
		}

		deps = append(deps, fetched...)
	}

	if slices.Contains(deps, pair.First) {
		return errors.Wrapf(
			ErrCyclicDependency,
			"task %d",
			pair.First,
		)
	}

	return nil
}
