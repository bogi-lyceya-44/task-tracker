package tasks

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s *TaskStorage) GetTasks(
	ctx context.Context,
	ids []int64,
) ([]models.Task, error) {
	sql, args, err := sq.
		Select(allColumns...).
		From(tableName).
		Where(map[string]any{
			columnID:        ids,
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

	fetched, err := pgx.CollectRows(rows, pgx.RowToStructByName[Task])
	if err != nil {
		return nil, errors.Wrap(err, "collecting rows")
	}

	return utils.Map(
		fetched,
		mapStorageTaskToDomain,
	), nil
}

func mapStorageTaskToDomain(task Task) models.Task {
	return models.Task{
		ID:           task.ID,
		Name:         task.Name,
		Description:  task.Description,
		Dependencies: task.Dependencies,
		Priority:     models.Priority(task.Priority),
		StartTime:    task.StartTime,
		FinishTime:   task.FinishTime,
		CreatedAt:    task.CreatedAt,
		UpdatedAt:    task.UpdatedAt,
	}
}
