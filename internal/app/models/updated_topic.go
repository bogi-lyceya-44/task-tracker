package models

type UpdatedTopic struct {
	ID      int64
	Name    *string
	TaskIDs []int64
}
