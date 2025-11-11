package aBoard

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(newBoardRepository),
	fx.Provide(NewBoardService),
)
