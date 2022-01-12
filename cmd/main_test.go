package main_test

import (
	"context"
	"github.com/apple/pkl-go-examples/gen/appconfig"
	"github.com/apple/pkl-go-examples/gen/redisconfig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadAppConfig(t *testing.T) {
	cfg, err := appconfig.LoadFromPath(context.Background(), "../pkl/dev/config.pkl")
	if assert.NoError(t, err) {
		assert.Equal(t, cfg.Port, uint16(5051))
		assert.Equal(t,
			cfg.Redis.Auth,
			&redisconfig.Auth{
				Username: "redis",
				Password: "redis",
			},
		)
	}
}
