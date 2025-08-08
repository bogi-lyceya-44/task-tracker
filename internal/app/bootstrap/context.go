package bootstrap

import (
	"context"
	"log"

	"github.com/bogi-lyceya-44/common/pkg/closer"
	"github.com/pkg/errors"
)

func InitGlobalContext() (context.Context, error) {
	ctx, cancel := context.WithCancel(context.Background())

	if err := closer.AddCallback(
		CloserGroupGlobalContext,
		func() error {
			log.Print("cancel context")
			cancel()
			return nil
		},
	); err != nil {
		return nil, errors.Wrap(err, "context callback")
	}

	return ctx, nil
}
