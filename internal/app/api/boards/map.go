package boards

import (
	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/boards"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapCreateBoardPrototypeToDomain(board *desc.CreateBoardsRequest_BoardPrototype) models.Board {
	return models.Board{
		Name:     board.GetName(),
		TopicIds: board.GetTopicIds(),
	}
}

func MapDomainBoardToProto(board models.Board) *desc.Board {
	return &desc.Board{
		Id:        board.ID,
		Name:      board.Name,
		TopicIds:  board.TopicIds,
		CreatedAt: timestamppb.New(board.CreatedAt),
		UpdatedAt: timestamppb.New(board.UpdatedAt),
	}
}

func MapUpdateBoardPrototypeToUpdatedBoard(board *desc.UpdateBoardsRequest_BoardPrototype) (models.UpdatedBoard, error) {
	result := models.UpdatedBoard{ID: board.GetId()}

	if board.Name != nil {
		result.Name = board.Name
	}

	if board.TopicIds != nil {
		result.TopicIds = append(result.TopicIds, board.TopicIds...)
	}

	return result, nil
}
