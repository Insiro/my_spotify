package server

import (
	"context"
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
				go func() {
					err := e.Start(":8000")
					if err != nil {
						e.Logger.Error(err)
						panic(err)
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return e.Shutdown(ctx)
			},
		},
	)
}
