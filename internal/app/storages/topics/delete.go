package topics

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (s *TopicStorage) DeleteTopics(
	ctx context.Context,
	ids []int64,
) error {
	sql, args, err := sq.
		Update(tableName).
		Set(columnIsDeleted, true).
		Where(map[string]any{
			columnID: ids,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "building sql")
	}

	if _, err := s.pool.Exec(ctx, sql, args...); err != nil {
		return errors.Wrap(err, "executing query")
	}

	return nil
}
