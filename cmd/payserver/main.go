// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package main

import (

	// Adjust the import path as necessary
	"context"

	payserver "github.com/costa92/micros-service/internal/payserver/server"
	v1 "github.com/costa92/micros-service/pkg/api/payserver/v1"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func main() {

	opts := []grpc.ServerOption{
		grpc.Address(":9000"),
	}
	srv := grpc.NewServer(opts...)
	payServer := payserver.NewPayServer()
	// Register gRPC server endpoint
	v1.RegisterPayServerServer(srv, payServer)

	// Register http server endpoint
	httpOpts := []http.ServerOption{
		http.Middleware(),
		http.Address(":8000"),
	}
	httpSrv := http.NewServer(httpOpts...)
	v1.RegisterPayServerHTTPServer(httpSrv, payServer)

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
