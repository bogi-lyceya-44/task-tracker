package bootstrap

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/bogi-lyceya-44/common/pkg/closer"
	"github.com/bogi-lyceya-44/task-tracker/config"
	"github.com/bogi-lyceya-44/task-tracker/docs"
	"github.com/bogi-lyceya-44/task-tracker/internal/app/api/tasks"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"github.com/flowchartsman/swaggerui"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	ShutdownTimeoutInSeconds = 5
)

func RunApp(
	appCtx context.Context,
	cfg config.Config,
	taskService *tasks.Implementation,
) error {
	eg, ctx := errgroup.WithContext(appCtx)

	grpcAddr := net.JoinHostPort(cfg.GRPC.Host, cfg.GRPC.Port)
	lis, err := net.Listen(
		"tcp",
		grpcAddr,
	)
	if err != nil {
		return errors.Wrap(err, "start listening")
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	desc.RegisterTaskServiceServer(grpcServer, taskService)

	mux := runtime.NewServeMux()
	if err = desc.RegisterTaskServiceHandlerFromEndpoint(
		ctx,
		mux,
		grpcAddr,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	); err != nil {
		return errors.Wrap(err, "registering mux")
	}

	gatewayAddr := net.JoinHostPort(cfg.Gateway.Host, cfg.Gateway.Port)
	if err = mux.HandlePath(
		"GET",
		"/docs",
		func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
		},
	); err != nil {
		return errors.Wrap(err, "registering swagger json")
	}

	if err = mux.HandlePath(
		"GET",
		"/swagger/{path=**}",
		func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			http.StripPrefix("/swagger", swaggerui.Handler(docs.Spec)).ServeHTTP(w, r)
		},
	); err != nil {
		return errors.Wrap(err, "registering swagger json")
	}

	httpServer := http.Server{
		Addr:    gatewayAddr,
		Handler: mux,
	}

	eg.Go(func() error {
		if err = httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return errors.Wrap(err, "failed listening http")
		}

		return nil
	})

	eg.Go(func() error {
		if err = grpcServer.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			return errors.Wrap(err, "failed listening grpc")
		}

		return nil
	})

	if err = closer.AddCallback(
		CloserGroupApp,
		func() error {
			defer lis.Close()
			grpcServer.GracefulStop()

			shutdownCtx, cancel := context.WithTimeout(
				context.Background(),
				ShutdownTimeoutInSeconds*time.Second,
			)
			defer cancel()

			if err = httpServer.Shutdown(shutdownCtx); err != nil {
				return errors.Wrap(err, "shutting down http")
			}

			return nil
		},
	); err != nil {
		return errors.Wrap(err, "app callback")
	}

	log.Print("app started")

	err = eg.Wait()

	// wait for the parent context to close
	<-appCtx.Done()

	return err
}
