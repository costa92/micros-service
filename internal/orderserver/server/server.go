// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package server

import (
	krtlog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewServers, NewGRPCServer, NewHTTPServer, NewMiddlewares)

// NewServers is a wire provider function that creates and returns a slice of transport servers.
func NewServers(hs *http.Server, gs *grpc.Server) []transport.Server {
	return []transport.Server{hs, gs}
}

// NewMiddlewares return middlewares used by grpc and http server both.
func NewMiddlewares(logger krtlog.Logger) []middleware.Middleware {
	return []middleware.Middleware{
		logging.Server(logger),
	}
}
