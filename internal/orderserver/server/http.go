package server

import (
	"context"

	service "github.com/costa92/micros-service/internal/orderserver/service"
	"github.com/costa92/micros-service/internal/pkg/pprof"
	v1 "github.com/costa92/micros-service/pkg/api/orderserver/v1"
	kratos_middleware "github.com/costa92/micros-service/pkg/middleware/kratos"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whitelist := make(map[string]struct{})
	return func(ctx context.Context, operation string) bool {
		if _, ok := whitelist[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer creates a new HTTP server with middleware and handler chain.
func NewHTTPServer(c *Config, gw *service.OrderService, middlewares []middleware.Middleware) *http.Server {
	// Define the server options with the middleware chain and other configuration.
	opts := []http.ServerOption{
		// http.WithDiscovery(nil),
		// http.WithEndpoint("discovery:///matrix.creation.service.grpc"),
		// Define the middleware chain with variable options.
		http.Middleware(middlewares...),
		// Add filter options to the middleware chain.
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{
				"X-Requested-With",
				"Content-Type",
				"Authorization",
				"X-Idempotent-ID",
			}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowCredentials(),
		)),
	}
	if c.HTTP.Network != "" {
		opts = append(opts, http.Network(c.HTTP.Network))
	}
	if c.HTTP.Timeout != 0 {
		opts = append(opts, http.Timeout(c.HTTP.Timeout))
	}
	if c.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.HTTP.Addr))
	}
	if c.TLS.UseTLS {
		opts = append(opts, http.TLSConfig(c.TLS.MustTLSConfig()))
	}

	opts = append(opts, kratos_middleware.ResponseEncoder())
	// Create and return the server instance.
	srv := http.NewServer(opts...)
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/openapi/", h)
	srv.Handle("/metrics", promhttp.Handler())
	srv.Handle("", pprof.NewHandler())

	v1.RegisterOrderServerHTTPServer(srv, gw)
	return srv
}
