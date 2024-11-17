package server

import (
	"github.com/costa92/micros-service/internal/orderserver/service"
	v1 "github.com/costa92/micros-service/pkg/api/orderserver/v1"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(c *Config, uc *service.OrderService, middlewares []middleware.Middleware) *grpc.Server {
	// Set the middleware options for the server.
	opts := []grpc.ServerOption{
		// grpc.WithDiscovery(nil),
		// grpc.WithEndpoint("discovery:///matrix.creation.service.grpc"),
		// Define the middleware chain with variable options.
		grpc.Middleware(middlewares...),
	}
	if c.GRPC.Network != "" {
		opts = append(opts, grpc.Network(c.GRPC.Network))
	}
	if c.GRPC.Timeout != 0 {
		opts = append(opts, grpc.Timeout(c.GRPC.Timeout))
	}
	if c.GRPC.Addr != "" {
		opts = append(opts, grpc.Address(c.GRPC.Addr))
	}
	// TODO: Need an elegant way to determine whether to open
	/*
		if c.TLS.UseTLS {
			opts = append(opts, grpc.TLSConfig(c.TLS.MustTLSConfig()))
		}
	*/

	// Create a new gRPC server with the middleware options.
	srv := grpc.NewServer(opts...)
	v1.RegisterOrderServerServer(srv, uc)
	return srv
}
