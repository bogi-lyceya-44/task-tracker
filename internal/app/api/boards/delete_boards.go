package boards

import (
	"context"

	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteBoards(
	ctx context.Context,
	req *desc.DeleteBoardsRequest,
) (*desc.DeleteBoardsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	err := i.boardService.DeleteBoards(ctx, req.Ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "deleting boards: %v", err)
	}

	return &desc.DeleteBoardsResponse{}, nil
}
