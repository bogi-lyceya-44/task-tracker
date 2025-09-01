package topics

import (
	"context"

	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/topics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteTopics(
	ctx context.Context,
	req *desc.DeleteTopicsRequest,
) (*desc.DeleteTopicsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	err := i.topicService.DeleteTopics(ctx, req.Ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "deleting topics: %v", err)
	}

	return &desc.DeleteTopicsResponse{}, nil
}
