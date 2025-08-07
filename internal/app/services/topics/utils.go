package topics

import (
	"context"
	"strconv"
	"strings"

	"github.com/bogi-lyceya-44/common/pkg/set"
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *TopicService) checkForTaskExistence(
	ctx context.Context,
	taskIds [][]int64,
) error {
	totalSize := utils.Reduce(
		taskIds,
		func(accumulator int, value []int64) int {
			return accumulator + len(value)
		},
		0,
	)

	taskIdsMap := set.NewEmptyWithCapacity[int64](totalSize)

	for _, ids := range taskIds {
		taskIdsMap.Add(ids...)
	}

	uniqueTaskIds := taskIdsMap.Slice()

	found, err := s.taskStorage.GetTasks(ctx, uniqueTaskIds)
	if err != nil {
		return errors.Wrap(err, "checking for task existence")
	}

	foundIds := set.New(
		utils.Map(
			found,
			(models.Task).GetID,
		)...,
	)

	if len(uniqueTaskIds) > len(foundIds) {
		nonExistentTasks := make([]int, 0, len(uniqueTaskIds)-len(foundIds))

		for _, ids := range taskIds {
			for _, id := range ids {
				if !foundIds.Contains(id) {
					nonExistentTasks = append(nonExistentTasks, int(id))
				}
			}
		}

		if len(nonExistentTasks) > 0 {
			stringified := utils.Map(
				nonExistentTasks,
				strconv.Itoa,
			)

			return errors.WithMessagef(
				ErrTaskDoesNotExist,
				"ids: %v",
				strings.Join(stringified, ","),
			)
		}
	}

	return nil
}
