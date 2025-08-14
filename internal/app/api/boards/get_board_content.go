package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	tasks_api "github.com/bogi-lyceya-44/task-tracker/internal/app/api/tasks"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/boards"
	desc_tasks "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	desc_topics "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/topics"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) GetBoardContent(
	ctx context.Context,
	req *desc.GetBoardContentRequest,
) (*desc.GetBoardContentResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validating: %v", err)
	}

	boardIds := utils.Unique(req.GetIds())

	boards, err := i.boardService.GetBoards(ctx, boardIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting boards: %v", err)
	}

	topicIds := utils.Unique(
		utils.Flatten(
			utils.Map(
				boards,
				func(board models.Board) []int64 {
					return board.TopicIds
				},
			),
		),
	)

	topics, err := i.topicService.GetTopics(ctx, topicIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting topics: %v", err)
	}

	taskIds := utils.Unique(
		utils.Flatten(
			utils.Map(
				topics,
				func(topic models.Topic) []int64 {
					return topic.TaskIds
				},
			),
		),
	)

	tasks, err := i.taskService.GetTasks(ctx, taskIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting tasks: %v", err)
	}

	topicsWithFetchedTasks, err := formTopicsWithFetchedTasks(topics, tasks)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "forming topics: %v", err)
	}

	result := make(map[int64]*desc.GetBoardContentResponse_BoardContent, len(boards))

	for _, board := range boards {
		result[board.ID] = &desc.GetBoardContentResponse_BoardContent{
			Topics: utils.Map(
				board.TopicIds,
				func(id int64) *desc_topics.TopicWithFetchedTasks {
					return topicsWithFetchedTasks[id]
				},
			),
		}
	}

	return &desc.GetBoardContentResponse{
		ContentById: result,
	}, nil
}

func formTopicsWithFetchedTasks(
	topics []models.Topic,
	tasks []models.Task,
) (map[int64]*desc_topics.TopicWithFetchedTasks, error) {
	mappedTasks, err := utils.MapWithError(
		tasks,
		tasks_api.MapDomainTaskToProto,
	)
	if err != nil {
		return nil, errors.Wrap(err, "mapping domain tasks to proto")
	}

	taskById := utils.GroupBySingle(mappedTasks, (*desc_tasks.Task).GetId)
	result := make(map[int64]*desc_topics.TopicWithFetchedTasks, len(topics))

	for _, topic := range topics {
		result[topic.ID] = &desc_topics.TopicWithFetchedTasks{
			Id:   topic.ID,
			Name: topic.Name,
			Tasks: utils.Map(
				topic.TaskIds,
				func(id int64) *desc_tasks.Task {
					return taskById[id]
				},
			),
			CreatedAt: timestamppb.New(topic.CreatedAt),
			UpdatedAt: timestamppb.New(topic.UpdatedAt),
		}
	}

	return result, nil
}
