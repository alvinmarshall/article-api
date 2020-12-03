package api

import (
	"article_api/src/handlers"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/health", handlers.HealthCheck)
}
