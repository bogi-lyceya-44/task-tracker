package topics

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *TopicStorage) UpdateTopics(
	ctx context.Context,
	topics []models.UpdatedTopic,
) error {
	for _, topic := range topics {
		updatedColumns := map[string]any{
			columnUpdatedAt: time.Now().UTC(),
		}

		if topic.Name != nil {
			updatedColumns[columnName] = *topic.Name
		}

		if topic.TaskIds != nil {
			updatedColumns[columnTaskIDs] = topic.TaskIds
		}

		sql, args, err := sq.
			Update(tableName).
			Where(
				sq.Eq{
					columnID: topic.ID,
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
