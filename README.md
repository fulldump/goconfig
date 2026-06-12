# Goconfig

![Logo](logo.png)

<p align="center">
  <a href="https://github.com/fulldump/goconfig/actions/workflows/ci.yml"><img src="https://github.com/fulldump/goconfig/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://pkg.go.dev/github.com/fulldump/goconfig"><img src="https://pkg.go.dev/badge/github.com/fulldump/goconfig.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/fulldump/goconfig"><img src="https://goreportcard.com/badge/github.com/fulldump/goconfig" alt="Go Report Card"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT License"></a>
</p>

`goconfig` is a lightweight Go library that fills structs from:

1. JSON config file(s)
2. Environment variables
3. Command-line flags

It is designed for apps that want production-ready configuration with minimal
boilerplate.

## Installation

```bash
go get github.com/fulldump/goconfig
```

## Quick Start

```go
package main

import (
	"log"
	"time"

	"github.com/fulldump/goconfig"
)

type DB struct {
	Host string `usage:"Database host"`
	Port int    `usage:"Database port"`
}

type Config struct {
	ServiceName string        `usage:"Service name"`
	Timeout     time.Duration `usage:"Request timeout"`
	DB          DB
}

func main() {
	cfg := Config{
		ServiceName: "payments",
		Timeout:     5 * time.Second,
		DB: DB{
			Host: "localhost",
			Port: 5432,
		},
	}

	if err := goconfig.Load(&cfg); err != nil {
		log.Fatal(err)
	}
}
```

If you want legacy one-liner behaviour (exit on error):

```go
goconfig.Read(&cfg)
```

## Copy/Paste Recipes

### 1) API service

```go
type Config struct {
	HTTPPort int           `usage:"HTTP port"`
	Timeout  time.Duration `usage:"Request timeout"`
	DB struct {
		Host string `usage:"Database host"`
		Port int    `usage:"Database port"`
	}
}

cfg := Config{HTTPPort: 8080, Timeout: 3 * time.Second}
if err := goconfig.Load(&cfg); err != nil {
	log.Fatal(err)
}
```

### 2) Worker service

```go
type Config struct {
	Concurrency int           `usage:"Worker concurrency"`
	PollEvery   time.Duration `usage:"Polling interval"`
	Queues      []string      `usage:"Enabled queues"`
}

cfg := Config{Concurrency: 4, PollEvery: 2 * time.Second}
if err := goconfig.Load(&cfg); err != nil {
	log.Fatal(err)
}
```

Environment example:

```bash
export QUEUES='["emails", "billing"]'
export CONCURRENCY=8
```

### 3) CLI tool with deterministic args/env (tests)

```go
cfg := Config{}
err := goconfig.Load(&cfg,
	goconfig.WithArgs([]string{"-verbose", "-config", "./testdata/config.json"}),
	goconfig.WithEnvLookup(func(k string) (string, bool) {
		if k == "VERBOSE" {
			return "true", true
		}
		return "", false
	}),
	goconfig.WithoutImplicitConfigFile(),
)
```

## Precedence

Highest priority wins:

1. Command-line flags
2. Environment variables
3. JSON config file(s)
4. Struct default values

If `-config` is not provided and `./config.json` exists in the current working
directory, `goconfig` loads it automatically before env vars and flags.

## Multiple Config Files

`-config` accepts multiple JSON file paths separated by commas. Files are loaded
from left to right over the same struct, so later files override values from
previous files while omitted fields keep their existing values.

```bash
myapp -config ./config.base.json,./config.prod.json,./config.local.json
```

The same comma-separated format also works with `WithConfigFile`:

```go
err := goconfig.Load(&cfg,
	goconfig.WithConfigFile("./config.base.json,./config.prod.json"),
)
```

Environment variables and command-line flags are still applied after all JSON
files, so they keep higher precedence.

## Naming Convention

Given this struct:

```go
type Config struct {
	App struct {
		Port int
	}
}
```

- Flag name: `-app.port`
- Environment variable: `APP_PORT`
- JSON object:

```json
{
  "app": {
    "port": 8080
  }
}
```

## Built-in Flags

- `-help`: displays generated help with usage and env names
- `-config`: JSON file path or comma-separated paths to load before env and flags

When `-config` is not set, `goconfig` auto-loads `config.json` if it exists.

## Supported Types

- bool
- string
- float32, float64
- int, int32, int64
- uint, uint32, uint64
- slices (`[]T`, as JSON arrays for env/flags)
- nested structs
- pointers to structs
- `time.Duration` (duration string like `"15s"` or nanoseconds)

## Public API

### `Load`

`Load` returns errors instead of exiting. This is the recommended API for
libraries and services.

```go
err := goconfig.Load(&cfg)
```

Optional behavior can be controlled with options:

- `WithArgs([]string)`
- `WithProgramName("myapp")`
- `WithConfigFile("/etc/myapp/config.json")`
- `WithConfigFlagName("settings")`
- `WithImplicitConfigFile("myconfig.json")`
- `WithoutImplicitConfigFile()`
- `WithEnvLookup(func(string) (string, bool))`

### `Read`

`Read` keeps backward compatibility and exits process on error.

```go
goconfig.Read(&cfg)
```

## Why Teams Use goconfig

- Minimal integration cost for existing Go projects
- Predictable override order across local/dev/prod
- Built-in generated help for operations teams
- No external runtime dependencies

## Development

```bash
go test ./...
```

For local workflows:

```bash
make test
make coverage
```

## Open Source Health

- CI: GitHub Actions (`.github/workflows/ci.yml`)
- Contributing guide: `CONTRIBUTING.md`
- Code of conduct: `CODE_OF_CONDUCT.md`
- Security policy: `SECURITY.md`
- Changelog: `CHANGELOG.md`
- Release process: `RELEASING.md`
- Launch/promotion plan: `PROMOTION_PLAN.md`
- Issue and PR templates: `.github/ISSUE_TEMPLATE` and `.github/pull_request_template.md`

## Contributing

Issues and pull requests are welcome.

## Roadmap

See `ROADMAP.md` for the proposed adoption roadmap and high-impact issues.

## License

MIT License. See `LICENSE`.
