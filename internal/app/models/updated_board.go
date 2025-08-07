package models

type UpdatedBoard struct {
	ID       int64
	Name     *string
	TopicIDs []int64
}
