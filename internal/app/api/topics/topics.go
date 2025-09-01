package topics

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/topics"
)

type TopicService interface {
	GetTopics(ctx context.Context, ids []int64) ([]models.Topic, error)
	InsertTopics(ctx context.Context, tasks []models.Topic) ([]int64, error)
	UpdateTopics(ctx context.Context, tasks []models.UpdatedTopic) error
	DeleteTopics(ctx context.Context, ids []int64) error
}

type Implementation struct {
	desc.UnimplementedTopicServiceServer

	topicService TopicService
}

func New(topicService TopicService) *Implementation {
	return &Implementation{
		topicService: topicService,
	}
}
