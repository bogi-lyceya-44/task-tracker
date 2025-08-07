package boards

import (
	"context"
	"strconv"
	"strings"

	"github.com/bogi-lyceya-44/common/pkg/set"
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *BoardService) checkForTopicExistence(
	ctx context.Context,
	topicIds [][]int64,
) error {
	totalSize := utils.Reduce(
		topicIds,
		func(accumulator int, value []int64) int {
			return accumulator + len(value)
		},
		0,
	)

	topicIdsMap := set.NewEmptyWithCapacity[int64](totalSize)

	for _, ids := range topicIds {
		topicIdsMap.Add(ids...)
	}

	uniqueTopicIds := topicIdsMap.Slice()

	found, err := s.topicStorage.GetTopics(ctx, uniqueTopicIds)
	if err != nil {
		return errors.Wrap(err, "checking for topic existence")
	}

	foundIds := set.New(
		utils.Map(
			found,
			(models.Topic).GetID,
		)...,
	)

	if len(uniqueTopicIds) > len(foundIds) {
		nonExistentTopics := make([]int, 0, len(uniqueTopicIds)-len(foundIds))

		for _, ids := range topicIds {
			for _, id := range ids {
				if !foundIds.Contains(id) {
					nonExistentTopics = append(nonExistentTopics, int(id))
				}
			}
		}

		if len(nonExistentTopics) > 0 {
			stringified := utils.Map(
				nonExistentTopics,
				strconv.Itoa,
			)

			return errors.WithMessagef(
				ErrTopicDoesNotExist,
				"ids: %v",
				strings.Join(stringified, ","),
			)
		}
	}

	return nil
}
