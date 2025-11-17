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

package internal

import (
	"fmt"

	"github.com/apple/pkl-go-examples/buildtimeeval/gen/appconfig"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type server struct {
	gin    *gin.Engine
	config *appconfig.AppConfig
	redis  *redis.Client
}

type Server interface {
	Run() error
}

var _ Server = (*server)(nil)

func (s server) Run() error {
	return s.gin.Run(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
}

func NewServer(config *appconfig.AppConfig) Server {
	s := &server{
		gin:    gin.Default(),
		config: config,
		redis: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
			Password: config.Redis.Auth.Password,
			Username: config.Redis.Auth.Username,
		}),
	}
	routes(s)
	return s
}
