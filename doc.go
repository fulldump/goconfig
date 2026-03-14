// Package goconfig loads configuration values into Go structs from three sources:
// JSON config files, environment variables and command-line flags.
//
// Sources are applied in this order:
//
//  1. JSON file
//  2. Environment variables
//  3. Command-line flags
//
// This means command-line flags have the highest precedence.
//
// Quick start:
//
//	type Config struct {
//		Port int `usage:"HTTP port"`
//	}
//
//	cfg := Config{Port: 8080}
//	if err := goconfig.Load(&cfg); err != nil {
//		log.Fatal(err)
//	}
package goconfig
