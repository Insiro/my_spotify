package cmd

import (
	controllers "github.com/Insiro/my_spotify/internal/controller/root"
	"github.com/Insiro/my_spotify/pkg/database"
	"github.com/Insiro/my_spotify/pkg/server"
	"go.uber.org/fx"
)

func StartServer() {
	fx.New(
		controllers.Modules,
		fx.Provide(database.NewPostgres),
		fx.Provide(server.New),
		fx.Invoke(server.Start),
	).Run()
}
