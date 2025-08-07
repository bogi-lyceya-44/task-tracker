package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetBoards(
	ctx context.Context,
	req *desc.GetBoardsRequest,
) (*desc.GetBoardsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	tasks, err := i.boardService.GetBoards(ctx, req.GetIds())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting boards: %v", err)
	}

	mappedBoards := utils.Map(
		tasks,
		MapDomainBoardToProto,
	)

	return &desc.GetBoardsResponse{Boards: mappedBoards}, nil
}
