package boards

import (
	"context"

	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/boards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetBoardOrder(
	ctx context.Context,
	req *desc.GetBoardOrderRequest,
) (*desc.GetBoardOrderResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	order, err := i.boardService.GetOrder(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting order: %v", err)
	}

	result := make(map[int64]int32, len(order))

	for _, pair := range order {
		result[pair.First] = pair.Second
	}

	return &desc.GetBoardOrderResponse{Order: result}, nil
}
