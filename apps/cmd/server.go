package cmd

import (
	controllers "github.com/Insiro/my_spotify/internal/controller"
	"github.com/Insiro/my_spotify/pkg/server"
	"go.uber.org/fx"
)

func StartServer() {
	fx.New(
		controllers.Modules,
		fx.Provide(server.New),
		fx.Invoke(server.Start),
	).Run()
}
