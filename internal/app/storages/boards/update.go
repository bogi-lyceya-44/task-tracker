package boards

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *BoardStorage) UpdateBoards(
	ctx context.Context,
	boards []models.UpdatedBoard,
) error {
	for _, board := range boards {
		updatedColumns := map[string]any{
			columnUpdatedAt: time.Now().UTC(),
		}

		if board.Name != nil {
			updatedColumns[columnName] = *board.Name
		}

		if board.TopicIDs != nil {
			updatedColumns[columnTopicIDs] = board.TopicIDs
		}

		sql, args, err := sq.
			Update(tableName).
			Where(
				sq.Eq{
					columnID: board.ID,
				},
			).
			SetMap(updatedColumns).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return errors.Wrap(err, "building sql")
		}

		if _, err = s.pool.Exec(ctx, sql, args...); err != nil {
			return errors.Wrap(err, "executing query")
		}
	}

	return nil
}
