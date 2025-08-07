package topics

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/topics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateTopics(
	ctx context.Context,
	req *desc.CreateTopicsRequest,
) (*desc.CreateTopicsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	topics := utils.Map(
		req.TopicsToCreate,
		MapCreateTopicPrototypeToDomain,
	)

	ids, err := i.topicService.InsertTopics(ctx, topics)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inserting topics: %v", err)
	}

	return &desc.CreateTopicsResponse{Ids: ids}, nil
}
