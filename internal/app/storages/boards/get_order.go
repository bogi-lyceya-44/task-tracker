package boards

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s *BoardStorage) GetOrder(ctx context.Context) ([]utils.Pair[int64, int32], error) {
	sql, args, err := sq.
		Select(columnID, columnPlace).
		From(orderTableName).
		InnerJoin(fmt.Sprintf("%s USING (%s)", boardsTableName, columnID)).
		Where(map[string]any{
			columnIsDeleted: false,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "building sql")
	}

	rows, err := s.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "fetching rows")
	}

	fetched, err := pgx.CollectRows(rows, pgx.RowToStructByPos[utils.Pair[int64, int32]])
	if err != nil {
		return nil, errors.Wrap(err, "collecting rows")
	}

	return fetched, nil
}
