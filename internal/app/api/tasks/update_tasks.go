package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	tasks_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/tasks"
	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/tasks"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateTasks(
	ctx context.Context,
	req *desc.UpdateTasksRequest,
) (*desc.UpdateTasksResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	tasks, err := utils.MapWithError(
		req.TasksToUpdate,
		MapUpdateTaskPrototypeToUpdatedTask,
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "mapping tasks: %v", err)
	}

	if err = i.taskService.UpdateTasks(ctx, tasks); err != nil {
		errorCode := codes.Internal

		if errors.Is(err, tasks_service.ErrSelfDependent) ||
			errors.Is(err, tasks_service.ErrTaskDoesNotExist) ||
			errors.Is(err, tasks_service.ErrCyclicDependency) {
			errorCode = codes.InvalidArgument
		}

		return nil, status.Errorf(errorCode, "updating tasks: %v", err)
	}

	return &desc.UpdateTasksResponse{}, nil
}
