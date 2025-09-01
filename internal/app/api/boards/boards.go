package boards

import (
	"context"

	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/pkg/pb/api/boards"
)

type BoardService interface {
	GetBoards(ctx context.Context, ids []int64) ([]models.Board, error)
	InsertBoards(ctx context.Context, boards []models.Board) ([]int64, error)
	UpdateBoards(ctx context.Context, boards []models.UpdatedBoard) error
	DeleteBoards(ctx context.Context, ids []int64) error

	GetOrder(ctx context.Context) ([]utils.Pair[int64, int32], error)
	ChangeOrder(ctx context.Context, ids []int64, places []int32) error
}

type TopicService interface {
	GetTopics(ctx context.Context, ids []int64) ([]models.Topic, error)
}

type TaskService interface {
	GetTasks(ctx context.Context, ids []int64) ([]models.Task, error)
}

type Implementation struct {
	desc.UnimplementedBoardServiceServer

	boardService BoardService
	topicService TopicService
	taskService  TaskService
}

func New(
	boardService BoardService,
	topicService TopicService,
	taskService TaskService,
) *Implementation {
	return &Implementation{
		boardService: boardService,
		topicService: topicService,
		taskService:  taskService,
	}
}
