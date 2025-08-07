package topics

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s *TopicStorage) InsertTopics(
	ctx context.Context,
	topics []models.Topic,
) ([]int64, error) {
	builder := sq.
		Insert(tableName).
		Columns(
			columnName,
			columnTaskIDs,
			columnCreatedAt,
			columnUpdatedAt,
		).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	now := time.Now().UTC()

	for _, topic := range topics {
		builder = builder.Values(
			topic.Name,
			topic.TaskIds,
			now,
			now,
		)
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "building sql")
	}

	rows, err := s.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "querying rows")
	}

	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, errors.Wrap(err, "collecting rows")
	}

	return ids, nil
}
