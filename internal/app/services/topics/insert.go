package topics

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

func (s *TopicService) InsertTopics(
	ctx context.Context,
	tasks []models.Topic,
) ([]int64, error) {
	return s.storage.InsertTopics(ctx, tasks)
}
