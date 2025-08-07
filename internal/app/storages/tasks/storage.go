package tasks

import "github.com/jackc/pgx/v5/pgxpool"

type TaskStorage struct {
	pool *pgxpool.Pool
}

func NewTaskStorage(pool *pgxpool.Pool) *TaskStorage {
	return &TaskStorage{
		pool: pool,
	}
}
