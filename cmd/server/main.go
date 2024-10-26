package main

import (
	"github.com/disco07/gRPC/generated/calculator"
	"github.com/disco07/gRPC/generated/users"
	"github.com/disco07/gRPC/internal"
	"github.com/disco07/grpc-lib/server"
	"go.uber.org/fx"
	"log/slog"
)

func newLogger() *slog.Logger {
	return slog.Default()
}

func newGRPCConfigServer() server.GRPCConfigServer {
	return server.YAMLGRPCConfigServer{
		ValuePort: 50051,
	}
}

var Module = fx.Options(
	server.Module,
	internal.Module,

	fx.Provide(
		newLogger, newGRPCConfigServer,
	),

	fx.Invoke(
		calculator.RegisterCalculatorServiceServer,
		users.RegisterUserServiceServer,
	),
)

func main() {
	app := fx.New(
		Module,
	)
	app.Run()
}
