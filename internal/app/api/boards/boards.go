package boards

import (
	"context"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/boards"
)

type BoardService interface {
	GetBoards(ctx context.Context, ids []int64) ([]models.Board, error)
	InsertBoards(ctx context.Context, boards []models.Board) ([]int64, error)
	UpdateBoards(ctx context.Context, boards []models.UpdatedBoard) error
	DeleteBoards(ctx context.Context, ids []int64) error
}

type Implementation struct {
	desc.UnimplementedBoardServiceServer

	boardService BoardService
}

func New(boardService BoardService) *Implementation {
	return &Implementation{
		boardService: boardService,
	}
}
