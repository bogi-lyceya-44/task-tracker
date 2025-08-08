package boards

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

type boardStorage interface {
	GetBoards(ctx context.Context, ids []int64) ([]models.Board, error)
	InsertBoards(ctx context.Context, tasks []models.Board) ([]int64, error)
	UpdateBoards(ctx context.Context, tasks []models.UpdatedBoard) error
	DeleteBoards(ctx context.Context, ids []int64) error
}

type topicStorage interface {
	GetTopics(ctx context.Context, ids []int64) ([]models.Topic, error)
}

type BoardService struct {
	boardStorage boardStorage
	topicStorage topicStorage
}

func NewBoardService(
	boardStorage boardStorage,
	topicStorage topicStorage,
) *BoardService {
	return &BoardService{
		boardStorage: boardStorage,
		topicStorage: topicStorage,
	}
}
