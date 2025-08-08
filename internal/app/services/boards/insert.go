package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *BoardService) InsertBoards(
	ctx context.Context,
	boards []models.Board,
) ([]int64, error) {
	allTopicIds := utils.Map(
		boards,
		func(board models.Board) []int64 {
			return board.TopicIds
		},
	)

	if err := s.checkForTopicExistence(ctx, allTopicIds); err != nil {
		return nil, errors.Wrap(err, "checking for task existence")
	}

	return s.boardStorage.InsertBoards(ctx, boards)
}
