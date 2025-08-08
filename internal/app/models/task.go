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

func (t Task) GetID() int64 {
	return t.ID
}

func (t Task) GetDependencies() []int64 {
	return t.Dependencies
}
