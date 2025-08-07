package main

import (
	"log"

	"github.com/bogi-lyceya-44/task-tracker/config"
	tasks_api "github.com/bogi-lyceya-44/task-tracker/internal/app/api/tasks"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/bootstrap"
	tasks_service "github.com/bogi-lyceya-44/task-tracker/internal/app/services/tasks"
	tasks_storage "github.com/bogi-lyceya-44/task-tracker/internal/app/storages/tasks"
	"github.com/pkg/errors"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(errors.Wrap(err, "init config"))
	}

	bootstrap.InitCloser()
	ctx, err := bootstrap.InitGlobalContext()
	if err != nil {
		log.Fatal(errors.Wrap(err, "init global context"))
	}

	pool, err := bootstrap.InitPostgresPool(
		ctx,
		cfg.Postgres.URL,
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "init pool"))
	}

	tasksStorage := tasks_storage.NewTaskStorage(pool)
	taskService := tasks_service.NewTaskService(tasksStorage)
	impl := tasks_api.New(taskService)

	if err = bootstrap.RunApp(ctx, *cfg, impl); err != nil {
		log.Fatal(errors.Wrap(err, "running app"))
	}
}
