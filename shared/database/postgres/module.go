package postgres

import (
	"github.com/Nebuska/neblab/shared/database"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(database.NewConfig),
	fx.Provide(NewPostgres),
)
