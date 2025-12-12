package mysql

import (
	"github.com/Nebuska/neblab/shared/database"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMySql),
	fx.Provide(database.NewConfig),
)
