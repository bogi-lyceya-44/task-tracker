package boards

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

func (s *BoardService) InsertBoards(
	ctx context.Context,
	boards []models.Board,
) ([]int64, error) {
	return s.storage.InsertBoards(ctx, boards)
}
