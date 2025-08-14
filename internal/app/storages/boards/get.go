package boards

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s *BoardStorage) GetBoards(
	ctx context.Context,
	ids []int64,
) ([]models.Board, error) {
	filters := map[string]any{
		columnIsDeleted: false,
	}

	if len(ids) > 0 {
		filters[columnID] = ids
	}

	sql, args, err := sq.
		Select(allColumns...).
		From(boardsTableName).
		Where(filters).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "building sql")
	}

	rows, err := s.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "fetching rows")
	}

	fetched, err := pgx.CollectRows(rows, pgx.RowToStructByName[Board])
	if err != nil {
		return nil, errors.Wrap(err, "collecting rows")
	}

	return utils.Map(
		fetched,
		mapStorageBoardToDomain,
	), nil
}

func mapStorageBoardToDomain(board Board) models.Board {
	return models.Board{
		ID:        board.ID,
		Name:      board.Name,
		TopicIds:  board.TopicIds,
		CreatedAt: board.CreatedAt,
		UpdatedAt: board.UpdatedAt,
	}
}
