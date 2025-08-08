package bootstrap

import (
	"context"

	tasks_api "github.com/bogi-lyceya-44/task-tracker/internal/app/api/tasks"
	tasks_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/tasks"
	tasks_storage "github.com/bogi-lyceya-44/task-tracker/internal/app/storages/tasks"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitTaskService(
	ctx context.Context,
	app *App,
	pool *pgxpool.Pool,
) error {
	tasksStorage := tasks_storage.NewTaskStorage(pool)
	taskService := tasks_service.NewTaskService(tasksStorage)
	impl := tasks_api.New(taskService)

	desc.RegisterTaskServiceServer(app.grpcServer, impl)

	if err := desc.RegisterTaskServiceHandlerFromEndpoint(
		ctx,
		app.mux,
		app.grpcAddr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		return errors.Wrap(err, "registering task service gateway")
	}

	return nil
}
