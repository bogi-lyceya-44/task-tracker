package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
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

	var mappingError error
	mappedTasks := utils.Map(
		tasks,
		func(task models.Task) *desc.Task {
			result, err := MapDomainTaskToProto(task)
			if err != nil {
				mappingError = err
			}
			return result
		},
	)

	if mappingError != nil {
		return nil, status.Error(codes.Internal, mappingError.Error())
	}

	return &desc.GetTasksResponse{Tasks: mappedTasks}, nil
}
