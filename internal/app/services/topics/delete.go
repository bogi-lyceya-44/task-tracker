package topics

import "context"

func (s *TopicService) DeleteTopics(
	ctx context.Context,
	ids []int64,
) error {
	return s.topicStorage.DeleteTopics(ctx, ids)
}
