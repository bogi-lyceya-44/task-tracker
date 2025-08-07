package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateTasks(
	ctx context.Context,
	req *desc.UpdateTasksRequest,
) (*desc.UpdateTasksResponse, error) {
	// TODO: implement utils.MapWithError
	var mappingError error
	tasks := utils.Map(
		req.TasksToUpdate,
		func(task *desc.UpdateTasksRequest_TaskPrototype) models.UpdatedTask {
			result, err := MapUpdateTaskPrototypeToUpdatedTask(task)
			if err != nil {
				mappingError = err
			}
			return result
		},
	)

	if mappingError != nil {
		return nil, status.Error(codes.InvalidArgument, mappingError.Error())
	}

	err := i.taskService.UpdateTasks(ctx, tasks)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.UpdateTasksResponse{}, nil
}
