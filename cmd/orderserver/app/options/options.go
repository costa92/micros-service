package options

import (
	"github.com/costa92/micros-service/internal/orderserver"
	"github.com/costa92/micros-service/internal/pkg/client"
	"github.com/costa92/micros-service/pkg/app"
	"github.com/costa92/micros-service/pkg/log"
	genericoptions "github.com/costa92/micros-service/pkg/options"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	cliflag "k8s.io/component-base/cli/flag"
)

var _ app.CliOptions = (*Options)(nil)

type Options struct {
	// gRPC options for configuring gRPC related options.
	GRPCOptions *genericoptions.GRPCOptions `json:"grpc" mapstructure:"grpc"`
	// HTTP options for configuring HTTP related options.
	HTTPOptions *genericoptions.HTTPOptions `json:"http" mapstructure:"http"`

	// MySQL options for configuring MySQL database related options.
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	// Etcd options for configuring Etcd related options.
	EtcdOptions *genericoptions.EtcdOptions `json:"etcd" mapstructure:"etcd"`
	// Log options for configuring log related options.
	Log *log.Options `json:"log" mapstructure:"log"`
}

func NewOptions() *Options {
	o := &Options{
		GRPCOptions:  genericoptions.NewGRPCOptions(),
		HTTPOptions:  genericoptions.NewHTTPOptions(),
		MySQLOptions: genericoptions.NewMySQLOptions(),
		EtcdOptions:  genericoptions.NewEtcdOptions(),
		Log:          log.NewOptions(),
	}

	return o
}

func (o *Options) Complete() error {

	return nil

}

// ApplyTo fills up onex-usercenter config with options.
func (o *Options) ApplyTo(c *orderserver.Config) error {
	c.GRPCOptions = o.GRPCOptions
	c.HTTPOptions = o.HTTPOptions
	c.MySQLOptions = o.MySQLOptions
	c.EtcdOptions = o.EtcdOptions
	return nil

}

// Flags returns flags for a specific server by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.HTTPOptions.AddFlags(fss.FlagSet("http"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.EtcdOptions.AddFlags(fss.FlagSet("etcd"))
	o.Log.AddFlags(fss.FlagSet("log"))

	fs := fss.FlagSet("misc")

	client.AddFlags(fs)

	return fss

}

// Validate validates all the required options.
func (o *Options) Validate() error {
	errs := []error{}
	errs = append(errs, o.GRPCOptions.Validate()...)
	errs = append(errs, o.HTTPOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.EtcdOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)
	return utilerrors.NewAggregate(errs)
}

// Config return an onex-usercenter config object.
func (o *Options) Config() (*orderserver.Config, error) {
	c := &orderserver.Config{}

	if err := o.ApplyTo(c); err != nil {
		return nil, err
	}

	return c, nil
}
