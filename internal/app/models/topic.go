package models

import "time"

type Topic struct {
	ID      int64
	Name    string
	TaskIds []int64

	CreatedAt time.Time
	UpdatedAt time.Time
}
