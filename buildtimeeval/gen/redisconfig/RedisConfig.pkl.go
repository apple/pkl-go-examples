// Code generated from Pkl module `org.pkl.golang.example.RedisConfig`. DO NOT EDIT.
package redisconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type RedisConfig struct {
	// Whether Redis connections are enabled or not.
	Enabled bool `pkl:"enabled"`

	// The hostname that Redis listens on
	Host string `pkl:"host"`

	// The port that Redis listens on
	Port uint16 `pkl:"port"`

	// Authorization settings for Redis
	Auth *Auth `pkl:"auth"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a RedisConfig
func LoadFromPath(ctx context.Context, path string) (ret *RedisConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a RedisConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*RedisConfig, error) {
	var ret RedisConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
