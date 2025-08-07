package boards

import "github.com/jackc/pgx/v5/pgxpool"

type BoardStorage struct {
	pool *pgxpool.Pool
}

func NewBoardStorage(pool *pgxpool.Pool) *BoardStorage {
	return &BoardStorage{
		pool: pool,
	}
}
