package models

type UpdatedTopic struct {
	ID      int64
	Name    *string
	TaskIds []int64
}
