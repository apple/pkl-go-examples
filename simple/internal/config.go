package internal

import (
	"context"
	"net/url"
	"path/filepath"
	"runtime"

	"github.com/apple/pkl-go-examples/simple/gen/appconfig"
	"github.com/apple/pkl-go/pkl"
)

// LoadAppConfig loads an app config given the configured evaluator, with the set project dir.
// Alternatively, there's also the shorthand `appconfig.LoadFromPath` which uses only the preconfigured evaluator.
func LoadAppConfig() appconfig.AppConfig {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error: could not get caller information.")
	}
	path := filepath.Join(filepath.Dir(filename), "..")
	u := &url.URL{Path: path, Scheme: "file"}
	evaluator, err := pkl.NewProjectEvaluator(context.Background(), u, pkl.PreconfiguredOptions)
	if err != nil {
		panic(err)
	}
	defer evaluator.Close()
	cfg, err := appconfig.Load(context.Background(), evaluator, pkl.FileSource(filename, "../../pkl/dev/config.pkl"))
	if err != nil {
		panic(err)
	}
	return cfg
}
