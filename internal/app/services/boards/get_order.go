package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
)

func (s *BoardService) GetOrder(
	ctx context.Context,
) ([]utils.Pair[int64, int32], error) {
	return s.boardStorage.GetOrder(ctx)
}
