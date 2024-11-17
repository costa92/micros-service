//go:build wireinject
// +build wireinject

package orderserver

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/costa92/micros-service/internal/orderserver/server"
	"github.com/costa92/micros-service/internal/orderserver/service" // Add this import
	"github.com/costa92/micros-service/internal/pkg/bootstrap"       // Add this import
	"github.com/costa92/micros-service/pkg/db"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func wireApp(
	bootstrap.AppInfo,
	*server.Config,
	*db.MySQLOptions,
) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		server.ProviderSet,
		service.ProviderSet,
	)
	return nil, nil, nil
}
