package topics

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

func (s *TopicService) UpdateTopics(
	ctx context.Context,
	tasks []models.UpdatedTopic,
) error {
	return s.storage.UpdateTopics(ctx, tasks)
}
