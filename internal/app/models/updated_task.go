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

func (t UpdatedTask) GetID() int64 {
	return t.ID
}

func (t UpdatedTask) GetDependencies() []int64 {
	return t.Dependencies
}
