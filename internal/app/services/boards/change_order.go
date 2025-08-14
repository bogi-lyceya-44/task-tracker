package boards

import (
	"context"
)

func (s *BoardService) ChangeOrder(
	ctx context.Context,
	ids []int64,
	places []int32,
) error {
	return s.boardStorage.ChangeOrder(ctx, ids, places)
}
