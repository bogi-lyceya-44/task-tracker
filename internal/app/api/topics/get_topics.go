package topics

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/topics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetTopics(
	ctx context.Context,
	req *desc.GetTopicsRequest,
) (*desc.GetTopicsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	ids := utils.Unique(req.GetIds())

	tasks, err := i.topicService.GetTopics(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting topics: %v", err)
	}

	mappedTopics := utils.Map(
		tasks,
		MapDomainTopicToProto,
	)

	return &desc.GetTopicsResponse{Topics: mappedTopics}, nil
}
