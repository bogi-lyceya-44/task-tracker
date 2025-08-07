package bootstrap

import (
	"context"

	boards_api "github.com/bogi-lyceya-44/task-tracker/internal/app/api/boards"
	boards_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/boards"
	boards_storage "github.com/bogi-lyceya-44/task-tracker/internal/app/storages/boards"
	topics_storage "github.com/bogi-lyceya-44/task-tracker/internal/app/storages/topics"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/boards"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitBoardService(
	ctx context.Context,
	app *App,
	pool *pgxpool.Pool,
) error {
	boardStorage := boards_storage.NewBoardStorage(pool)
	topicStorage := topics_storage.NewTopicStorage(pool)
	boardService := boards_service.NewBoardService(boardStorage, topicStorage)
	impl := boards_api.New(boardService)

	desc.RegisterBoardServiceServer(app.grpcServer, impl)

	if err := desc.RegisterBoardServiceHandlerFromEndpoint(
		ctx,
		app.mux,
		app.grpcAddr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		return errors.Wrap(err, "registering board service gateway")
	}

	return nil
}
