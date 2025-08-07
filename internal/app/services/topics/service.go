package topics

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

type storage interface {
	GetTopics(ctx context.Context, ids []int64) ([]models.Topic, error)
	InsertTopics(ctx context.Context, tasks []models.Topic) ([]int64, error)
	UpdateTopics(ctx context.Context, tasks []models.UpdatedTopic) error
	DeleteTopics(ctx context.Context, ids []int64) error
}

type TopicService struct {
	storage storage
}

func NewTopicService(storage storage) *TopicService {
	return &TopicService{
		storage: storage,
	}
}
