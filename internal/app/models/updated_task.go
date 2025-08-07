package models

import "time"

type UpdatedTask struct {
	ID           int64
	Name         *string
	Description  *string
	Dependencies []int64
	Priority     *Priority
	StartTime    *time.Time
	FinishTime   *time.Time
}
