package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateTasks(
	ctx context.Context,
	req *desc.UpdateTasksRequest,
) (*desc.UpdateTasksResponse, error) {
	tasks, err := utils.MapWithError(
		req.TasksToUpdate,
		MapUpdateTaskPrototypeToUpdatedTask,
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = i.taskService.UpdateTasks(ctx, tasks); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.UpdateTasksResponse{}, nil
}
