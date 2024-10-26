package main

import (
	"context"
	"github.com/disco07/gRPC/generated/calculator"
	"github.com/disco07/gRPC/generated/users"
	"github.com/disco07/grpc-lib/client"
	"github.com/disco07/grpc-lib/server"
	"go.uber.org/fx"
	"google.golang.org/grpc/metadata"
)

func newContext() context.Context {
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("key", "value"))
	return ctx

}

func newGRPCConfigServer() server.GRPCConfigServer {
	return server.YAMLGRPCConfigServer{
		ValuePort: 50051,
	}
}

func newGRPCConfigClient() client.GRPCConfigClient {
	return client.YAMLGRPCConfigClient{
		ValuePort: 8090,
	}
}

var Module = fx.Options(
	client.Module,
	fx.Provide(newGRPCConfigClient, newContext, newGRPCConfigServer),
	fx.Invoke(
		calculator.RegisterCalculatorServiceHandler,
		users.RegisterUserServiceHandler,
	),
)

func main() {
	app := fx.New(
		Module,
	)
	app.Run()
}
