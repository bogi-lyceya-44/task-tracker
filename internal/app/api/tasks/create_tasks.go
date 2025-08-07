package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateTasks(
	ctx context.Context,
	req *desc.CreateTasksRequest,
) (*desc.CreateTasksResponse, error) {
	tasks, err := utils.MapWithError(
		req.TasksToCreate,
		MapCreateTaskPrototypeToDomain,
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ids, err := i.taskService.InsertTasks(ctx, tasks)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateTasksResponse{Ids: ids}, nil
}
