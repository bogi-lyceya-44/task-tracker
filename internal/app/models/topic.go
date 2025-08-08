package models

import "time"

type Topic struct {
	ID      int64
	Name    string
	TaskIds []int64

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t Topic) GetID() int64 {
	return t.ID
}
