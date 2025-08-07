package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetTasks(
	ctx context.Context,
	req *desc.GetTasksRequest,
) (*desc.GetTasksResponse, error) {
	tasks, err := i.taskService.GetTasks(ctx, req.GetIds())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	mappedTasks, err := utils.MapWithError(
		tasks,
		MapDomainTaskToProto,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.GetTasksResponse{Tasks: mappedTasks}, nil
}
