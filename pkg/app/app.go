// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package app

import (
	"fmt"
	"os"

	"github.com/costa92/micros-service/pkg/log"
	"github.com/costa92/micros-service/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/component-base/cli"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/term"

	genericoptions "github.com/costa92/micros-service/pkg/options"
)

type App struct {
	name        string
	shortDesc   string
	description string
	run         RunFunc
	cmd         *cobra.Command
	args        cobra.PositionalArgs

	// +optional
	healthCheckFunc HealthCheckFunc

	// +optional
	options CliOptions

	// +optional
	silence bool

	// +optional
	noConfig bool

	// watching and re-reading config file
	// +optional
	watch bool
}

type RunFunc func() error

type HealthCheckFunc func() error

// Option is a function on the options for an App
type Option func(*App)

// WithOptions to open the application's function to read from the command line
func WithOptions(options CliOptions) Option {
	return func(app *App) {
		app.options = options
	}
}

// WithRunFunc to open the application's function to run
func WithRunFunc(run RunFunc) Option {
	return func(app *App) {
		app.run = run
	}
}

// WithDescription to open the application's function to set the description
func WithDescription(description string) Option {
	return func(app *App) {
		app.description = description
	}
}

// WithHealthCheckFunc to open the application's function to set the health check function
func WithHealthCheckFunc(fn HealthCheckFunc) Option {
	return func(app *App) {
		app.healthCheckFunc = fn
	}
}

// WithDefaultHealthCheckFunc to open the application's function to set the default health check function
func WithDefaultHealthCheckFunc() Option {
	fn := func() HealthCheckFunc {
		return func() error {
			go genericoptions.NewHealthOptions().ServeHealthCheck()
			return nil
		}
	}
	return WithHealthCheckFunc(fn())
}

func WithSilence() Option {
	return func(app *App) {
		app.silence = true
	}
}

func WithNoConfig() Option {
	return func(app *App) {
		app.noConfig = true
	}
}

// WithValidArgs to open the application's function to validate the arguments
func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(app *App) {
		app.args = args
	}
}

func WithDefaultValidArgs() Option {
	return func(app *App) {
		app.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		}
	}
}

// WithWatch to open the application's function to watch
func WithWatch() Option {
	return func(app *App) {
		app.watch = true
	}
}

// NewApp to create a new application
func NewApp(name, shortDesc string, opts ...Option) *App {
	app := &App{
		name:      name,
		shortDesc: shortDesc,
		run:       func() error { return nil },
	}

	for _, opt := range opts {
		opt(app)
	}

	app.buildCommand()

	return app
}

// buildCommand to build the command
func (app *App) buildCommand() {
	cmd := &cobra.Command{
		Use:   app.name,
		Short: app.shortDesc,
		Long:  app.description,
		RunE:  app.runCommand,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Args: app.args,
	}

	if !cmd.SilenceUsage {
		cmd.SilenceUsage = true
		cmd.SetFlagErrorFunc(func(c *cobra.Command, err error) error {
			// Re-enable usage printing.
			c.SilenceUsage = false
			return err
		})
	}

	// In all cases error printing is done below.
	cmd.SilenceErrors = true

	cmd.SetOutput(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true

	var fss cliflag.NamedFlagSets
	if app.options != nil {
		fss = app.options.Flags()
	}

	version.AddFlags(fss.FlagSet("global"))

	if !app.noConfig {
		AddConfigFlag(fss.FlagSet("global"), app.name, app.watch)
	}

	for _, f := range fss.FlagSets {
		cmd.Flags().AddFlagSet(f)
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, fss, cols)

	app.cmd = cmd
}

// Run is used to launch the application.
func (app *App) Run() {
	os.Exit(cli.Run(app.cmd))
}

func (app *App) runCommand(cmd *cobra.Command, args []string) error {
	version.PrintAndExitIfRequested(app.name)

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	if app.options != nil {
		if err := viper.Unmarshal(app.options); err != nil {
			return err
		}

		// set default options
		if err := app.options.Complete(); err != nil {
			return err
		}

		// validate options
		if err := app.options.Validate(); err != nil {
			return err
		}
	}
	// 初始化日志
	log.Init(logOptions())
	defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中

	if !app.silence {
		log.Infow("Starting application", "name", app.name, "version", version.Get().ToJSON())
		log.Infow("Golang settings", "GOGC", os.Getenv("GOGC"), "GOMAXPROCS", os.Getenv("GOMAXPROCS"), "GOTRACEBACK", os.Getenv("GOTRACEBACK"))
		if !app.noConfig {
			PrintConfig()
		} else if app.options != nil {
			cliflag.PrintFlags(cmd.Flags())
		}
	}

	if app.healthCheckFunc != nil {
		if err := app.healthCheckFunc(); err != nil {
			return err
		}
	}

	// run application
	return app.run()
}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回.
// 注意：`viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		EnableColor:       viper.GetBool("log.enable-color"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

func init() {
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "console")
	viper.SetDefault("log.output-paths", []string{"stdout"})
}
