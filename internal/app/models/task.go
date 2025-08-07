package models

import "time"

type Task struct {
	ID           int64
	Name         string
	Description  string
	Dependencies []int64
	Priority     Priority

	StartTime  time.Time
	FinishTime time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
