package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func New() *echo.Echo {
	return echo.New()
}

func Start(lc fx.Lifecycle, e *echo.Echo) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				err := e.Start(":8000")
				if err != nil {
					e.Logger.Error(err)
					return err
				}
				e.GET("/", func(c echo.Context) error {
					return c.String(http.StatusOK, "Hello, World!")
				})
				fmt.Println("Server started on port 8000")

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return e.Shutdown(ctx)
			},
		},
	)
}
