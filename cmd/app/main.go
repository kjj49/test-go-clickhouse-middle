package main

import (
	"log"

	"test-go-clickhouse-middle/config"
	"test-go-clickhouse-middle/internal/app"
)

// @title Test Go ClickHouse Middle
// @version 1.0
// @description API Service for Test Go ClickHouse Middle
// @host localhost:8080
func main() {
	// Configuration

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
