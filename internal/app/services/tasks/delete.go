package tasks

import "context"

func (s *TaskService) DeleteTasks(
	ctx context.Context,
	ids []int64,
) error {
	return s.storage.DeleteTasks(ctx, ids)
}
