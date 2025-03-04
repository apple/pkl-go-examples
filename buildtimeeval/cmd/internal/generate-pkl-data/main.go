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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apple/pkl-go/pkl"
)

func main() {
	evaluator, err := pkl.NewEvaluator(context.Background(), pkl.PreconfiguredOptions)
	if err != nil {
		panic(err)
	}
	configFiles, err := filepath.Glob("pkl/*/config.pkl")
	if err != nil {
		panic(err)
	}
	for i := range configFiles {
		filename := configFiles[i]
		bytes, err := evaluator.EvaluateExpressionRaw(context.Background(), pkl.FileSource(filename), "module")
		if err != nil {
			panic(err)
		}
		path := "configdata/" + strings.Replace(filename, ".pkl", ".msgpack", 1)
		if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			panic(err)
		}
		if err = os.WriteFile(path, bytes, 0o644); err != nil {
			panic(err)
		}
		fmt.Printf("Generated %s", path)
	}
}
