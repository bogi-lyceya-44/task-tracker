package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	tasks_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/tasks"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateTasks(
	ctx context.Context,
	req *desc.CreateTasksRequest,
) (*desc.CreateTasksResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	tasks, err := utils.MapWithError(
		req.TasksToCreate,
		MapCreateTaskPrototypeToDomain,
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "mapping tasks: %v", err)
	}

	ids, err := i.taskService.InsertTasks(ctx, tasks)
	if err != nil {
		errorCode := codes.Internal

		if errors.Is(err, tasks_service.ErrSelfDependent) ||
			errors.Is(err, tasks_service.ErrTaskDoesNotExist) {
			errorCode = codes.InvalidArgument
		}

		return nil, status.Errorf(errorCode, "inserting tasks: %v", err)
	}

	return &desc.CreateTasksResponse{Ids: ids}, nil
}
