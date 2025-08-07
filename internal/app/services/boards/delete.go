package boards

import "context"

func (s *BoardService) DeleteBoards(
	ctx context.Context,
	ids []int64,
) error {
	return s.storage.DeleteBoards(ctx, ids)
}
