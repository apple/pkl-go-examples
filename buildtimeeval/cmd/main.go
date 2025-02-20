// ===----------------------------------------------------------------------===//
// Copyright © 2025 Apple Inc. and the Pkl project authors. All rights reserved.
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
package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/apple/pkl-go-examples/buildtimeeval/configdata"
	"github.com/apple/pkl-go-examples/buildtimeeval/gen/appconfig"
	"github.com/apple/pkl-go-examples/buildtimeeval/gen/appconfig/loglevel"
	"github.com/apple/pkl-go-examples/buildtimeeval/internal"
	"github.com/apple/pkl-go/pkl"
)

// Loads the app config from Pkl's byte representation.
//
// This relies on the application's configuration to be pre-baked at build time, and embedded via the `go:embed`
// directive.
//
// With the bytes, we can then unmarshal them into the AppConfig struct.
func LoadAppConfig() (*appconfig.AppConfig, error) {
	var cfg appconfig.AppConfig
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	bytes, err := configdata.EmbedFs.ReadFile(fmt.Sprintf("pkl/%s/config.msgpack", env))
	if err != nil {
		return nil, err
	}
	err = pkl.Unmarshal(bytes, &cfg)
	return &cfg, err
}

func main() {
	cfg, err := LoadAppConfig()
	if err != nil {
		panic(err)
	}

	var programLevel = new(slog.LevelVar)

	switch cfg.LogLevel {
	case loglevel.Info:
		programLevel.Set(slog.LevelInfo)
	case loglevel.Warn:
		programLevel.Set(slog.LevelWarn)
	case loglevel.Error:
		programLevel.Set(slog.LevelError)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel}))
	slog.SetDefault(logger)

	slog.Info("Starting server", "port", cfg.Port)

	if err = internal.NewServer(cfg).Run(); err != nil {
		panic(err)
	}
}
