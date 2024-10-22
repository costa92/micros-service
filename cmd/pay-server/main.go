package main

import (

	// Adjust the import path as necessary
	"context"

	payserver "github.com/costa92/micros-service/internal/pay-server/server"
	v1 "github.com/costa92/micros-service/pkg/api/pay-server/v1"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func main() {

	opts := []grpc.ServerOption{
		grpc.Address(":9000"),
	}
	srv := grpc.NewServer(opts...)
	payService := payserver.NewPayService()
	// Register gRPC server endpoint
	v1.RegisterPayServiceServer(srv, payService)

	// Register http server endpoint
	httpOpts := []http.ServerOption{
		http.Middleware(),
		http.Address(":8000"),
	}
	httpSrv := http.NewServer(httpOpts...)
	v1.RegisterPayServiceHTTPServer(httpSrv, payService)

	// Run both servers in parallel
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		if err := srv.Start(ctx); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := httpSrv.Start(ctx); err != nil {
			panic(err)
		}
	}()
	select {}

}
