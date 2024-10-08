package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/disco07/gRPC/generated/calculator"
	"github.com/disco07/gRPC/internal"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var Module = fx.Options(
	internal.Module,
	fx.Provide(
		newLogger,
		NewGPRCServer,
	),
	fx.Invoke(
		calculator.RegisterCalculatorServiceServer,
	),
)

func newLogger() *slog.Logger {
	return slog.Default()
}

func NewGPRCServer(lifecycle fx.Lifecycle, logger *slog.Logger) grpc.ServiceRegistrar {
	const addr = ":50051"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Warn(err.Error())
	}

	keepaliveOptions := grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    time.Minute,
		Timeout: 3 * time.Second,
	})

	keepaliveEnforcementOptions := grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		MinTime:             30 * time.Second,
		PermitWithoutStream: true,
	})

	server := grpc.NewServer(
		keepaliveOptions,
		keepaliveEnforcementOptions,
	)

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				reflection.Register(server)

				if err = server.Serve(listener); err != nil {
					logger.Warn(err.Error())
				}
			}()
			logger.With(fmt.Sprintf("%s://%s", listener.Addr().Network(), listener.Addr().String())).Info("Listening")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()

			return nil
		},
	})

	return server
}

func main() {
	app := fx.New(
		Module,
	)
	app.Run()
}
