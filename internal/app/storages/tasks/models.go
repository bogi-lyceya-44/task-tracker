package tasks

import "time"

type Task struct {
	ID           int64   `db:"id"`
	Name         string  `db:"name"`
	Description  string  `db:"description"`
	Dependencies []int64 `db:"dependencies"`
	Priority     int32   `db:"priority"`

	StartTime  time.Time `db:"start_time"`
	FinishTime time.Time `db:"finish_time"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
