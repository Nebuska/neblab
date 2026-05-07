package main

import (
	"github.com/Nebuska/neblab/account/api"
	"github.com/Nebuska/neblab/account/internal"
	"github.com/Nebuska/neblab/shared/config"
	"github.com/Nebuska/neblab/shared/database/postgres"
	"github.com/Nebuska/neblab/shared/jwtAuth"
	"github.com/Nebuska/neblab/shared/logger"
	"github.com/Nebuska/neblab/shared/redis"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		logger.Module,
		config.Module,
		postgres.Module,
		redis.Module,
		jwtAuth.Module,
		internal.Module,

		fx.Provide(NewGinEngine),
		api.Module,

		fx.Invoke(GinStarter),
	)

	app.Run()
}
