package topics

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

func (s *TopicService) GetTopics(
	ctx context.Context,
	ids []int64,
) ([]models.Topic, error) {
	return s.storage.GetTopics(ctx, ids)
}
