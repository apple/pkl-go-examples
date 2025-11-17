//===----------------------------------------------------------------------===//
// Copyright © 2025 Apple Inc. and the Pkl project authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//===----------------------------------------------------------------------===//

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
