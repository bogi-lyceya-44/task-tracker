package bootstrap

import (
	"context"

	topics_api "github.com/bogi-lyceya-44/task-tracker/internal/app/api/topics"
	topics_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/topics"
	tasks_storage "github.com/bogi-lyceya-44/task-tracker/internal/app/storages/tasks"
	topics_storage "github.com/bogi-lyceya-44/task-tracker/internal/app/storages/topics"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/topics"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitTopicService(
	ctx context.Context,
	app *App,
	pool *pgxpool.Pool,
) error {
	topicStorage := topics_storage.NewTopicStorage(pool)
	tasksStorage := tasks_storage.NewTaskStorage(pool)
	topicService := topics_service.NewTopicService(topicStorage, tasksStorage)
	impl := topics_api.New(topicService)

	desc.RegisterTopicServiceServer(app.grpcServer, impl)

	if err := desc.RegisterTopicServiceHandlerFromEndpoint(
		ctx,
		app.mux,
		app.grpcAddr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		return errors.Wrap(err, "registering topic service gateway")
	}

	return nil
}
