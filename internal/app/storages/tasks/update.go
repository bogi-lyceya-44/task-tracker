package tasks

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/pkg/errors"
)

func (s *TaskStorage) UpdateTasks(
	ctx context.Context,
	tasks []models.UpdatedTask,
) error {
	for _, task := range tasks {
		updatedColumns := map[string]any{
			columnUpdatedAt: time.Now().UTC(),
		}

		if task.Name != nil {
			updatedColumns[columnName] = *task.Name
		}

		if task.Description != nil {
			updatedColumns[columnDescription] = *task.Description
		}

		if task.Dependencies != nil {
			updatedColumns[columnDependencies] = task.Dependencies
		}

		if task.Priority != nil {
			updatedColumns[columnPriority] = *task.Priority
		}

		if task.StartTime != nil {
			updatedColumns[columnStartTime] = *task.StartTime
		}

		if task.FinishTime != nil {
			updatedColumns[columnFinishTime] = *task.FinishTime
		}

		sql, args, err := sq.
			Update(tableName).
			Where(
				sq.Eq{
					columnID: task.ID,
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
