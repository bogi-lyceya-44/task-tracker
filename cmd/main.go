package main

import (
	"log"

	"github.com/bogi-lyceya-44/task-tracker/config"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/bootstrap"
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

	app := bootstrap.InitApp(cfg)

	if err = bootstrap.InitTaskService(ctx, app, pool); err != nil {
		log.Fatal(errors.Wrap(err, "init task service"))
	}

	if err = bootstrap.InitTopicService(ctx, app, pool); err != nil {
		log.Fatal(errors.Wrap(err, "init topic service"))
	}

	if err = bootstrap.InitBoardService(ctx, app, pool); err != nil {
		log.Fatal(errors.Wrap(err, "init topic service"))
	}

	if err = app.Run(ctx); err != nil {
		log.Fatal(errors.Wrap(err, "running app"))
	}
}
