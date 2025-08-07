package topics

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *TopicService) InsertTopics(
	ctx context.Context,
	topics []models.Topic,
) ([]int64, error) {
	allTaskIds := utils.Map(
		topics,
		func(topic models.Topic) []int64 {
			return topic.TaskIds
		},
	)

	if err := s.checkForTaskExistence(ctx, allTaskIds); err != nil {
		return nil, errors.Wrap(err, "checking for task existence")
	}

	return s.topicStorage.InsertTopics(ctx, topics)
}
