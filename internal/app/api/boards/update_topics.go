package boards

import (
	"context"
	"errors"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	board_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/boards"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateBoards(
	ctx context.Context,
	req *desc.UpdateBoardsRequest,
) (*desc.UpdateBoardsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	tasks, err := utils.MapWithError(
		req.BoardsToUpdate,
		MapUpdateBoardPrototypeToUpdatedBoard,
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "mapping boards: %v", err)
	}

	if err = i.boardService.UpdateBoards(ctx, tasks); err != nil {
		errorCode := codes.Internal

		if errors.Is(err, board_service.ErrTopicDoesNotExist) {
			errorCode = codes.InvalidArgument
		}

		return nil, status.Errorf(errorCode, "updating boards: %v", err)
	}

	return &desc.UpdateBoardsResponse{}, nil
}
