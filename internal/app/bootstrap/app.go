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
	"github.com/flowchartsman/swaggerui"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	ShutdownTimeoutInSeconds = 5
)

type App struct {
	grpcServer *grpc.Server
	mux        *runtime.ServeMux

	grpcAddr    string
	gatewayAddr string
}

func InitApp(cfg *config.Config) *App {
	grpcServer := grpc.NewServer()
	mux := runtime.NewServeMux()

	grpcAddr := net.JoinHostPort(cfg.GRPC.Host, cfg.GRPC.Port)
	gatewayAddr := net.JoinHostPort(cfg.Gateway.Host, cfg.Gateway.Port)

	reflection.Register(grpcServer)

	return &App{
		grpcServer:  grpcServer,
		mux:         mux,
		grpcAddr:    grpcAddr,
		gatewayAddr: gatewayAddr,
	}
}

func (app *App) Run(ctx context.Context) error {
	eg, _ := errgroup.WithContext(ctx)

	lis, err := net.Listen(
		"tcp",
		app.grpcAddr,
	)
	if err != nil {
		return errors.Wrap(err, "start listening")
	}

	if err = app.mux.HandlePath(
		"GET",
		"/docs",
		func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
		},
	); err != nil {
		return errors.Wrap(err, "registering swagger json")
	}

	if err = app.mux.HandlePath(
		"GET",
		"/swagger/{path=**}",
		func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			http.StripPrefix("/swagger", swaggerui.Handler(docs.Spec)).ServeHTTP(w, r)
		},
	); err != nil {
		return errors.Wrap(err, "registering swagger json")
	}

	httpServer := http.Server{
		Addr:    app.gatewayAddr,
		Handler: corsMiddleware(app.mux),
	}

	eg.Go(func() error {
		if err = httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return errors.Wrap(err, "failed listening http")
		}

		return nil
	})

	eg.Go(func() error {
		if err = app.grpcServer.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			return errors.Wrap(err, "failed listening grpc")
		}

		return nil
	})

	if err = closer.AddCallback(
		CloserGroupApp,
		func() error {
			defer lis.Close()
			app.grpcServer.GracefulStop()

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
	<-ctx.Done()

	return err
}

// allows any requests
// WARN: bad for production
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)

			return
		}

		next.ServeHTTP(w, r)
	})
}
