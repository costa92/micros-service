package server

import genericoptions "github.com/costa92/micros-service/pkg/options"

type Config struct {
	HTTP genericoptions.HTTPOptions
	GRPC genericoptions.GRPCOptions
	TLS  genericoptions.TLSOptions
}
