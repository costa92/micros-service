package app

import (
	"github.com/costa92/micros-service/cmd/orderserver/app/options"
	orderserver "github.com/costa92/micros-service/internal/orderserver"
	"github.com/costa92/micros-service/pkg/app"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

const commandDesc = `The order server is a microservice that provides an API for managing orders.`

func NewApp() *app.App {
	opts := options.NewOptions()
	application := app.NewApp(orderserver.Name, "Launch order server",
		app.WithDescription(commandDesc),
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application

}

// Returns the function to run the application.
func run(opts *options.Options) app.RunFunc {
	return func() error {
		cfg, err := opts.Config()
		if err != nil {
			return err
		}

		return Run(cfg, genericapiserver.SetupSignalHandler())
	}
}

// Run runs the specified APIServer. This should never exit.
func Run(c *orderserver.Config, stopCh <-chan struct{}) error {
	server, err := c.Complete().New(stopCh)
	if err != nil {
		return err
	}

	return server.Run(stopCh)
}
