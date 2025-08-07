package topics

import "time"

type Topic struct {
	ID      int64   `db:"id"`
	Name    string  `db:"name"`
	TaskIDs []int64 `db:"task_ids"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
