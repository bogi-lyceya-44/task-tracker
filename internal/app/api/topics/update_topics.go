package topics

import (
	"context"
	"errors"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	topics_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/topics"
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
		errorCode := codes.Internal

		if errors.Is(err, topics_service.ErrTaskDoesNotExist) {
			errorCode = codes.InvalidArgument
		}

		return nil, status.Errorf(errorCode, "updating topics: %v", err)
	}

	return &desc.UpdateTopicsResponse{}, nil
}
