package topics

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

type topicStorage interface {
	GetTopics(ctx context.Context, ids []int64) ([]models.Topic, error)
	InsertTopics(ctx context.Context, tasks []models.Topic) ([]int64, error)
	UpdateTopics(ctx context.Context, tasks []models.UpdatedTopic) error
	DeleteTopics(ctx context.Context, ids []int64) error
}

type taskStorage interface {
	GetTasks(ctx context.Context, ids []int64) ([]models.Task, error)
}

type TopicService struct {
	topicStorage topicStorage
	taskStorage  taskStorage
}

func NewTopicService(
	topicStorage topicStorage,
	taskStorage taskStorage,
) *TopicService {
	return &TopicService{
		topicStorage: topicStorage,
		taskStorage:  taskStorage,
	}
}
