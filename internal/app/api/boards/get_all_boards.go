package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetAllBoards(
	ctx context.Context,
	req *desc.GetAllBoardsRequest,
) (*desc.GetAllBoardsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	tasks, err := i.boardService.GetBoards(ctx, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting boards: %v", err)
	}

	mappedBoards := utils.Map(
		tasks,
		MapDomainBoardToProto,
	)

	return &desc.GetAllBoardsResponse{Boards: mappedBoards}, nil
}
