package main

import (
	"context"
	"fmt"
	"github.com/disco07/gRPC/generated/users"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"

	"github.com/disco07/gRPC/generated/calculator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClientConn(lc fx.Lifecycle) (*grpc.ClientConn, error) {
	grpcAddr := "localhost:50051"

	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not connect to order service: %w", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			fmt.Println("Closing gRPC connection")
			return conn.Close()
		},
	})

	return conn, nil
}

func NewServeMux() *runtime.ServeMux {
	return runtime.NewServeMux()
}

func NewContext() context.Context {
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("key", "value"))
	return ctx
}

func StartHTTPServer(lc fx.Lifecycle, mux *runtime.ServeMux) {
	addr := "0.0.0.0:8080"
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Println("API gateway server is running on " + addr)
				if err := http.ListenAndServe(addr, mux); err != nil {
					log.Fatalf("gateway server closed abruptly: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping HTTP server")
			return nil
		},
	})
}

var Module = fx.Options(
	fx.Provide(
		NewGRPCClientConn,
		NewServeMux,
		NewContext,
	),
	fx.Invoke(
		calculator.RegisterCalculatorServiceHandler,
		users.RegisterUserServiceHandler,
		StartHTTPServer,
	),
)

func main() {
	app := fx.New(Module)
	app.Run()
}
