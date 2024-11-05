// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package main

import (
	"context"

	orderServer "github.com/costa92/micros-service/internal/order-server/server"
	v1 "github.com/costa92/micros-service/pkg/api/order-server/v1"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func main() {

	opts := []grpc.ServerOption{
		grpc.Address(":9100"),
	}
	srv := grpc.NewServer(opts...)
	orderService := orderServer.NewOrderService()
	// Register gRPC server endpoint
	v1.RegisterOrderServiceServer(srv, orderService)

	// Register http server endpoint
	httpOpts := []http.ServerOption{
		http.Middleware(),
		http.Address(":8080"),
	}
	httpSrv := http.NewServer(httpOpts...)
	v1.RegisterOrderServiceHTTPServer(httpSrv, orderService)

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
