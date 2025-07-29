package internal_test

import (
	"testing"

	"github.com/apple/pkl-go-examples/simple/gen/redisconfig"
	"github.com/apple/pkl-go-examples/simple/internal"
	"github.com/stretchr/testify/assert"
)

func TestLoadAppConfig(t *testing.T) {
	cfg := internal.LoadAppConfig()
	assert.Equal(t, cfg.Port, uint16(5051))
	assert.Equal(t,
		cfg.Redis.Auth,
		&redisconfig.Auth{
			Username: "redis",
			Password: "redis",
		},
	)
}
