package bootstrap

import (
	"context"
	"log"

	"github.com/bogi-lyceya-44/common/pkg/closer"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

func InitPostgresPool(
	ctx context.Context,
	connectionString string,
) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "init postgres")
	}

	if err = closer.AddCallback(
		CloserGroupConnections,
		func() error {
			log.Print("cancel postgres")
			pool.Close()
			return nil
		},
	); err != nil {
		return nil, errors.Wrap(err, "postgres callback")
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "readiness probe")
	}

	return pool, nil
}
