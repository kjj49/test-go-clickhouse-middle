// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test-go-clickhouse-middle/config"
	http "test-go-clickhouse-middle/internal/controller/http"
	"test-go-clickhouse-middle/internal/usecase"
	"test-go-clickhouse-middle/internal/usecase/repository"
	"test-go-clickhouse-middle/pkg/clickhouse"
	"test-go-clickhouse-middle/pkg/httpserver"
	"test-go-clickhouse-middle/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level)

	// Repository
	ch, err := clickhouse.New(
		[]string{fmt.Sprintf("%s:%d", cfg.ClickHouse.Host, cfg.ClickHouse.Port)},
		cfg.ClickHouse.DB, cfg.ClickHouse.Username, cfg.ClickHouse.Password)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - clickhouse.New: %w", err))
	}
	defer ch.Close()

	// Use case
	eventUseCase := usecase.New(
		repository.New(ch),
	)

	// HTTP Server
	handler := gin.New()
	http.NewRouter(handler, l, eventUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}
	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
