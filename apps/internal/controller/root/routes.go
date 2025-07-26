package root

import "github.com/labstack/echo/v4"

func BindRoutes(e *echo.Echo, c Controller) {
	e.GET("/", c.Heath)
}
