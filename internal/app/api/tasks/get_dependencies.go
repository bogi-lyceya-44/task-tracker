package tasks

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetAllDependencies(
	ctx context.Context,
	req *desc.GetAllDependenciesRequest,
) (*desc.GetAllDependenciesResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	ids := utils.Unique(req.GetIds())

	dependenciesById, err := i.taskService.GetDependencies(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting all dependencies: %v", err)
	}

	response := &desc.GetAllDependenciesResponse{
		DependenciesById: make(
			map[int64]*desc.GetAllDependenciesResponse_IdArray,
			len(dependenciesById),
		),
	}

	for id, deps := range dependenciesById {
		response.DependenciesById[id] = &desc.GetAllDependenciesResponse_IdArray{
			Ids: deps,
		}
	}

	return response, nil
}
