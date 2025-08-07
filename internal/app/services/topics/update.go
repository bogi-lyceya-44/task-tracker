package topics

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *TopicService) UpdateTopics(
	ctx context.Context,
	topics []models.UpdatedTopic,
) error {
	allTaskIds := utils.Map(
		topics,
		func(topic models.UpdatedTopic) []int64 {
			return topic.TaskIds
		},
	)

	if err := s.checkForTaskExistence(ctx, allTaskIds); err != nil {
		return errors.Wrap(err, "checking for task existence")
	}

	return s.topicStorage.UpdateTopics(ctx, topics)
}
