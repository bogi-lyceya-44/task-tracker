package topics

import (
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/topics"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapCreateTopicPrototypeToDomain(topic *desc.CreateTopicsRequest_TopicPrototype) models.Topic {
	return models.Topic{
		Name:    topic.GetName(),
		TaskIds: topic.GetTaskIds(),
	}
}

func MapDomainTopicToProto(topic models.Topic) *desc.Topic {
	return &desc.Topic{
		Id:        topic.ID,
		Name:      topic.Name,
		TaskIds:   topic.TaskIds,
		CreatedAt: timestamppb.New(topic.CreatedAt),
		UpdatedAt: timestamppb.New(topic.UpdatedAt),
	}
}

func MapUpdateTopicPrototypeToUpdatedTopic(topic *desc.UpdateTopicsRequest_TopicPrototype) (models.UpdatedTopic, error) {
	result := models.UpdatedTopic{ID: topic.GetId()}

	if topic.Name != nil {
		result.Name = topic.Name
	}

	if topic.TaskIds != nil {
		result.TaskIds = append(result.TaskIds, topic.TaskIds...)
	}

	return result, nil
}
