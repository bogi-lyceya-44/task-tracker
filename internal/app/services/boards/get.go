package boards

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

func (s *BoardService) GetBoards(
	ctx context.Context,
	ids []int64,
) ([]models.Board, error) {
	return s.boardStorage.GetBoards(ctx, ids)
}
