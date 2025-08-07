package topics

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s *TopicStorage) GetTopics(
	ctx context.Context,
	ids []int64,
) ([]models.Topic, error) {
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

	fetched, err := pgx.CollectRows(rows, pgx.RowToStructByName[Topic])
	if err != nil {
		return nil, errors.Wrap(err, "collecting rows")
	}

	return utils.Map(
		fetched,
		mapStorageTopicToDomain,
	), nil
}

func mapStorageTopicToDomain(topic Topic) models.Topic {
	return models.Topic{
		ID:        topic.ID,
		Name:      topic.Name,
		TaskIds:   topic.TaskIds,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
	}
}
