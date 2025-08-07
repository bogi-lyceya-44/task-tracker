package tasks

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s *TaskStorage) InsertTasks(
	ctx context.Context,
	tasks []models.Task,
) ([]int64, error) {
	builder := sq.
		Insert(tableName).
		Columns(
			columnName,
			columnDescription,
			columnDependencies,
			columnPriority,
			columnStartTime,
			columnFinishTime,
			columnCreatedAt,
			columnUpdatedAt,
		).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	now := time.Now().UTC()

	for _, task := range tasks {
		builder = builder.Values(
			task.Name,
			task.Description,
			task.Dependencies,
			task.Priority,
			task.StartTime,
			task.FinishTime,
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
