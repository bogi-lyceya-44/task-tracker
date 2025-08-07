package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateBoards(
	ctx context.Context,
	req *desc.CreateBoardsRequest,
) (*desc.CreateBoardsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	boards := utils.Map(
		req.BoardsToCreate,
		MapCreateBoardPrototypeToDomain,
	)

	ids, err := i.boardService.InsertBoards(ctx, boards)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inserting boards: %v", err)
	}

	return &desc.CreateBoardsResponse{Ids: ids}, nil
}
