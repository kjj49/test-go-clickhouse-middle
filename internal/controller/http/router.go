// Package v1 implements routing paths. Each services in own file.
package http

import (
	"test-go-clickhouse-middle/internal/usecase"
	"test-go-clickhouse-middle/pkg/logger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs
	_ "test-go-clickhouse-middle/docs"

	"github.com/gin-gonic/gin"
)

// NewRouter
func NewRouter(handler *gin.Engine, l logger.Interface, u usecase.Event) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Routers
	h := handler.Group("/api")
	{
		newEventRoutes(h, u, l)
	}
}
