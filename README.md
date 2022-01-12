# pkl-go-example

This is a sample project that demonstrates Pkl <> Go integration.

This is a basic HTTP server that is configured via Pkl in the `pkl/` directory.

## Requirements

* Go <1.18
* Pkl <0.25.0

## Project structure

| Directory   | Description                   |
|-------------|-------------------------------|
| `pkl/`      | Pkl configuration sources     |
| `gen/`      | Generated Go sources from Pkl |
| `internal/` | Internal Go files             |
| `cmd/`      | Server entrypoint             |
| `scripts/`  | Scripts directory             |

## Codegen

To generate new Pkl sources for the AppConfig module, first install the Go code generator:

```bash
go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest
```

With that installed, generate Go sources via:

```bash
pkl-gen-go pkl/AppConfig.pkl
```

The code generator detects that the Go package for `AppConfig` is
`github.com/apple/pkl-go-example/gen/appconfig`, and the Go module name is
`github.com/apple/pkl-go-example` (via the go.mod file), and therefore places generated
sources in `gen/appconfig`.
