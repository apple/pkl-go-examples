// ===----------------------------------------------------------------------===//
// Copyright © 2024 Apple Inc. and the Pkl project authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ===----------------------------------------------------------------------===//
package internal

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func routes(s *server) {
	s.gin.GET("/", s.okay())
	s.gin.GET("/ping", s.ping())
}

func (s *server) okay() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
			"config": s.config,
		})
	}
}

func (s *server) ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := s.redis.Ping(context.Background()).Result()
		if err != nil {
			slog.Error("redis ping failed", "error", err)
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "OK",
		})
	}
}
