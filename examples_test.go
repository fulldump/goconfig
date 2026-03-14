package goconfig

import "time"

func ExampleLoad() {
	type Server struct {
		Host string `usage:"Server host"`
		Port int    `usage:"Server port"`
	}

	type Config struct {
		Timeout time.Duration `usage:"Request timeout"`
		Server  Server
	}

	cfg := Config{
		Timeout: 5 * time.Second,
		Server: Server{
			Host: "127.0.0.1",
			Port: 8080,
		},
	}

	_ = Load(&cfg,
		WithArgs([]string{"-server.port", "9090"}),
		WithoutImplicitConfigFile(),
		WithEnvLookup(func(string) (string, bool) { return "", false }),
	)
}

func ExampleLoad_apiService() {
	type Config struct {
		HTTPPort int           `usage:"HTTP port"`
		Timeout  time.Duration `usage:"Request timeout"`
		DB       struct {
			Host string `usage:"Database host"`
			Port int    `usage:"Database port"`
		}
	}

	cfg := Config{
		HTTPPort: 8080,
		Timeout:  3 * time.Second,
	}

	_ = Load(&cfg)
}

func ExampleLoad_workerService() {
	type Config struct {
		Concurrency int           `usage:"Worker concurrency"`
		PollEvery   time.Duration `usage:"Polling interval"`
		Queues      []string      `usage:"Enabled queues"`
	}

	cfg := Config{
		Concurrency: 4,
		PollEvery:   2 * time.Second,
		Queues:      []string{"emails", "billing"},
	}

	_ = Load(&cfg)
}

func ExampleLoad_cliTool() {
	type Config struct {
		ConfigPath string `usage:"Path to config file"`
		Verbose    bool   `usage:"Enable verbose output"`
	}

	cfg := Config{}

	_ = Load(&cfg,
		WithArgs([]string{"-verbose", "-configpath", "./dev.json"}),
		WithoutImplicitConfigFile(),
	)
}
