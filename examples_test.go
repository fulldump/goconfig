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
