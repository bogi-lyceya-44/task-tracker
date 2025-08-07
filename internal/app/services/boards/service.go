package boards

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
)

type storage interface {
	GetBoards(ctx context.Context, ids []int64) ([]models.Board, error)
	InsertBoards(ctx context.Context, tasks []models.Board) ([]int64, error)
	UpdateBoards(ctx context.Context, tasks []models.UpdatedBoard) error
	DeleteBoards(ctx context.Context, ids []int64) error
}

type BoardService struct {
	storage storage
}

func NewBoardService(storage storage) *BoardService {
	return &BoardService{
		storage: storage,
	}
}
