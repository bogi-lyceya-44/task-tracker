package tasks

import "errors"

var (
	ErrSelfDependent    = errors.New("found self dependency")
	ErrTaskDoesNotExist = errors.New("task does not exist")
)
