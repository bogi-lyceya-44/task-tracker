package tasks

import (
	"context"

	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteTasks(
	ctx context.Context,
	req *desc.DeleteTasksRequest,
) (*desc.DeleteTasksResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	err := i.taskService.DeleteTasks(ctx, req.Ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "deleting tasks: %v", err)
	}

	return &desc.DeleteTasksResponse{}, nil
}
