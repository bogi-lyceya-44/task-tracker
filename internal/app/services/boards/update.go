package boards

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

func (s *BoardService) UpdateBoards(
	ctx context.Context,
	boards []models.UpdatedBoard,
) error {
	return s.storage.UpdateBoards(ctx, boards)
}
