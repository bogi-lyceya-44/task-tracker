package topics

import "context"

func (s *TopicService) DeleteTopics(
	ctx context.Context,
	ids []int64,
) error {
	return s.storage.DeleteTopics(ctx, ids)
}
