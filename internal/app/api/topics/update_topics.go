package topics

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/topics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateTopics(
	ctx context.Context,
	req *desc.UpdateTopicsRequest,
) (*desc.UpdateTopicsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	tasks, err := utils.MapWithError(
		req.TopicsToUpdate,
		MapUpdateTopicPrototypeToUpdatedTopic,
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "mapping topics: %v", err)
	}

	if err = i.topicService.UpdateTopics(ctx, tasks); err != nil {
		return nil, status.Errorf(codes.Internal, "updating topics: %v", err)
	}

	return &desc.UpdateTopicsResponse{}, nil
}
