package tasks

import (
	"context"

	"github.com/pkg/errors"
)

func (s *TaskService) GetDependencies(
	ctx context.Context,
	ids []int64,
) (map[int64][]int64, error) {
	result := make(map[int64][]int64, len(ids))

	for _, id := range ids {
		fetched, err := s.storage.GetDependencies(ctx, id)
		if err != nil {
			return nil, errors.Wrap(err, "getting dependencies")
		}

		result[id] = fetched
	}

	return result, nil
}
