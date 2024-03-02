// Code generated from Pkl module `org.pkl.golang.example.AppConfig`. DO NOT EDIT.
package appconfig

import (
	"context"

	"github.com/apple/pkl-go-examples/gen/appconfig/loglevel"
	"github.com/apple/pkl-go-examples/gen/redisconfig"
	"github.com/apple/pkl-go/pkl"
)

// Application configuration for the Pkl Go Example application.
//
// See generated sources in the gen/ directory.
type AppConfig struct {
	// The host to listen on
	Host string `pkl:"host"`

	// The application port to listen on
	Port uint16 `pkl:"port"`

	// Redis settings for this application
	Redis *redisconfig.RedisConfig `pkl:"redis"`

	// The level of logging for the application.
	//
	// - "error": Log only error level messages
	// - "warn": Log error and warning messages
	// - "info": Log all messages
	LogLevel loglevel.LogLevel `pkl:"logLevel"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a AppConfig
func LoadFromPath(ctx context.Context, path string) (ret *AppConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a AppConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*AppConfig, error) {
	var ret AppConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
