package tasks

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

// GetDependencies fetches all (direct and indirect) dependencies of the task.
func (s *TaskStorage) GetDependencies(
	ctx context.Context,
	id int64,
) ([]int64, error) {
	sql := fmt.Sprintf(`
		WITH RECURSIVE dependencies AS (
			SELECT
				t.id,
				t.dependencies,
				'{}'::bigint[] AS path
			FROM task_tracker.tasks t
			WHERE t.id = %d
			UNION
			SELECT
				tt.id,
				tt.dependencies,
				d.path || tt.id
			FROM dependencies d
			CROSS JOIN LATERAL UNNEST(d.dependencies) AS dep(id)
			INNER JOIN task_tracker.tasks tt
				ON tt.id = dep.id
			WHERE NOT tt.id = ANY(d.path)
		)
		SELECT d.id
		FROM dependencies d
		GROUP BY d.id
		HAVING d.id <> %d OR COUNT(*) > 1;
		`,
		id,
		id,
	)

	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, "querying rows")
	}

	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, errors.Wrap(err, "collecting rows")
	}

	return ids, nil
}
