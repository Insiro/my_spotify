package root

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
}

func (controller Controller) Heath(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func NewController() Controller {
	return Controller{}
}
