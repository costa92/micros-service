package orderserver

import (
	"os"

	"github.com/costa92/micros-service/internal/orderserver/server"
	"github.com/costa92/micros-service/internal/pkg/bootstrap"
	"github.com/costa92/micros-service/pkg/db"
	"github.com/costa92/micros-service/pkg/log"
	genericoptions "github.com/costa92/micros-service/pkg/options"
	"github.com/costa92/micros-service/pkg/version"
	"github.com/go-kratos/kratos/v2"
	"github.com/jinzhu/copier"
)

var (
	// Name is the name of the compiled software.
	Name = "order-server"

	// ID contains the host name and any error encountered during the retrieval.
	ID, _ = os.Hostname()
)

type Config struct {
	GRPCOptions  *genericoptions.GRPCOptions
	HTTPOptions  *genericoptions.HTTPOptions
	MySQLOptions *genericoptions.MySQLOptions
	EtcdOptions  *genericoptions.EtcdOptions
}

// The address on which the server should listen for requests.

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() completedConfig {
	return completedConfig{cfg}
}

// completedConfig holds the configuration after it has been completed.
type completedConfig struct {
	*Config
}

// New returns a new instance of Server from the given config.
func (c completedConfig) New(stopCh <-chan struct{}) (*Server, error) {

	appInfo := bootstrap.NewAppInfo(ID, Name, version.Get().String())

	conf := &server.Config{
		HTTP: *c.HTTPOptions,
		GRPC: *c.GRPCOptions,
	}

	var dbOptions db.MySQLOptions
	_ = copier.Copy(&dbOptions, c.MySQLOptions)

	// Initialize Kratos application with the provided configurations.
	app, cleanup, err := wireApp(appInfo, conf, &dbOptions)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	return &Server{app: app}, nil
}

// Server represents the server.
type Server struct {
	app *kratos.App
}

// Run is a method of the Server struct that starts the server.
func (s *Server) Run(stopCh <-chan struct{}) error {
	go func() {
		if err := s.app.Run(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-stopCh

	log.Infof("Gracefully shutting down server ...")

	if err := s.app.Stop(); err != nil {
		log.Errorw(err, "Failed to gracefully shutdown kratos application")
		return err
	}

	return nil
}
