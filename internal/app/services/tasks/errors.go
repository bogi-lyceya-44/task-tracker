package tasks

import "errors"

var (
	ErrSelfDependent    = errors.New("found self dependency")
	ErrTaskDoesNotExist = errors.New("task does not exist")
	ErrCyclicDependency = errors.New("task has indirect cyclic dependency")
)
