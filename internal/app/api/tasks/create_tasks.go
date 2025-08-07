package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateTasks(
	ctx context.Context,
	req *desc.CreateTasksRequest,
) (*desc.CreateTasksResponse, error) {
	// TODO: implement utils.MapWithError
	var mappingError error
	tasks := utils.Map(
		req.TasksToCreate,
		func(task *desc.CreateTasksRequest_TaskPrototype) models.Task {
			result, err := MapCreateTaskPrototypeToDomain(task)
			if err != nil {
				mappingError = err
			}
			return result
		},
	)

	if mappingError != nil {
		return nil, status.Error(codes.InvalidArgument, mappingError.Error())
	}

	ids, err := i.taskService.InsertTasks(ctx, tasks)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateTasksResponse{Ids: ids}, nil
}
