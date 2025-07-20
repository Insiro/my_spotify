package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type RootController struct {
}

func (controller RootController) Heath(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func NewController() RootController {
	return RootController{}
}

func BindRoutes(e *echo.Echo, c RootController) {
	e.GET("/", c.Heath)
}
