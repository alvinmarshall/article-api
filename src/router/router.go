package router

import (
	"article_api/src/api"
	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	api.MainGroup(e)
	return e
}
