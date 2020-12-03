package handlers

import (
	"article_api/src/vo"
	"github.com/labstack/echo"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	healthResponse := vo.HealthResponse{
		Message: "API Service running OK",
	}
	return c.JSON(http.StatusOK, healthResponse)
}
