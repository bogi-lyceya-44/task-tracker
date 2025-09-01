package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ChangeBoardOrder(
	ctx context.Context,
	req *desc.ChangeBoardOrderRequest,
) (*desc.ChangeBoardOrderResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	ids := utils.Map(
		req.Changes,
		(*desc.BoardOrder).GetBoardId,
	)

	places := utils.Map(
		req.Changes,
		(*desc.BoardOrder).GetPlace,
	)

	err := i.boardService.ChangeOrder(ctx, ids, places)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "changing order: %v", err)
	}

	return &desc.ChangeBoardOrderResponse{}, nil
}
