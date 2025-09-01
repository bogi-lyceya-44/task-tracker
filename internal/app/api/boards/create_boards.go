package boards

import (
	"context"
	"errors"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	board_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/boards"
	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/boards"
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
		errorCode := codes.Internal

		if errors.Is(err, board_service.ErrTopicDoesNotExist) {
			errorCode = codes.InvalidArgument
		}

		return nil, status.Errorf(errorCode, "inserting boards: %v", err)
	}

	return &desc.CreateBoardsResponse{Ids: ids}, nil
}
