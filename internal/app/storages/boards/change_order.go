package boards

import (
	"context"

	"github.com/pkg/errors"
)

func (s *BoardStorage) ChangeOrder(
	ctx context.Context,
	ids []int64,
	places []int32,
) error {
	sql := `
		INSERT INTO task_tracker.board_order
		SELECT
			board_id,
			place
		FROM UNNEST(
			$1::bigint[],
			$2::integer[]
		) AS source (
			board_id,
			place
		)
		ON CONFLICT ON CONSTRAINT board_order_pkey
		DO UPDATE
		SET place = COALESCE(EXCLUDED.place, task_tracker.board_order.place);
	`

	_, err := s.pool.Exec(ctx, sql, ids, places)
	if err != nil {
		return errors.Wrap(err, "fetching rows")
	}

	return nil
}
