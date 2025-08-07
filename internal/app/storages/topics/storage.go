package topics

import "github.com/jackc/pgx/v5/pgxpool"

type TopicStorage struct {
	pool *pgxpool.Pool
}

func NewTopicStorage(pool *pgxpool.Pool) *TopicStorage {
	return &TopicStorage{
		pool: pool,
	}
}
