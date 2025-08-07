package models

import "time"

type Board struct {
	ID       int64
	Name     string
	TopicIDs []int64

	CreatedAt time.Time
	UpdatedAt time.Time
}
