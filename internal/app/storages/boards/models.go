package boards

import "time"

type Board struct {
	ID       int64   `db:"id"`
	Name     string  `db:"name"`
	TopicIDs []int64 `db:"topic_ids"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
