package tasks

import (
	"context"

	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteTasks(
	ctx context.Context,
	req *desc.DeleteTasksRequest,
) (*desc.DeleteTasksResponse, error) {
	err := i.taskService.DeleteTasks(ctx, req.Ids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.DeleteTasksResponse{}, nil
}
