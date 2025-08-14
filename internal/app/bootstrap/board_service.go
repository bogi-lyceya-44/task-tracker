package bootstrap

import (
	"context"

	boards_api "github.com/bogi-lyceya-44/task-tracker/internal/app/api/boards"
	boards_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/boards"
	tasks_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/tasks"
	topics_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/topics"
	boards_storage "github.com/bogi-lyceya-44/task-tracker/internal/app/storages/boards"
	tasks_storage "github.com/bogi-lyceya-44/task-tracker/internal/app/storages/tasks"
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
	taskStorage := tasks_storage.NewTaskStorage(pool)

	topicService := topics_service.NewTopicService(topicStorage, taskStorage)
	taskService := tasks_service.NewTaskService(taskStorage)
	boardService := boards_service.NewBoardService(boardStorage, topicStorage)

	impl := boards_api.New(boardService, topicService, taskService)

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
